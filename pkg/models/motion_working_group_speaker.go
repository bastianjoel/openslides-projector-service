package models

type MotionWorkingGroupSpeaker struct {
	ID            int  `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID int  `json:"meeting_user_id"`
	MotionID      int  `json:"motion_id"`
	Weight        *int `json:"weight"`
}

func (m MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}

func (m MotionWorkingGroupSpeaker) Get(field string) interface{} {
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
