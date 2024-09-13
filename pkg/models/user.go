package models

type User struct {
	IsPhysicalPerson            *bool
	IsPresentInMeetingIDs       []int
	LastEmailSent               *int
	LastLogin                   *int
	Username                    string
	IsDemoUser                  *bool
	LastName                    *string
	SamlID                      *string
	VoteIDs                     []int
	OptionIDs                   []int
	PollCandidateIDs            []int
	DefaultVoteWeight           *string
	ForwardingCommitteeIDs      []int
	Gender                      *string
	ID                          *int
	MeetingIDs                  []int
	MeetingUserIDs              []int
	PollVotedIDs                []int
	Email                       *string
	Password                    *string
	DefaultPassword             *string
	OrganizationID              int
	OrganizationManagementLevel *string
	CanChangeOwnPassword        *bool
	CommitteeManagementIDs      []int
	DelegatedVoteIDs            []int
	IsActive                    *bool
	MemberNumber                *string
	Pronoun                     *string
	CommitteeIDs                []int
	FirstName                   *string
	Title                       *string
}

func (m User) CollectionName() string {
	return "user"
}
