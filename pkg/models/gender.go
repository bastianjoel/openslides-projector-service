package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Gender struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	OrganizationID  int    `json:"organization_id"`
	UserIDs         []int  `json:"user_ids"`
	loadedRelations map[string]struct{}
	users           []User
	organization    *Organization
}

func (m *Gender) CollectionName() string {
	return "gender"
}

func (m *Gender) Users() []User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Gender which was not loaded.")
	}

	return m.users
}

func (m *Gender) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Gender which was not loaded.")
	}

	return *m.organization
}

func (m *Gender) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "user_ids":
			m.users = content.([]User)
		case "organization_id":
			m.organization = content.(*Organization)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Gender) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "user_ids":
		err := json.Unmarshal(content, &m.users)
		if err != nil {
			return err
		}
	case "organization_id":
		err := json.Unmarshal(content, &m.organization)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return nil
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

func (m *Gender) GetFqids(field string) []string {
	switch field {
	case "user_ids":
		r := make([]string, len(m.UserIDs))
		for i, id := range m.UserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}
	}
	return []string{}
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
