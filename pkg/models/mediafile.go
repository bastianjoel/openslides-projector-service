package models

import "encoding/json"

type Mediafile struct {
	OwnerID                             string
	ParentID                            *int
	IsDirectory                         *bool
	MeetingMediafileIDs                 []int
	Mimetype                            *string
	Token                               *string
	ChildIDs                            []int
	Filename                            *string
	PdfInformation                      json.RawMessage
	Filesize                            *int
	Title                               *string
	CreateTimestamp                     *int
	ID                                  *int
	PublishedToMeetingsInOrganizationID *int
}

func (m Mediafile) CollectionName() string {
	return "mediafile"
}
