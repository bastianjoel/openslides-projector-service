package models

type Option struct {
	Abstain                    *string
	UsedAsGlobalOptionInPollID *int
	VoteIDs                    []int
	Yes                        *string
	ContentObjectID            *string
	ID                         *int
	MeetingID                  int
	No                         *string
	PollID                     *int
	Text                       *string
	Weight                     *int
}

func (m Option) CollectionName() string {
	return "option"
}
