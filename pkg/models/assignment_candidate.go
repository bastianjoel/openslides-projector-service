package models

type AssignmentCandidate struct {
	AssignmentID  int  `json:"assignment_id"`
	ID            int  `json:"id"`
	MeetingID     int  `json:"meeting_id"`
	MeetingUserID *int `json:"meeting_user_id"`
	Weight        *int `json:"weight"`
}

func (m AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}

func (m AssignmentCandidate) Get(field string) interface{} {
	switch field {
	case "assignment_id":
		return m.AssignmentID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "weight":
		return m.Weight
	}

	return nil
}
