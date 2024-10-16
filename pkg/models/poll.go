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

func (m Poll) Update(data map[string]string) error {
	if val, ok := data["backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.Backend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content_object_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ContentObjectID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["crypt_key"]; ok {
		err := json.Unmarshal([]byte(val), &m.CryptKey)
		if err != nil {
			return err
		}
	}

	if val, ok := data["crypt_signature"]; ok {
		err := json.Unmarshal([]byte(val), &m.CryptSignature)
		if err != nil {
			return err
		}
	}

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
		if err != nil {
			return err
		}
	}

	if val, ok := data["entitled_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.EntitledGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["entitled_users_at_stop"]; ok {
		err := json.Unmarshal([]byte(val), &m.EntitledUsersAtStop)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_abstain"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalAbstain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_no"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalNo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalOptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_yes"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalYes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_pseudoanonymized"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPseudoanonymized)
		if err != nil {
			return err
		}
	}

	if val, ok := data["max_votes_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MaxVotesAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["max_votes_per_option"]; ok {
		err := json.Unmarshal([]byte(val), &m.MaxVotesPerOption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["min_votes_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MinVotesAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.OnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["pollmethod"]; ok {
		err := json.Unmarshal([]byte(val), &m.Pollmethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state"]; ok {
		err := json.Unmarshal([]byte(val), &m.State)
		if err != nil {
			return err
		}
	}

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
		if err != nil {
			return err
		}
	}

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_count"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteCount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["voted_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotedIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votes_raw"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotesRaw)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votes_signature"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotesSignature)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votescast"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votescast)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votesinvalid"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votesinvalid)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votesvalid"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votesvalid)
		if err != nil {
			return err
		}
	}

	return nil
}
