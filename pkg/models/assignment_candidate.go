package models

type AssignmentCandidate struct {
	AssignmentID  int
	ID            *int
	MeetingID     int
	MeetingUserID *int
	Weight        *int
}

func (m AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}
