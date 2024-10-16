package models

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
