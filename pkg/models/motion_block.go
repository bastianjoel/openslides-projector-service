package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionBlock struct {
	AgendaItemID     *int   `json:"agenda_item_id"`
	ID               int    `json:"id"`
	Internal         *bool  `json:"internal"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	MeetingID        int    `json:"meeting_id"`
	MotionIDs        []int  `json:"motion_ids"`
	ProjectionIDs    []int  `json:"projection_ids"`
	SequentialNumber int    `json:"sequential_number"`
	Title            string `json:"title"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	motions          *Motion
	projections      *Projection
	agendaItem       *AgendaItem
	listOfSpeakers   *ListOfSpeakers
}

func (m *MotionBlock) CollectionName() string {
	return "motion_block"
}

func (m *MotionBlock) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionBlock which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionBlock) Motions() *Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionBlock which was not loaded.")
	}

	return m.motions
}

func (m *MotionBlock) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of MotionBlock which was not loaded.")
	}

	return m.projections
}

func (m *MotionBlock) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of MotionBlock which was not loaded.")
	}

	return m.agendaItem
}

func (m *MotionBlock) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of MotionBlock which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *MotionBlock) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "title":
		return m.Title
	}

	return nil
}

func (m *MotionBlock) Update(data map[string]string) error {
	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
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

	if val, ok := data["internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.Internal)
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

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
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

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
		if err != nil {
			return err
		}
	}

	return nil
}
