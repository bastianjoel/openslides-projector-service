package models

type AssignmentCandidate struct {
	MeetingID     int
	MeetingUserID *int
	Weight        *int
	AssignmentID  int
	ID            *int
}

func (m AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}
