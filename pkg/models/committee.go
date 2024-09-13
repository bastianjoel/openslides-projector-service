package models

type Committee struct {
	DefaultMeetingID                   *int
	Description                        *string
	ExternalID                         *string
	ForwardToCommitteeIDs              []int
	ForwardingUserID                   *int
	ID                                 *int
	ManagerIDs                         []int
	MeetingIDs                         []int
	Name                               string
	OrganizationID                     int
	OrganizationTagIDs                 []int
	ReceiveForwardingsFromCommitteeIDs []int
	UserIDs                            []int
}

func (m Committee) CollectionName() string {
	return "committee"
}
