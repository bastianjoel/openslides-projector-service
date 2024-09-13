package models

type Option struct {
	Abstain                    *string
	ContentObjectID            *string
	ID                         *int
	MeetingID                  int
	No                         *string
	PollID                     *int
	Text                       *string
	UsedAsGlobalOptionInPollID *int
	VoteIDs                    []int
	Weight                     *int
	Yes                        *string
}

func (m Option) CollectionName() string {
	return "option"
}
