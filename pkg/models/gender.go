package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Gender struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	OrganizationID  int    `json:"organization_id"`
	UserIDs         []int  `json:"user_ids"`
	loadedRelations map[string]struct{}
	organization    *Organization
	users           *User
}

func (m *Gender) CollectionName() string {
	return "gender"
}

func (m *Gender) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Gender which was not loaded.")
	}

	return *m.organization
}

func (m *Gender) Users() *User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Gender which was not loaded.")
	}

	return m.users
}

func (m *Gender) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "user_ids":
		return m.UserIDs
	}

	return nil
}

func (m *Gender) Update(data map[string]string) error {
	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
