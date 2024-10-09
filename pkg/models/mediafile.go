package models

import "encoding/json"

type Mediafile struct {
	ChildIDs                            []int           `json:"child_ids"`
	CreateTimestamp                     *int            `json:"create_timestamp"`
	Filename                            *string         `json:"filename"`
	Filesize                            *int            `json:"filesize"`
	ID                                  int             `json:"id"`
	IsDirectory                         *bool           `json:"is_directory"`
	MeetingMediafileIDs                 []int           `json:"meeting_mediafile_ids"`
	Mimetype                            *string         `json:"mimetype"`
	OwnerID                             string          `json:"owner_id"`
	ParentID                            *int            `json:"parent_id"`
	PdfInformation                      json.RawMessage `json:"pdf_information"`
	PublishedToMeetingsInOrganizationID *int            `json:"published_to_meetings_in_organization_id"`
	Title                               *string         `json:"title"`
	Token                               *string         `json:"token"`
}

func (m Mediafile) CollectionName() string {
	return "mediafile"
}
