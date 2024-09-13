package models

type Committee struct {
	ForwardToCommitteeIDs              []int
	ForwardingUserID                   *int
	ID                                 *int
	MeetingIDs                         []int
	OrganizationID                     int
	UserIDs                            []int
	ReceiveForwardingsFromCommitteeIDs []int
	DefaultMeetingID                   *int
	Description                        *string
	ExternalID                         *string
	ManagerIDs                         []int
	Name                               string
	OrganizationTagIDs                 []int
}

func (m Committee) CollectionName() string {
	return "committee"
}
