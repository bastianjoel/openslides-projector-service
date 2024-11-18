package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

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
	loadedRelations                     map[string]struct{}
	publishedToMeetingsInOrganization   *Organization
	childs                              *Mediafile
	parent                              *Mediafile
	meetingMediafiles                   *MeetingMediafile
}

func (m *Mediafile) CollectionName() string {
	return "mediafile"
}

func (m *Mediafile) PublishedToMeetingsInOrganization() *Organization {
	if _, ok := m.loadedRelations["published_to_meetings_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access PublishedToMeetingsInOrganization relation of Mediafile which was not loaded.")
	}

	return m.publishedToMeetingsInOrganization
}

func (m *Mediafile) Childs() *Mediafile {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of Mediafile which was not loaded.")
	}

	return m.childs
}

func (m *Mediafile) Parent() *Mediafile {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of Mediafile which was not loaded.")
	}

	return m.parent
}

func (m *Mediafile) MeetingMediafiles() *MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafiles relation of Mediafile which was not loaded.")
	}

	return m.meetingMediafiles
}

func (m *Mediafile) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "create_timestamp":
		return m.CreateTimestamp
	case "filename":
		return m.Filename
	case "filesize":
		return m.Filesize
	case "id":
		return m.ID
	case "is_directory":
		return m.IsDirectory
	case "meeting_mediafile_ids":
		return m.MeetingMediafileIDs
	case "mimetype":
		return m.Mimetype
	case "owner_id":
		return m.OwnerID
	case "parent_id":
		return m.ParentID
	case "pdf_information":
		return m.PdfInformation
	case "published_to_meetings_in_organization_id":
		return m.PublishedToMeetingsInOrganizationID
	case "title":
		return m.Title
	case "token":
		return m.Token
	}

	return nil
}

func (m *Mediafile) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["create_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.CreateTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["filename"]; ok {
		err := json.Unmarshal([]byte(val), &m.Filename)
		if err != nil {
			return err
		}
	}

	if val, ok := data["filesize"]; ok {
		err := json.Unmarshal([]byte(val), &m.Filesize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_directory"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsDirectory)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["mimetype"]; ok {
		err := json.Unmarshal([]byte(val), &m.Mimetype)
		if err != nil {
			return err
		}
	}

	if val, ok := data["owner_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OwnerID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ParentID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["pdf_information"]; ok {
		err := json.Unmarshal([]byte(val), &m.PdfInformation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["published_to_meetings_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PublishedToMeetingsInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
		if err != nil {
			return err
		}
	}

	if val, ok := data["token"]; ok {
		err := json.Unmarshal([]byte(val), &m.Token)
		if err != nil {
			return err
		}
	}

	return nil
}
