package models

import "encoding/json"

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
}

func (m MotionCategory) CollectionName() string {
	return "motion_category"
}

func (m MotionCategory) Get(field string) interface{} {
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

func (m MotionCategory) Update(data map[string]string) error {
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
