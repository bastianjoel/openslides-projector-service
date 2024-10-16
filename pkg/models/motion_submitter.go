package models

type MotionSubmitter struct {
	ID            int  `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID int  `json:"meeting_user_id"`
	MotionID      int  `json:"motion_id"`
	Weight        *int `json:"weight"`
}

func (m MotionSubmitter) CollectionName() string {
	return "motion_submitter"
}

func (m MotionSubmitter) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "motion_id":
		return m.MotionID
	case "weight":
		return m.Weight
	}

	return nil
}
