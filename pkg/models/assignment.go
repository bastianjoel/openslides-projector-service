package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Assignment struct {
	AgendaItemID                  *int    `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	CandidateIDs                  []int   `json:"candidate_ids"`
	DefaultPollDescription        *string `json:"default_poll_description"`
	Description                   *string `json:"description"`
	ID                            int     `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	NumberPollCandidates          *bool   `json:"number_poll_candidates"`
	OpenPosts                     *int    `json:"open_posts"`
	Phase                         *string `json:"phase"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	TagIDs                        []int   `json:"tag_ids"`
	Title                         string  `json:"title"`
	loadedRelations               map[string]struct{}
	candidates                    *AssignmentCandidate
	polls                         *Poll
	projections                   *Projection
	attachmentMeetingMediafiles   *MeetingMediafile
	tags                          *Tag
	agendaItem                    *AgendaItem
	meeting                       *Meeting
	listOfSpeakers                *ListOfSpeakers
}

func (m *Assignment) CollectionName() string {
	return "assignment"
}

func (m *Assignment) Candidates() *AssignmentCandidate {
	if _, ok := m.loadedRelations["candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access Candidates relation of Assignment which was not loaded.")
	}

	return m.candidates
}

func (m *Assignment) Polls() *Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Assignment which was not loaded.")
	}

	return m.polls
}

func (m *Assignment) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Assignment which was not loaded.")
	}

	return m.projections
}

func (m *Assignment) AttachmentMeetingMediafiles() *MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Assignment which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Assignment) Tags() *Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Assignment which was not loaded.")
	}

	return m.tags
}

func (m *Assignment) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Assignment which was not loaded.")
	}

	return m.agendaItem
}

func (m *Assignment) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Assignment which was not loaded.")
	}

	return *m.meeting
}

func (m *Assignment) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Assignment which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Assignment) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "candidate_ids":
		return m.CandidateIDs
	case "default_poll_description":
		return m.DefaultPollDescription
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "number_poll_candidates":
		return m.NumberPollCandidates
	case "open_posts":
		return m.OpenPosts
	case "phase":
		return m.Phase
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "tag_ids":
		return m.TagIDs
	case "title":
		return m.Title
	}

	return nil
}

func (m *Assignment) Update(data map[string]string) error {
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

	if val, ok := data["candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_poll_description"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultPollDescription)
		if err != nil {
			return err
		}
	}

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
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

	if val, ok := data["number_poll_candidates"]; ok {
		err := json.Unmarshal([]byte(val), &m.NumberPollCandidates)
		if err != nil {
			return err
		}
	}

	if val, ok := data["open_posts"]; ok {
		err := json.Unmarshal([]byte(val), &m.OpenPosts)
		if err != nil {
			return err
		}
	}

	if val, ok := data["phase"]; ok {
		err := json.Unmarshal([]byte(val), &m.Phase)
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

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
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
