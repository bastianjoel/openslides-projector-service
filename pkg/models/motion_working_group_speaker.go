package models

type MotionWorkingGroupSpeaker struct {
	Weight        *int
	ID            *int
	MeetingID     int
	MeetingUserID int
	MotionID      int
}

func (m MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}
