package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type StructureLevelListOfSpeakers struct {
	AdditionalTime   *float32 `json:"additional_time"`
	CurrentStartTime *int     `json:"current_start_time"`
	ID               int      `json:"id"`
	InitialTime      int      `json:"initial_time"`
	ListOfSpeakersID int      `json:"list_of_speakers_id"`
	MeetingID        int      `json:"meeting_id"`
	RemainingTime    float32  `json:"remaining_time"`
	SpeakerIDs       []int    `json:"speaker_ids"`
	StructureLevelID int      `json:"structure_level_id"`
	loadedRelations  map[string]struct{}
	speakers         *Speaker
	structureLevel   *StructureLevel
	meeting          *Meeting
	listOfSpeakers   *ListOfSpeakers
}

func (m StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
}

func (m *StructureLevelListOfSpeakers) Speakers() *Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return m.speakers
}

func (m *StructureLevelListOfSpeakers) StructureLevel() StructureLevel {
	if _, ok := m.loadedRelations["structure_level_id"]; !ok {
		log.Panic().Msg("Tried to access StructureLevel relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.structureLevel
}

func (m *StructureLevelListOfSpeakers) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.meeting
}

func (m *StructureLevelListOfSpeakers) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m StructureLevelListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "additional_time":
		return m.AdditionalTime
	case "current_start_time":
		return m.CurrentStartTime
	case "id":
		return m.ID
	case "initial_time":
		return m.InitialTime
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "remaining_time":
		return m.RemainingTime
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_id":
		return m.StructureLevelID
	}

	return nil
}

func (m StructureLevelListOfSpeakers) Update(data map[string]string) error {
	if val, ok := data["additional_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdditionalTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["current_start_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentStartTime)
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

	if val, ok := data["initial_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.InitialTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersID)
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

	if val, ok := data["remaining_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.RemainingTime)
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

	if val, ok := data["structure_level_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelID)
		if err != nil {
			return err
		}
	}

	return nil
}
