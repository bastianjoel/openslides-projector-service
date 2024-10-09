package models

type User struct {
	CanChangeOwnPassword        *bool   `json:"can_change_own_password"`
	CommitteeIDs                []int   `json:"committee_ids"`
	CommitteeManagementIDs      []int   `json:"committee_management_ids"`
	DefaultPassword             *string `json:"default_password"`
	DefaultVoteWeight           *string `json:"default_vote_weight"`
	DelegatedVoteIDs            []int   `json:"delegated_vote_ids"`
	Email                       *string `json:"email"`
	FirstName                   *string `json:"first_name"`
	ForwardingCommitteeIDs      []int   `json:"forwarding_committee_ids"`
	GenderID                    *int    `json:"gender_id"`
	ID                          int     `json:"id"`
	IsActive                    *bool   `json:"is_active"`
	IsDemoUser                  *bool   `json:"is_demo_user"`
	IsPhysicalPerson            *bool   `json:"is_physical_person"`
	IsPresentInMeetingIDs       []int   `json:"is_present_in_meeting_ids"`
	LastEmailSent               *int    `json:"last_email_sent"`
	LastLogin                   *int    `json:"last_login"`
	LastName                    *string `json:"last_name"`
	MeetingIDs                  []int   `json:"meeting_ids"`
	MeetingUserIDs              []int   `json:"meeting_user_ids"`
	MemberNumber                *string `json:"member_number"`
	OptionIDs                   []int   `json:"option_ids"`
	OrganizationID              int     `json:"organization_id"`
	OrganizationManagementLevel *string `json:"organization_management_level"`
	Password                    *string `json:"password"`
	PollCandidateIDs            []int   `json:"poll_candidate_ids"`
	PollVotedIDs                []int   `json:"poll_voted_ids"`
	Pronoun                     *string `json:"pronoun"`
	SamlID                      *string `json:"saml_id"`
	Title                       *string `json:"title"`
	Username                    string  `json:"username"`
	VoteIDs                     []int   `json:"vote_ids"`
}

func (m User) CollectionName() string {
	return "user"
}
