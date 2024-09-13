package models

import "encoding/json"

type Poll struct {
	Title                 string
	Backend               string
	ContentObjectID       string
	CryptSignature        *string
	EntitledUsersAtStop   json.RawMessage
	MeetingID             int
	MinVotesAmount        *int
	Pollmethod            string
	EntitledGroupIDs      []int
	GlobalOptionID        *int
	SequentialNumber      int
	VotesSignature        *string
	CryptKey              *string
	GlobalAbstain         *bool
	VotedIDs              []int
	ProjectionIDs         []int
	State                 *string
	VoteCount             *int
	VotesRaw              *string
	Votesvalid            *string
	GlobalNo              *bool
	MaxVotesAmount        *int
	MaxVotesPerOption     *int
	ID                    *int
	IsPseudoanonymized    *bool
	OnehundredPercentBase string
	Votesinvalid          *string
	Type                  string
	Description           *string
	GlobalYes             *bool
	OptionIDs             []int
	Votescast             *string
}

func (m Poll) CollectionName() string {
	return "poll"
}
