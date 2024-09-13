package models

type MotionEditor struct {
	ID            *int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        *int
}

func (m MotionEditor) CollectionName() string {
	return "motion_editor"
}
