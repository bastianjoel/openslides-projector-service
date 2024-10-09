package models

type MotionEditor struct {
	ID            int  `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID int  `json:"meeting_user_id"`
	MotionID      int  `json:"motion_id"`
	Weight        *int `json:"weight"`
}

func (m MotionEditor) CollectionName() string {
	return "motion_editor"
}
