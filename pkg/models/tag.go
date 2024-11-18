package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Tag struct {
	ID              int      `json:"id"`
	MeetingID       int      `json:"meeting_id"`
	Name            string   `json:"name"`
	TaggedIDs       []string `json:"tagged_ids"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
}

func (m *Tag) CollectionName() string {
	return "tag"
}

func (m *Tag) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Tag which was not loaded.")
	}

	return *m.meeting
}

func (m *Tag) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "tagged_ids":
		return m.TaggedIDs
	}

	return nil
}

func (m *Tag) Update(data map[string]string) error {
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["tagged_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TaggedIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
