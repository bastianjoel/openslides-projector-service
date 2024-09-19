package models

type MotionComment struct {
	Comment   *string `json:"comment"`
	ID        *int    `json:"id"`
	MeetingID int     `json:"meeting_id"`
	MotionID  int     `json:"motion_id"`
	SectionID int     `json:"section_id"`
}

func (m MotionComment) CollectionName() string {
	return "motion_comment"
}
