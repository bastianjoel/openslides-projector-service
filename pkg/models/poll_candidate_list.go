package models

type PollCandidateList struct {
	ID               *int
	MeetingID        int
	OptionID         int
	PollCandidateIDs []int
}

func (m PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}
