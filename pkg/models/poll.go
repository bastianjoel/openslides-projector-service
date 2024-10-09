package models

import "encoding/json"

type Poll struct {
	Backend               string          `json:"backend"`
	ContentObjectID       string          `json:"content_object_id"`
	CryptKey              *string         `json:"crypt_key"`
	CryptSignature        *string         `json:"crypt_signature"`
	Description           *string         `json:"description"`
	EntitledGroupIDs      []int           `json:"entitled_group_ids"`
	EntitledUsersAtStop   json.RawMessage `json:"entitled_users_at_stop"`
	GlobalAbstain         *bool           `json:"global_abstain"`
	GlobalNo              *bool           `json:"global_no"`
	GlobalOptionID        *int            `json:"global_option_id"`
	GlobalYes             *bool           `json:"global_yes"`
	ID                    int             `json:"id"`
	IsPseudoanonymized    *bool           `json:"is_pseudoanonymized"`
	MaxVotesAmount        *int            `json:"max_votes_amount"`
	MaxVotesPerOption     *int            `json:"max_votes_per_option"`
	MeetingID             int             `json:"meeting_id"`
	MinVotesAmount        *int            `json:"min_votes_amount"`
	OnehundredPercentBase string          `json:"onehundred_percent_base"`
	OptionIDs             []int           `json:"option_ids"`
	Pollmethod            string          `json:"pollmethod"`
	ProjectionIDs         []int           `json:"projection_ids"`
	SequentialNumber      int             `json:"sequential_number"`
	State                 *string         `json:"state"`
	Title                 string          `json:"title"`
	Type                  string          `json:"type"`
	VoteCount             *int            `json:"vote_count"`
	VotedIDs              []int           `json:"voted_ids"`
	VotesRaw              *string         `json:"votes_raw"`
	VotesSignature        *string         `json:"votes_signature"`
	Votescast             *string         `json:"votescast"`
	Votesinvalid          *string         `json:"votesinvalid"`
	Votesvalid            *string         `json:"votesvalid"`
}

func (m Poll) CollectionName() string {
	return "poll"
}
