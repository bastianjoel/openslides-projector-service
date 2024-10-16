package models

type PollCandidateList struct {
	ID               int   `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
}

func (m PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}

func (m PollCandidateList) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	}

	return nil
}
