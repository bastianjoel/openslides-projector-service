package models

type MotionSubmitter struct {
	ID            *int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        *int
}

func (m MotionSubmitter) CollectionName() string {
	return "motion_submitter"
}
