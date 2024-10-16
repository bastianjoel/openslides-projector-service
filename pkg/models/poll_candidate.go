package models

type PollCandidate struct {
	ID                  int  `json:"id"`
	MeetingID           int  `json:"meeting_id"`
	PollCandidateListID int  `json:"poll_candidate_list_id"`
	UserID              *int `json:"user_id"`
	Weight              int  `json:"weight"`
}

func (m PollCandidate) CollectionName() string {
	return "poll_candidate"
}

func (m PollCandidate) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "poll_candidate_list_id":
		return m.PollCandidateListID
	case "user_id":
		return m.UserID
	case "weight":
		return m.Weight
	}

	return nil
}
