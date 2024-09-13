package models

import "encoding/json"

type Organization struct {
	ActiveMeetingIDs           []int
	ArchivedMeetingIDs         []int
	CommitteeIDs               []int
	DefaultLanguage            string
	Description                *string
	EnableChat                 *bool
	EnableElectronicVoting     *bool
	Genders                    []string
	ID                         *int
	LegalNotice                *string
	LimitOfMeetings            *int
	LimitOfUsers               *int
	LoginText                  *string
	MediafileIDs               []int
	Name                       *string
	OrganizationTagIDs         []int
	PrivacyPolicy              *string
	PublishedMediafileIDs      []int
	RequireDuplicateFrom       *bool
	ResetPasswordVerboseErrors *bool
	SamlAttrMapping            json.RawMessage
	SamlEnabled                *bool
	SamlLoginButtonText        *string
	SamlMetadataIDp            *string
	SamlMetadataSp             *string
	SamlPrivateKey             *string
	TemplateMeetingIDs         []int
	ThemeID                    int
	ThemeIDs                   []int
	Url                        *string
	UserIDs                    []int
	UsersEmailBody             *string
	UsersEmailReplyto          *string
	UsersEmailSender           *string
	UsersEmailSubject          *string
	VoteDecryptPublicMainKey   *string
}

func (m Organization) CollectionName() string {
	return "organization"
}
