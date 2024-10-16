package models

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
}

func (m Committee) CollectionName() string {
	return "committee"
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
