package models

type PollCandidate struct {
	ID                  *int `json:"id"`
	MeetingID           int  `json:"meeting_id"`
	PollCandidateListID int  `json:"poll_candidate_list_id"`
	UserID              *int `json:"user_id"`
	Weight              int  `json:"weight"`
}

func (m PollCandidate) CollectionName() string {
	return "poll_candidate"
}
