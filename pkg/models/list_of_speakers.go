package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type ListOfSpeakers struct {
	Closed                          *bool  `json:"closed"`
	ContentObjectID                 string `json:"content_object_id"`
	ID                              int    `json:"id"`
	MeetingID                       int    `json:"meeting_id"`
	ProjectionIDs                   []int  `json:"projection_ids"`
	SequentialNumber                int    `json:"sequential_number"`
	SpeakerIDs                      []int  `json:"speaker_ids"`
	StructureLevelListOfSpeakersIDs []int  `json:"structure_level_list_of_speakers_ids"`
	loadedRelations                 map[string]struct{}
	projections                     *Projection
	speakers                        *Speaker
	structureLevelListOfSpeakerss   *StructureLevelListOfSpeakers
	meeting                         *Meeting
}

func (m ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}

func (m *ListOfSpeakers) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ListOfSpeakers which was not loaded.")
	}

	return m.projections
}

func (m *ListOfSpeakers) Speakers() *Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of ListOfSpeakers which was not loaded.")
	}

	return m.speakers
}

func (m *ListOfSpeakers) StructureLevelListOfSpeakerss() *StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of ListOfSpeakers which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *ListOfSpeakers) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ListOfSpeakers which was not loaded.")
	}

	return *m.meeting
}

func (m ListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "closed":
		return m.Closed
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}

func (m ListOfSpeakers) Update(data map[string]string) error {
	if val, ok := data["closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.Closed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content_object_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ContentObjectID)
		if err != nil {
			return err
		}
	}

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

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
