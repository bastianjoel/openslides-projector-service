package models

type MotionComment struct {
	Comment   *string
	ID        *int
	MeetingID int
	MotionID  int
	SectionID int
}

func (m MotionComment) CollectionName() string {
	return "motion_comment"
}
