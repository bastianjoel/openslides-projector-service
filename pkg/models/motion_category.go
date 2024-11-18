package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionCategory struct {
	ChildIDs         []int   `json:"child_ids"`
	ID               int     `json:"id"`
	Level            *int    `json:"level"`
	MeetingID        int     `json:"meeting_id"`
	MotionIDs        []int   `json:"motion_ids"`
	Name             string  `json:"name"`
	ParentID         *int    `json:"parent_id"`
	Prefix           *string `json:"prefix"`
	SequentialNumber int     `json:"sequential_number"`
	Weight           *int    `json:"weight"`
	loadedRelations  map[string]struct{}
	motions          *Motion
	parent           *MotionCategory
	childs           *MotionCategory
	meeting          *Meeting
}

func (m *MotionCategory) CollectionName() string {
	return "motion_category"
}

func (m *MotionCategory) Motions() *Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionCategory which was not loaded.")
	}

	return m.motions
}

func (m *MotionCategory) Parent() *MotionCategory {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of MotionCategory which was not loaded.")
	}

	return m.parent
}

func (m *MotionCategory) Childs() *MotionCategory {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of MotionCategory which was not loaded.")
	}

	return m.childs
}

func (m *MotionCategory) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionCategory which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionCategory) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "id":
		return m.ID
	case "level":
		return m.Level
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "name":
		return m.Name
	case "parent_id":
		return m.ParentID
	case "prefix":
		return m.Prefix
	case "sequential_number":
		return m.SequentialNumber
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionCategory) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
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

	if val, ok := data["level"]; ok {
		err := json.Unmarshal([]byte(val), &m.Level)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ParentID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.Prefix)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
