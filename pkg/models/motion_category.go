package models

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
