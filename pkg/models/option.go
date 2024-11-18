package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Option struct {
	Abstain                    *string `json:"abstain"`
	ContentObjectID            *string `json:"content_object_id"`
	ID                         int     `json:"id"`
	MeetingID                  int     `json:"meeting_id"`
	No                         *string `json:"no"`
	PollID                     *int    `json:"poll_id"`
	Text                       *string `json:"text"`
	UsedAsGlobalOptionInPollID *int    `json:"used_as_global_option_in_poll_id"`
	VoteIDs                    []int   `json:"vote_ids"`
	Weight                     *int    `json:"weight"`
	Yes                        *string `json:"yes"`
	loadedRelations            map[string]struct{}
	meeting                    *Meeting
	poll                       *Poll
	usedAsGlobalOptionInPoll   *Poll
	votes                      *Vote
}

func (m *Option) CollectionName() string {
	return "option"
}

func (m *Option) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Option which was not loaded.")
	}

	return *m.meeting
}

func (m *Option) Poll() *Poll {
	if _, ok := m.loadedRelations["poll_id"]; !ok {
		log.Panic().Msg("Tried to access Poll relation of Option which was not loaded.")
	}

	return m.poll
}

func (m *Option) UsedAsGlobalOptionInPoll() *Poll {
	if _, ok := m.loadedRelations["used_as_global_option_in_poll_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsGlobalOptionInPoll relation of Option which was not loaded.")
	}

	return m.usedAsGlobalOptionInPoll
}

func (m *Option) Votes() *Vote {
	if _, ok := m.loadedRelations["vote_ids"]; !ok {
		log.Panic().Msg("Tried to access Votes relation of Option which was not loaded.")
	}

	return m.votes
}

func (m *Option) Get(field string) interface{} {
	switch field {
	case "abstain":
		return m.Abstain
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "no":
		return m.No
	case "poll_id":
		return m.PollID
	case "text":
		return m.Text
	case "used_as_global_option_in_poll_id":
		return m.UsedAsGlobalOptionInPollID
	case "vote_ids":
		return m.VoteIDs
	case "weight":
		return m.Weight
	case "yes":
		return m.Yes
	}

	return nil
}

func (m *Option) Update(data map[string]string) error {
	if val, ok := data["abstain"]; ok {
		err := json.Unmarshal([]byte(val), &m.Abstain)
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

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
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

	if val, ok := data["no"]; ok {
		err := json.Unmarshal([]byte(val), &m.No)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_global_option_in_poll_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsGlobalOptionInPollID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["yes"]; ok {
		err := json.Unmarshal([]byte(val), &m.Yes)
		if err != nil {
			return err
		}
	}

	return nil
}
