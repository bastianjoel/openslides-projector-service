package models

import "encoding/json"

type Mediafile struct {
	ChildIDs                            []int
	CreateTimestamp                     *int
	Filename                            *string
	Filesize                            *int
	ID                                  *int
	IsDirectory                         *bool
	MeetingMediafileIDs                 []int
	Mimetype                            *string
	OwnerID                             string
	ParentID                            *int
	PdfInformation                      json.RawMessage
	PublishedToMeetingsInOrganizationID *int
	Title                               *string
	Token                               *string
}

func (m Mediafile) CollectionName() string {
	return "mediafile"
}
