package projector

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projector struct {
	Content string
	Data    models.Projector
	Channel chan projectorUpdateEvent
}

type projectorUpdateEvent struct{}

func newProjector(id int, db *datastore.Datastore) (*projector, error) {
	projectorData, err := db.Collection(&models.Projector{}).SetIds(id).AsSingle().Run()
	if err != nil {
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}

	tmpl, err := template.ParseFiles("templates/projector.html")
	if err != nil {
		return nil, fmt.Errorf("error reading projector template %w", err)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Projector": projectorData,
	})
	if err != nil {
		return nil, fmt.Errorf("error generating projector template %w", err)
	}

	return &projector{
		Content: content.String(),
		Data:    projectorData.(models.Projector),
		Channel: make(chan projectorUpdateEvent),
	}, nil
}
