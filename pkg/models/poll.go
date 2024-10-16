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

func (m Poll) Get(field string) interface{} {
	switch field {
	case "backend":
		return m.Backend
	case "content_object_id":
		return m.ContentObjectID
	case "crypt_key":
		return m.CryptKey
	case "crypt_signature":
		return m.CryptSignature
	case "description":
		return m.Description
	case "entitled_group_ids":
		return m.EntitledGroupIDs
	case "entitled_users_at_stop":
		return m.EntitledUsersAtStop
	case "global_abstain":
		return m.GlobalAbstain
	case "global_no":
		return m.GlobalNo
	case "global_option_id":
		return m.GlobalOptionID
	case "global_yes":
		return m.GlobalYes
	case "id":
		return m.ID
	case "is_pseudoanonymized":
		return m.IsPseudoanonymized
	case "max_votes_amount":
		return m.MaxVotesAmount
	case "max_votes_per_option":
		return m.MaxVotesPerOption
	case "meeting_id":
		return m.MeetingID
	case "min_votes_amount":
		return m.MinVotesAmount
	case "onehundred_percent_base":
		return m.OnehundredPercentBase
	case "option_ids":
		return m.OptionIDs
	case "pollmethod":
		return m.Pollmethod
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "state":
		return m.State
	case "title":
		return m.Title
	case "type":
		return m.Type
	case "vote_count":
		return m.VoteCount
	case "voted_ids":
		return m.VotedIDs
	case "votes_raw":
		return m.VotesRaw
	case "votes_signature":
		return m.VotesSignature
	case "votescast":
		return m.Votescast
	case "votesinvalid":
		return m.Votesinvalid
	case "votesvalid":
		return m.Votesvalid
	}

	return nil
}
