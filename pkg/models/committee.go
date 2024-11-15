package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Committee struct {
	DefaultMeetingID                   *int    `json:"default_meeting_id"`
	Description                        *string `json:"description"`
	ExternalID                         *string `json:"external_id"`
	ForwardToCommitteeIDs              []int   `json:"forward_to_committee_ids"`
	ForwardingUserID                   *int    `json:"forwarding_user_id"`
	ID                                 int     `json:"id"`
	ManagerIDs                         []int   `json:"manager_ids"`
	MeetingIDs                         []int   `json:"meeting_ids"`
	Name                               string  `json:"name"`
	OrganizationID                     int     `json:"organization_id"`
	OrganizationTagIDs                 []int   `json:"organization_tag_ids"`
	ReceiveForwardingsFromCommitteeIDs []int   `json:"receive_forwardings_from_committee_ids"`
	UserIDs                            []int   `json:"user_ids"`
	loadedRelations                    map[string]struct{}
	defaultMeeting                     *Meeting
	forwardToCommittees                *Committee
	managers                           *User
	meetings                           *Meeting
	users                              *User
	forwardingUser                     *User
	organization                       *Organization
	organizationTags                   *OrganizationTag
	receiveForwardingsFromCommittees   *Committee
}

func (m Committee) CollectionName() string {
	return "committee"
}

func (m *Committee) DefaultMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultMeeting relation of Committee which was not loaded.")
	}

	return m.defaultMeeting
}

func (m *Committee) ForwardToCommittees() *Committee {
	if _, ok := m.loadedRelations["forward_to_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardToCommittees relation of Committee which was not loaded.")
	}

	return m.forwardToCommittees
}

func (m *Committee) Managers() *User {
	if _, ok := m.loadedRelations["manager_ids"]; !ok {
		log.Panic().Msg("Tried to access Managers relation of Committee which was not loaded.")
	}

	return m.managers
}

func (m *Committee) Meetings() *Meeting {
	if _, ok := m.loadedRelations["meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access Meetings relation of Committee which was not loaded.")
	}

	return m.meetings
}

func (m *Committee) Users() *User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Committee which was not loaded.")
	}

	return m.users
}

func (m *Committee) ForwardingUser() *User {
	if _, ok := m.loadedRelations["forwarding_user_id"]; !ok {
		log.Panic().Msg("Tried to access ForwardingUser relation of Committee which was not loaded.")
	}

	return m.forwardingUser
}

func (m *Committee) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Committee which was not loaded.")
	}

	return *m.organization
}

func (m *Committee) OrganizationTags() *OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Committee which was not loaded.")
	}

	return m.organizationTags
}

func (m *Committee) ReceiveForwardingsFromCommittees() *Committee {
	if _, ok := m.loadedRelations["receive_forwardings_from_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ReceiveForwardingsFromCommittees relation of Committee which was not loaded.")
	}

	return m.receiveForwardingsFromCommittees
}

func (m Committee) Get(field string) interface{} {
	switch field {
	case "default_meeting_id":
		return m.DefaultMeetingID
	case "description":
		return m.Description
	case "external_id":
		return m.ExternalID
	case "forward_to_committee_ids":
		return m.ForwardToCommitteeIDs
	case "forwarding_user_id":
		return m.ForwardingUserID
	case "id":
		return m.ID
	case "manager_ids":
		return m.ManagerIDs
	case "meeting_ids":
		return m.MeetingIDs
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "receive_forwardings_from_committee_ids":
		return m.ReceiveForwardingsFromCommitteeIDs
	case "user_ids":
		return m.UserIDs
	}

	return nil
}

func (m Committee) Update(data map[string]string) error {
	if val, ok := data["default_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultMeetingID)
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

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forward_to_committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardToCommitteeIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarding_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardingUserID)
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

	if val, ok := data["manager_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ManagerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingIDs)
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

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["receive_forwardings_from_committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReceiveForwardingsFromCommitteeIDs)
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
