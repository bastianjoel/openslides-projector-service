package models

import "encoding/json"

type Organization struct {
	ActiveMeetingIDs           []int           `json:"active_meeting_ids"`
	ArchivedMeetingIDs         []int           `json:"archived_meeting_ids"`
	CommitteeIDs               []int           `json:"committee_ids"`
	DefaultLanguage            string          `json:"default_language"`
	Description                *string         `json:"description"`
	EnableAnonymous            *bool           `json:"enable_anonymous"`
	EnableChat                 *bool           `json:"enable_chat"`
	EnableElectronicVoting     *bool           `json:"enable_electronic_voting"`
	GenderIDs                  []int           `json:"gender_ids"`
	ID                         int             `json:"id"`
	LegalNotice                *string         `json:"legal_notice"`
	LimitOfMeetings            *int            `json:"limit_of_meetings"`
	LimitOfUsers               *int            `json:"limit_of_users"`
	LoginText                  *string         `json:"login_text"`
	MediafileIDs               []int           `json:"mediafile_ids"`
	Name                       *string         `json:"name"`
	OrganizationTagIDs         []int           `json:"organization_tag_ids"`
	PrivacyPolicy              *string         `json:"privacy_policy"`
	PublishedMediafileIDs      []int           `json:"published_mediafile_ids"`
	RequireDuplicateFrom       *bool           `json:"require_duplicate_from"`
	ResetPasswordVerboseErrors *bool           `json:"reset_password_verbose_errors"`
	SamlAttrMapping            json.RawMessage `json:"saml_attr_mapping"`
	SamlEnabled                *bool           `json:"saml_enabled"`
	SamlLoginButtonText        *string         `json:"saml_login_button_text"`
	SamlMetadataIDp            *string         `json:"saml_metadata_idp"`
	SamlMetadataSp             *string         `json:"saml_metadata_sp"`
	SamlPrivateKey             *string         `json:"saml_private_key"`
	TemplateMeetingIDs         []int           `json:"template_meeting_ids"`
	ThemeID                    int             `json:"theme_id"`
	ThemeIDs                   []int           `json:"theme_ids"`
	Url                        *string         `json:"url"`
	UserIDs                    []int           `json:"user_ids"`
	UsersEmailBody             *string         `json:"users_email_body"`
	UsersEmailReplyto          *string         `json:"users_email_replyto"`
	UsersEmailSender           *string         `json:"users_email_sender"`
	UsersEmailSubject          *string         `json:"users_email_subject"`
	VoteDecryptPublicMainKey   *string         `json:"vote_decrypt_public_main_key"`
}

func (m Organization) CollectionName() string {
	return "organization"
}

func (m Organization) Get(field string) interface{} {
	switch field {
	case "active_meeting_ids":
		return m.ActiveMeetingIDs
	case "archived_meeting_ids":
		return m.ArchivedMeetingIDs
	case "committee_ids":
		return m.CommitteeIDs
	case "default_language":
		return m.DefaultLanguage
	case "description":
		return m.Description
	case "enable_anonymous":
		return m.EnableAnonymous
	case "enable_chat":
		return m.EnableChat
	case "enable_electronic_voting":
		return m.EnableElectronicVoting
	case "gender_ids":
		return m.GenderIDs
	case "id":
		return m.ID
	case "legal_notice":
		return m.LegalNotice
	case "limit_of_meetings":
		return m.LimitOfMeetings
	case "limit_of_users":
		return m.LimitOfUsers
	case "login_text":
		return m.LoginText
	case "mediafile_ids":
		return m.MediafileIDs
	case "name":
		return m.Name
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "privacy_policy":
		return m.PrivacyPolicy
	case "published_mediafile_ids":
		return m.PublishedMediafileIDs
	case "require_duplicate_from":
		return m.RequireDuplicateFrom
	case "reset_password_verbose_errors":
		return m.ResetPasswordVerboseErrors
	case "saml_attr_mapping":
		return m.SamlAttrMapping
	case "saml_enabled":
		return m.SamlEnabled
	case "saml_login_button_text":
		return m.SamlLoginButtonText
	case "saml_metadata_idp":
		return m.SamlMetadataIDp
	case "saml_metadata_sp":
		return m.SamlMetadataSp
	case "saml_private_key":
		return m.SamlPrivateKey
	case "template_meeting_ids":
		return m.TemplateMeetingIDs
	case "theme_id":
		return m.ThemeID
	case "theme_ids":
		return m.ThemeIDs
	case "url":
		return m.Url
	case "user_ids":
		return m.UserIDs
	case "users_email_body":
		return m.UsersEmailBody
	case "users_email_replyto":
		return m.UsersEmailReplyto
	case "users_email_sender":
		return m.UsersEmailSender
	case "users_email_subject":
		return m.UsersEmailSubject
	case "vote_decrypt_public_main_key":
		return m.VoteDecryptPublicMainKey
	}

	return nil
}
