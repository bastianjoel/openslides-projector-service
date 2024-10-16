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

func (m User) Get(field string) interface{} {
	switch field {
	case "can_change_own_password":
		return m.CanChangeOwnPassword
	case "committee_ids":
		return m.CommitteeIDs
	case "committee_management_ids":
		return m.CommitteeManagementIDs
	case "default_password":
		return m.DefaultPassword
	case "default_vote_weight":
		return m.DefaultVoteWeight
	case "delegated_vote_ids":
		return m.DelegatedVoteIDs
	case "email":
		return m.Email
	case "first_name":
		return m.FirstName
	case "forwarding_committee_ids":
		return m.ForwardingCommitteeIDs
	case "gender_id":
		return m.GenderID
	case "id":
		return m.ID
	case "is_active":
		return m.IsActive
	case "is_demo_user":
		return m.IsDemoUser
	case "is_physical_person":
		return m.IsPhysicalPerson
	case "is_present_in_meeting_ids":
		return m.IsPresentInMeetingIDs
	case "last_email_sent":
		return m.LastEmailSent
	case "last_login":
		return m.LastLogin
	case "last_name":
		return m.LastName
	case "meeting_ids":
		return m.MeetingIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "member_number":
		return m.MemberNumber
	case "option_ids":
		return m.OptionIDs
	case "organization_id":
		return m.OrganizationID
	case "organization_management_level":
		return m.OrganizationManagementLevel
	case "password":
		return m.Password
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	case "poll_voted_ids":
		return m.PollVotedIDs
	case "pronoun":
		return m.Pronoun
	case "saml_id":
		return m.SamlID
	case "title":
		return m.Title
	case "username":
		return m.Username
	case "vote_ids":
		return m.VoteIDs
	}

	return nil
}
