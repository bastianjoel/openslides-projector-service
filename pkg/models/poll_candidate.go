package models

type PollCandidate struct {
	ID                  *int
	MeetingID           int
	PollCandidateListID int
	UserID              *int
	Weight              int
}

func (m PollCandidate) CollectionName() string {
	return "poll_candidate"
}
