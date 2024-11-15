package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionChangeRecommendation struct {
	CreationTime     *int    `json:"creation_time"`
	ID               int     `json:"id"`
	Internal         *bool   `json:"internal"`
	LineFrom         *int    `json:"line_from"`
	LineTo           *int    `json:"line_to"`
	MeetingID        int     `json:"meeting_id"`
	MotionID         int     `json:"motion_id"`
	OtherDescription *string `json:"other_description"`
	Rejected         *bool   `json:"rejected"`
	Text             *string `json:"text"`
	Type             *string `json:"type"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	motion           *Motion
}

func (m MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}

func (m *MotionChangeRecommendation) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionChangeRecommendation) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.motion
}

func (m MotionChangeRecommendation) Get(field string) interface{} {
	switch field {
	case "creation_time":
		return m.CreationTime
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "line_from":
		return m.LineFrom
	case "line_to":
		return m.LineTo
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "other_description":
		return m.OtherDescription
	case "rejected":
		return m.Rejected
	case "text":
		return m.Text
	case "type":
		return m.Type
	}

	return nil
}

func (m MotionChangeRecommendation) Update(data map[string]string) error {
	if val, ok := data["creation_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CreationTime)
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

	if val, ok := data["line_from"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineFrom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["line_to"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineTo)
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

	if val, ok := data["motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["other_description"]; ok {
		err := json.Unmarshal([]byte(val), &m.OtherDescription)
		if err != nil {
			return err
		}
	}

	if val, ok := data["rejected"]; ok {
		err := json.Unmarshal([]byte(val), &m.Rejected)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
		if err != nil {
			return err
		}
	}

	return nil
}
