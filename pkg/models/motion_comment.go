package models

import "encoding/json"

type MotionComment struct {
	Comment   *string `json:"comment"`
	ID        int     `json:"id"`
	MeetingID int     `json:"meeting_id"`
	MotionID  int     `json:"motion_id"`
	SectionID int     `json:"section_id"`
}

func (m MotionComment) CollectionName() string {
	return "motion_comment"
}

func (m MotionComment) Get(field string) interface{} {
	switch field {
	case "comment":
		return m.Comment
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "section_id":
		return m.SectionID
	}

	return nil
}

func (m MotionComment) Update(data map[string]string) error {
	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
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

	if val, ok := data["motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["section_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SectionID)
		if err != nil {
			return err
		}
	}

	return nil
}
