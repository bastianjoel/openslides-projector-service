package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Vote struct {
	DelegatedUserID *int    `json:"delegated_user_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	OptionID        int     `json:"option_id"`
	UserID          *int    `json:"user_id"`
	UserToken       string  `json:"user_token"`
	Value           *string `json:"value"`
	Weight          *string `json:"weight"`
	loadedRelations map[string]struct{}
	delegatedUser   *User
	meeting         *Meeting
	option          *Option
	user            *User
}

func (m *Vote) CollectionName() string {
	return "vote"
}

func (m *Vote) DelegatedUser() *User {
	if _, ok := m.loadedRelations["delegated_user_id"]; !ok {
		log.Panic().Msg("Tried to access DelegatedUser relation of Vote which was not loaded.")
	}

	return m.delegatedUser
}

func (m *Vote) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Vote which was not loaded.")
	}

	return *m.meeting
}

func (m *Vote) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of Vote which was not loaded.")
	}

	return *m.option
}

func (m *Vote) User() *User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of Vote which was not loaded.")
	}

	return m.user
}

func (m *Vote) Get(field string) interface{} {
	switch field {
	case "delegated_user_id":
		return m.DelegatedUserID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "user_id":
		return m.UserID
	case "user_token":
		return m.UserToken
	case "value":
		return m.Value
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *Vote) Update(data map[string]string) error {
	if val, ok := data["delegated_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DelegatedUserID)
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

	if val, ok := data["option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_token"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserToken)
		if err != nil {
			return err
		}
	}

	if val, ok := data["value"]; ok {
		err := json.Unmarshal([]byte(val), &m.Value)
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

	return nil
}
