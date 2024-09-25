package projector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projector struct {
	db             *datastore.Datastore
	id             int
	listeners      []chan *ProjectorUpdateEvent
	Content        string
	Data           models.Projector
	AddListener    chan chan *ProjectorUpdateEvent
	RemoveListener chan (<-chan *ProjectorUpdateEvent)
}

type ProjectorUpdateEvent struct {
	Event string
	Data  string
}

func newProjector(id int, db *datastore.Datastore) (*projector, error) {
	p := &projector{
		db:             db,
		id:             id,
		AddListener:    make(chan chan *ProjectorUpdateEvent),
		RemoveListener: make(chan (<-chan *ProjectorUpdateEvent)),
	}
	err := p.updateFullContent()
	if err != nil {
		return nil, fmt.Errorf("error generating projector content: %w", err)
	}

	go p.subscribeProjector()

	return p, nil
}

func (p *projector) subscribeProjector() {
	subscription := p.db.Collection(&models.Projector{}).SetIds(p.id).AsSingle().Subscribe()
	// defer p.db.Collection(&models.Projector{}).SetIds(p.id).AsSingle().Unsubscribe()
	// TODO: Subscribe on projector settings updates
	// Ignore e.g. projection defaults and [...]_projection_ids
	// If header active: Meeting name + description need to be subscribed

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
		case data, ok := <-subscription:
			if !ok {
				return
			}

			projectorData := data["projector/"+strconv.Itoa(p.id)]
			if projectorData == nil {
				p.sendToAll(&ProjectorUpdateEvent{"deleted", ""})
				return
			}

			encodedData, err := json.Marshal(projectorData)
			if err != nil {
				fmt.Println(err)
			}

			p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
		}
	}

	// TODO: Slides subscription needs to pass events to projector channel
	// TODO: Handle projector deletion
}

func (p *projector) sendToAll(event *ProjectorUpdateEvent) {
	for _, listener := range p.listeners {
		listener <- event
	}
}

func (p *projector) updateFullContent() error {
	projectorQuery := p.db.Collection(&models.Projector{}).SetIds(p.id).AsSingle()
	data, err := projectorQuery.Run()
	if err != nil {
		return fmt.Errorf("error fetching projector from db %w", err)
	}

	if data == nil {
		return fmt.Errorf("projector not found")
	}
	projectorData := data.(*models.Projector)

	tmpl, err := template.ParseFiles("templates/projector.html")
	if err != nil {
		return fmt.Errorf("error reading projector template %w", err)
	}

	var projections []string
	for _, id := range projectorData.CurrentProjectionIDs {
		projection, err := p.getProjectionContent(id)
		if err != nil {
			log.Print(fmt.Errorf("error calculating projection: %w", err))
		}

		if projection != nil {
			projections = append(projections, *projection)
		}
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Projector":   projectorData,
		"Projections": projections,
	})
	if err != nil {
		return fmt.Errorf("error generating projector template %w", err)
	}

	p.Content = content.String()
	p.Data = *projectorData

	return nil
}

func (p *projector) getProjectionContent(id int) (*string, error) {
	data, err := p.db.Collection(&models.Projection{}).SetIds(p.id).AsSingle().Run()
	if err != nil {
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}
	projectionData := data.(*models.Projection)

	return projectionData.Type, nil
}
