package models

type MotionWorkingGroupSpeaker struct {
	ID            *int `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID int  `json:"meeting_user_id"`
	MotionID      int  `json:"motion_id"`
	Weight        *int `json:"weight"`
}

func (m MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}
