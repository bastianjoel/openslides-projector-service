package models

type MotionComment struct {
	MeetingID int
	MotionID  int
	SectionID int
	Comment   *string
	ID        *int
}

func (m MotionComment) CollectionName() string {
	return "motion_comment"
}
