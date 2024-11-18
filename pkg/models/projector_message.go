package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type ProjectorMessage struct {
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	Message         *string `json:"message"`
	ProjectionIDs   []int   `json:"projection_ids"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	projections     *Projection
}

func (m *ProjectorMessage) CollectionName() string {
	return "projector_message"
}

func (m *ProjectorMessage) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ProjectorMessage which was not loaded.")
	}

	return *m.meeting
}

func (m *ProjectorMessage) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ProjectorMessage which was not loaded.")
	}

	return m.projections
}

func (m *ProjectorMessage) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "message":
		return m.Message
	case "projection_ids":
		return m.ProjectionIDs
	}

	return nil
}

func (m *ProjectorMessage) Update(data map[string]string) error {
	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["message"]; ok {
		err := json.Unmarshal([]byte(val), &m.Message)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
