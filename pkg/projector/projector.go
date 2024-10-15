package projector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector/slide"
	"github.com/rs/zerolog/log"
)

type projector struct {
	ctxCancel      context.CancelFunc
	db             *datastore.Datastore
	slideRouter    *slide.SlideRouter
	projector      *models.Projector
	listeners      []chan *ProjectorUpdateEvent
	Content        string
	Projections    map[int]string
	AddListener    chan chan *ProjectorUpdateEvent
	RemoveListener chan (<-chan *ProjectorUpdateEvent)
}

type ProjectorUpdateEvent struct {
	Event string
	Data  string
}

func newProjector(id int, db *datastore.Datastore) (*projector, error) {
	projectorQuery := datastore.Collection(db, &models.Projector{}).SetIds(id)
	data, err := projectorQuery.GetOne()
	if err != nil {
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}

	if data == nil {
		return nil, fmt.Errorf("projector not found")
	}

	ctx, cancel := context.WithCancel(context.Background())
	p := &projector{
		ctxCancel:      cancel,
		db:             db,
		projector:      data,
		slideRouter:    slide.New(ctx, db),
		Projections:    make(map[int]string),
		AddListener:    make(chan chan *ProjectorUpdateEvent),
		RemoveListener: make(chan (<-chan *ProjectorUpdateEvent)),
	}

	if err = p.updateFullContent(); err != nil {
		return nil, fmt.Errorf("error generating projector content: %w", err)
	}

	go p.subscribeProjector()

	return p, nil
}

func (p *projector) subscribeProjector() {
	defer p.ctxCancel()
	// TODO: Subscribe on projector settings updates
	// Ignore e.g. projection defaults and [...]_projection_ids
	// If header active: Meeting name + description need to be subscribed
	projectorSub := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).Subscribe()
	defer projectorSub.Unsubscribe()

	projectionUpdate, projections, unsubscibeProjections := p.getProjectionSubscription()
	defer unsubscibeProjections()

	for {
		select {
		case listener := <-p.AddListener:
			p.listeners = append(p.listeners, listener)
			listener <- &ProjectorUpdateEvent{
				Event: "connected",
				Data:  "",
			}
		case listener := <-p.RemoveListener:
			i := slices.IndexFunc(p.listeners, func(el chan *ProjectorUpdateEvent) bool { return el == listener })
			if i > -1 {
				close(p.listeners[i])
				p.listeners[i] = p.listeners[len(p.listeners)-1]
				p.listeners = p.listeners[:len(p.listeners)-1]
			}
		case data, ok := <-projectorSub.Channel:
			if !ok {
				return
			}

			projectorData := data["projector/"+strconv.Itoa(p.projector.ID)]
			if projectorData == nil {
				p.sendToAll(&ProjectorUpdateEvent{"deleted", ""})
				return
			}

			encodedData, err := json.Marshal(projectorData)
			if err != nil {
				log.Error().Err(err).Msg("could not encode projector data")
			} else {
				p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
			}
		case data, ok := <-projectionUpdate:
			if !ok {
				return
			}

			p.processProjectionUpdate(data, projections)
		}
	}
}

func (p *projector) processProjectionUpdate(updated []int, projections map[int]string) {
	if updated == nil {
		return
	}

	for _, projectionId := range updated {
		if projection, ok := projections[projectionId]; ok {
			p.Projections[projectionId] = projection
			defer p.sendToAll(&ProjectorUpdateEvent{"projection-updated", projection})
		} else {
			delete(p.Projections, projectionId)
			defer p.sendToAll(&ProjectorUpdateEvent{"projection-deleted", strconv.Itoa(projectionId)})
		}
	}

	if err := p.updateFullContent(); err != nil {
		log.Error().Err(err).Msg("failed to generate projector content")
	}
}

func (p *projector) sendToAll(event *ProjectorUpdateEvent) {
	for _, listener := range p.listeners {
		listener <- event
	}
}

func (p *projector) updateFullContent() error {
	tmpl, err := template.ParseFiles("templates/projector.html")
	if err != nil {
		return fmt.Errorf("error reading projector template %w", err)
	}

	// TODO: Queue projections update if `p.Projections` is nil

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Projector":   p.projector,
		"Projections": p.Projections,
	})
	if err != nil {
		return fmt.Errorf("error generating projector template %w", err)
	}

	p.Content = content.String()

	return nil
}

func (p *projector) getProjectionSubscription() (<-chan []int, map[int]string, func()) {
	query := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).SetFields("current_projection_ids")
	updateChannel := make(chan []int)
	projections := make(map[int]string)
	addProjection := make(chan int)
	removeProjection := make(chan int)

	go func() {
		projectionIDs := []int{}
		sub := query.SubscribeField(&projectionIDs)
		defer sub.Unsubscribe()

		projector, err := query.GetOne()
		if err != nil {
			log.Warn().Err(err).Msg("retrieving current projection ids")
		}

		projectionChannel := p.slideRouter.SubscribeContent(addProjection, removeProjection)
		for _, id := range projector.CurrentProjectionIDs {
			addProjection <- id
			projections[id] = ""
		}

		for {
			select {
			case <-sub.Channel:
				updated := []int{}
				for id := range projections {
					if !slices.Contains(projectionIDs, id) {
						updated = append(updated, id)
						removeProjection <- id
						delete(projections, id)
					}
				}

				for _, id := range projectionIDs {
					if _, ok := projections[id]; !ok {
						updated = append(updated, id)
						addProjection <- id
						projections[id] = ""
					}
				}

				updateChannel <- updated
			case update := <-projectionChannel:
				projections[update.ID] = update.Content
				updateChannel <- []int{update.ID}
			}
		}
	}()

	return updateChannel, projections, func() {
		// TODO: Unsubscribe
	}
}
