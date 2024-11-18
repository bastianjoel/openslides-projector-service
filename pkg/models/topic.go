package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Topic struct {
	AgendaItemID                  int     `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	ID                            int     `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	Text                          *string `json:"text"`
	Title                         string  `json:"title"`
	loadedRelations               map[string]struct{}
	agendaItem                    *AgendaItem
	projections                   *Projection
	meeting                       *Meeting
	polls                         *Poll
	attachmentMeetingMediafiles   *MeetingMediafile
	listOfSpeakers                *ListOfSpeakers
}

func (m *Topic) CollectionName() string {
	return "topic"
}

func (m *Topic) AgendaItem() AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Topic which was not loaded.")
	}

	return *m.agendaItem
}

func (m *Topic) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Topic which was not loaded.")
	}

	return m.projections
}

func (m *Topic) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Topic which was not loaded.")
	}

	return *m.meeting
}

func (m *Topic) Polls() *Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Topic which was not loaded.")
	}

	return m.polls
}

func (m *Topic) AttachmentMeetingMediafiles() *MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Topic which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Topic) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Topic which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Topic) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "text":
		return m.Text
	case "title":
		return m.Title
	}

	return nil
}

func (m *Topic) Update(data map[string]string) error {
	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["attachment_meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AttachmentMeetingMediafileIDs)
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

	if val, ok := data["list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
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

	return nil
}
