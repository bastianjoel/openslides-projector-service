package models

type User struct {
	CanChangeOwnPassword        *bool
	CommitteeIDs                []int
	CommitteeManagementIDs      []int
	DefaultPassword             *string
	DefaultVoteWeight           *string
	DelegatedVoteIDs            []int
	Email                       *string
	FirstName                   *string
	ForwardingCommitteeIDs      []int
	Gender                      *string
	ID                          *int
	IsActive                    *bool
	IsDemoUser                  *bool
	IsPhysicalPerson            *bool
	IsPresentInMeetingIDs       []int
	LastEmailSent               *int
	LastLogin                   *int
	LastName                    *string
	MeetingIDs                  []int
	MeetingUserIDs              []int
	MemberNumber                *string
	OptionIDs                   []int
	OrganizationID              int
	OrganizationManagementLevel *string
	Password                    *string
	PollCandidateIDs            []int
	PollVotedIDs                []int
	Pronoun                     *string
	SamlID                      *string
	Title                       *string
	Username                    string
	VoteIDs                     []int
}

func (m User) CollectionName() string {
	return "user"
}
