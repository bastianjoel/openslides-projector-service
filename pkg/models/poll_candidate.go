package models

type PollCandidate struct {
	Weight              int
	ID                  *int
	MeetingID           int
	PollCandidateListID int
	UserID              *int
}

func (m PollCandidate) CollectionName() string {
	return "poll_candidate"
}
