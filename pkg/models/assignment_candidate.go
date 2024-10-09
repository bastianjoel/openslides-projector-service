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
