package models

type PollCandidateList struct {
	MeetingID        int
	OptionID         int
	PollCandidateIDs []int
	ID               *int
}

func (m PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}
