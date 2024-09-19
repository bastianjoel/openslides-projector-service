package models

type PollCandidateList struct {
	ID               *int  `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
}

func (m PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}
