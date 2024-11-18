package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type StructureLevel struct {
	Color                           *string `json:"color"`
	DefaultTime                     *int    `json:"default_time"`
	ID                              int     `json:"id"`
	MeetingID                       int     `json:"meeting_id"`
	MeetingUserIDs                  []int   `json:"meeting_user_ids"`
	Name                            string  `json:"name"`
	StructureLevelListOfSpeakersIDs []int   `json:"structure_level_list_of_speakers_ids"`
	loadedRelations                 map[string]struct{}
	structureLevelListOfSpeakerss   []StructureLevelListOfSpeakers
	meeting                         *Meeting
	meetingUsers                    []MeetingUser
}

func (m *StructureLevel) CollectionName() string {
	return "structure_level"
}

func (m *StructureLevel) StructureLevelListOfSpeakerss() []StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of StructureLevel which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *StructureLevel) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of StructureLevel which was not loaded.")
	}

	return *m.meeting
}

func (m *StructureLevel) MeetingUsers() []MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of StructureLevel which was not loaded.")
	}

	return m.meetingUsers
}

func (m *StructureLevel) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "structure_level_list_of_speakers_ids":
			m.structureLevelListOfSpeakerss = content.([]StructureLevelListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_ids":
			m.meetingUsers = content.([]MeetingUser)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *StructureLevel) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "structure_level_list_of_speakers_ids":
		err := json.Unmarshal(content, &m.structureLevelListOfSpeakerss)
		if err != nil {
			return err
		}
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "meeting_user_ids":
		err := json.Unmarshal(content, &m.meetingUsers)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return nil
}

func (m *StructureLevel) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "default_time":
		return m.DefaultTime
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}

func (m *StructureLevel) GetFqids(field string) []string {
	switch field {
	case "structure_level_list_of_speakers_ids":
		r := make([]string, len(m.StructureLevelListOfSpeakersIDs))
		for i, id := range m.StructureLevelListOfSpeakersIDs {
			r[i] = "structure_level_list_of_speakers/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *StructureLevel) Update(data map[string]string) error {
	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultTime)
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

	if val, ok := data["meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserIDs)
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

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
