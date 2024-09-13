package models

import "encoding/json"

type Poll struct {
	Backend               string
	ContentObjectID       string
	CryptKey              *string
	CryptSignature        *string
	Description           *string
	EntitledGroupIDs      []int
	EntitledUsersAtStop   json.RawMessage
	GlobalAbstain         *bool
	GlobalNo              *bool
	GlobalOptionID        *int
	GlobalYes             *bool
	ID                    *int
	IsPseudoanonymized    *bool
	MaxVotesAmount        *int
	MaxVotesPerOption     *int
	MeetingID             int
	MinVotesAmount        *int
	OnehundredPercentBase string
	OptionIDs             []int
	Pollmethod            string
	ProjectionIDs         []int
	SequentialNumber      int
	State                 *string
	Title                 string
	Type                  string
	VoteCount             *int
	VotedIDs              []int
	VotesRaw              *string
	VotesSignature        *string
	Votescast             *string
	Votesinvalid          *string
	Votesvalid            *string
}

func (m Poll) CollectionName() string {
	return "poll"
}
