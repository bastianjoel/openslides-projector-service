package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/rs/zerolog/log"
)

func CurrentListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	projection := req.Projection

	var los models.ListOfSpeakers
	losSub, err := datastore.Collection(req.DB, &models.ListOfSpeakers{}).SetFqids(projection.ContentObjectID).SubscribeOne(&los)
	if err != nil {
		return nil, fmt.Errorf("CurrentListOfSpeakersSlideHandler: %w", err)
	}

	go func() {
		content <- getCurrentListOfSpeakersSlideContent(&los)

		for {
			select {
			case <-losSub.Channel:
				content <- getCurrentListOfSpeakersSlideContent(&los)
			}
		}
	}()

	return content, nil
}

func getCurrentListOfSpeakersSlideContent(los *models.ListOfSpeakers) string {
	tmpl, err := template.ParseFiles("templates/slides/current-list-of-speakers.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load current-list-of-speakers template")
		return ""
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"ListOfSpeakers": los,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute current-list-of-speakers template")
		return ""
	}

	return content.String()
}
