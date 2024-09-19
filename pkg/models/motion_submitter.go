package models

type MotionSubmitter struct {
	ID            *int `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID int  `json:"meeting_user_id"`
	MotionID      int  `json:"motion_id"`
	Weight        *int `json:"weight"`
}

func (m MotionSubmitter) CollectionName() string {
	return "motion_submitter"
}
