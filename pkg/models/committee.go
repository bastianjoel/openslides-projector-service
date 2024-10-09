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
