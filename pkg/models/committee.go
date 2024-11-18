package models

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	organizationTags                   []OrganizationTag
	receiveForwardingsFromCommittees   []Committee
	users                              []User
	defaultMeeting                     *Meeting
	managers                           []User
	meetings                           []Meeting
	organization                       *Organization
	forwardToCommittees                []Committee
	forwardingUser                     *User
}

func (m *Committee) CollectionName() string {
	return "committee"
}

func (m *Committee) OrganizationTags() []OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Committee which was not loaded.")
	}

	return m.organizationTags
}

func (m *Committee) ReceiveForwardingsFromCommittees() []Committee {
	if _, ok := m.loadedRelations["receive_forwardings_from_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ReceiveForwardingsFromCommittees relation of Committee which was not loaded.")
	}

	return m.receiveForwardingsFromCommittees
}

func (m *Committee) Users() []User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Committee which was not loaded.")
	}

	return m.users
}

func (m *Committee) DefaultMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultMeeting relation of Committee which was not loaded.")
	}

	return m.defaultMeeting
}

func (m *Committee) Managers() []User {
	if _, ok := m.loadedRelations["manager_ids"]; !ok {
		log.Panic().Msg("Tried to access Managers relation of Committee which was not loaded.")
	}

	return m.managers
}

func (m *Committee) Meetings() []Meeting {
	if _, ok := m.loadedRelations["meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access Meetings relation of Committee which was not loaded.")
	}

	return m.meetings
}

func (m *Committee) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Committee which was not loaded.")
	}

	return *m.organization
}

func (m *Committee) ForwardToCommittees() []Committee {
	if _, ok := m.loadedRelations["forward_to_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardToCommittees relation of Committee which was not loaded.")
	}

	return m.forwardToCommittees
}

func (m *Committee) ForwardingUser() *User {
	if _, ok := m.loadedRelations["forwarding_user_id"]; !ok {
		log.Panic().Msg("Tried to access ForwardingUser relation of Committee which was not loaded.")
	}

	return m.forwardingUser
}

func (m *Committee) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "organization_tag_ids":
			m.organizationTags = content.([]OrganizationTag)
		case "receive_forwardings_from_committee_ids":
			m.receiveForwardingsFromCommittees = content.([]Committee)
		case "user_ids":
			m.users = content.([]User)
		case "default_meeting_id":
			m.defaultMeeting = content.(*Meeting)
		case "manager_ids":
			m.managers = content.([]User)
		case "meeting_ids":
			m.meetings = content.([]Meeting)
		case "organization_id":
			m.organization = content.(*Organization)
		case "forward_to_committee_ids":
			m.forwardToCommittees = content.([]Committee)
		case "forwarding_user_id":
			m.forwardingUser = content.(*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Committee) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "organization_tag_ids":
		err := json.Unmarshal(content, &m.organizationTags)
		if err != nil {
			return err
		}
	case "receive_forwardings_from_committee_ids":
		err := json.Unmarshal(content, &m.receiveForwardingsFromCommittees)
		if err != nil {
			return err
		}
	case "user_ids":
		err := json.Unmarshal(content, &m.users)
		if err != nil {
			return err
		}
	case "default_meeting_id":
		err := json.Unmarshal(content, &m.defaultMeeting)
		if err != nil {
			return err
		}
	case "manager_ids":
		err := json.Unmarshal(content, &m.managers)
		if err != nil {
			return err
		}
	case "meeting_ids":
		err := json.Unmarshal(content, &m.meetings)
		if err != nil {
			return err
		}
	case "organization_id":
		err := json.Unmarshal(content, &m.organization)
		if err != nil {
			return err
		}
	case "forward_to_committee_ids":
		err := json.Unmarshal(content, &m.forwardToCommittees)
		if err != nil {
			return err
		}
	case "forwarding_user_id":
		err := json.Unmarshal(content, &m.forwardingUser)
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

func (m *Committee) Get(field string) interface{} {
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

func (m *Committee) GetFqids(field string) []string {
	switch field {
	case "organization_tag_ids":
		r := make([]string, len(m.OrganizationTagIDs))
		for i, id := range m.OrganizationTagIDs {
			r[i] = "organization_tag/" + strconv.Itoa(id)
		}
		return r

	case "receive_forwardings_from_committee_ids":
		r := make([]string, len(m.ReceiveForwardingsFromCommitteeIDs))
		for i, id := range m.ReceiveForwardingsFromCommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "user_ids":
		r := make([]string, len(m.UserIDs))
		for i, id := range m.UserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "default_meeting_id":
		if m.DefaultMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultMeetingID)}
		}

	case "manager_ids":
		r := make([]string, len(m.ManagerIDs))
		for i, id := range m.ManagerIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "meeting_ids":
		r := make([]string, len(m.MeetingIDs))
		for i, id := range m.MeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}

	case "forward_to_committee_ids":
		r := make([]string, len(m.ForwardToCommitteeIDs))
		for i, id := range m.ForwardToCommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "forwarding_user_id":
		if m.ForwardingUserID != nil {
			return []string{"user/" + strconv.Itoa(*m.ForwardingUserID)}
		}
	}
	return []string{}
}

func (m *Committee) Update(data map[string]string) error {
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
