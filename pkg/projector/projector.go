package projector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector/slide"
)

type projector struct {
	db             *datastore.Datastore
	slideRouter    *slide.SlideRouter
	projector      *models.Projector
	listeners      []chan *ProjectorUpdateEvent
	Content        string
	Projections    map[int]string
	Data           models.Projector
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

	p := &projector{
		db:             db,
		projector:      data,
		slideRouter:    slide.New(db, id),
		Projections:    make(map[int]string),
		AddListener:    make(chan chan *ProjectorUpdateEvent),
		RemoveListener: make(chan (<-chan *ProjectorUpdateEvent)),
	}

	if err = p.updateFullContent(data); err != nil {
		return nil, fmt.Errorf("error generating projector content: %w", err)
	}

	go p.subscribeProjector()

	return p, nil
}

func (p *projector) subscribeProjector() {
	projectorSub := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).Subscribe()
	defer projectorSub.Unsubscribe()

	projectionUpdate, projections := p.getProjectionSubscription()
	// TODO: Subscribe on projector settings updates
	// Ignore e.g. projection defaults and [...]_projection_ids
	// If header active: Meeting name + description need to be subscribed

	// TODO: Slides subscription needs to pass events to projector channel

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
				fmt.Println(err)
			}

			p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
		case data, ok := <-projectionUpdate:
			if !ok {
				return
			}

			if data == nil {
				continue
			}

			for _, projectionId := range data {
				if projection, ok := projections[projectionId]; ok {
					p.Projections[projectionId] = projection
					p.sendToAll(&ProjectorUpdateEvent{"projection-updated", projection})
				} else {
					delete(p.Projections, projectionId)
					p.sendToAll(&ProjectorUpdateEvent{"projection-deleted", strconv.Itoa(projectionId)})
				}
			}
		}
	}
}

func (p *projector) sendToAll(event *ProjectorUpdateEvent) {
	for _, listener := range p.listeners {
		listener <- event
	}
}

func (p *projector) updateFullContent(projector *models.Projector) error {
	tmpl, err := template.ParseFiles("templates/projector.html")
	if err != nil {
		return fmt.Errorf("error reading projector template %w", err)
	}

	// TODO: Queue projections update if `p.Projections` is nil

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Projector":   projector,
		"Projections": p.Projections,
	})
	if err != nil {
		return fmt.Errorf("error generating projector template %w", err)
	}

	p.Content = content.String()
	p.Data = *projector

	return nil
}

func (p *projector) getProjectionSubscription() (<-chan []int, map[int]string) {
	query := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).SetFields("current_projection_ids")
	updateChannel := make(chan []int)
	projections := make(map[int]string)

	go func() {
		projectionIDs := []int{}
		sub := query.SubscribeField(&projectionIDs)
		defer sub.Unsubscribe()

		projector, err := query.GetOne()
		if err != nil {
			fmt.Println("retrieving current projection ids: %w", err)
		}

		updateChannel <- projector.CurrentProjectionIDs

		for range sub.Channel {
			updated := []int{}
			for id := range projections {
				if !slices.Contains(projectionIDs, id) {
					updated = append(updated, id)
					delete(projections, id)
				}
			}

			for _, id := range projectionIDs {
				if _, ok := projections[id]; !ok {
					updated = append(updated, id)
					projections[id] = ""
				}
			}

			updateChannel <- updated
		}
	}()

	return updateChannel, projections
}
