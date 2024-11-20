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
	q := datastore.Collection(req.DB, &models.ListOfSpeakers{}).With("speaker_ids", nil).SetFqids(projection.ContentObjectID)
	speakersQ := q.GetSubquery("speaker_ids")
	meetingUsersQ := speakersQ.With("meeting_user_id", nil)
	meetingUsersQ.With("user_id", nil)

	losSub, err := q.SubscribeOne(&los)
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

	for _, speaker := range los.Speakers() {
		firstName := speaker.MeetingUser().User().Username
		fmt.Println(firstName)
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
