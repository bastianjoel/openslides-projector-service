package models

import "encoding/json"

type Organization struct {
	DefaultLanguage            string
	Genders                    []string
	PublishedMediafileIDs      []int
	SamlAttrMapping            json.RawMessage
	ThemeIDs                   []int
	UsersEmailBody             *string
	EnableChat                 *bool
	RequireDuplicateFrom       *bool
	SamlEnabled                *bool
	SamlMetadataSp             *string
	TemplateMeetingIDs         []int
	UsersEmailReplyto          *string
	VoteDecryptPublicMainKey   *string
	ActiveMeetingIDs           []int
	CommitteeIDs               []int
	SamlPrivateKey             *string
	Url                        *string
	ArchivedMeetingIDs         []int
	SamlMetadataIDp            *string
	UserIDs                    []int
	UsersEmailSender           *string
	UsersEmailSubject          *string
	LimitOfMeetings            *int
	ID                         *int
	LimitOfUsers               *int
	LoginText                  *string
	Name                       *string
	ResetPasswordVerboseErrors *bool
	Description                *string
	EnableElectronicVoting     *bool
	LegalNotice                *string
	OrganizationTagIDs         []int
	PrivacyPolicy              *string
	SamlLoginButtonText        *string
	MediafileIDs               []int
	ThemeID                    int
}

func (m Organization) CollectionName() string {
	return "organization"
}
