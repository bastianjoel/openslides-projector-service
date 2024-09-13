package models

type MotionEditor struct {
	MeetingUserID int
	MotionID      int
	Weight        *int
	ID            *int
	MeetingID     int
}

func (m MotionEditor) CollectionName() string {
	return "motion_editor"
}
