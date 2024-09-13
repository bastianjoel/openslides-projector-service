package models

type MotionWorkingGroupSpeaker struct {
	ID            *int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        *int
}

func (m MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}
