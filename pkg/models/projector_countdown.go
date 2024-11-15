package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type ProjectorCountdown struct {
	CountdownTime                          *float32 `json:"countdown_time"`
	DefaultTime                            *int     `json:"default_time"`
	Description                            *string  `json:"description"`
	ID                                     int      `json:"id"`
	MeetingID                              int      `json:"meeting_id"`
	ProjectionIDs                          []int    `json:"projection_ids"`
	Running                                *bool    `json:"running"`
	Title                                  string   `json:"title"`
	UsedAsListOfSpeakersCountdownMeetingID *int     `json:"used_as_list_of_speakers_countdown_meeting_id"`
	UsedAsPollCountdownMeetingID           *int     `json:"used_as_poll_countdown_meeting_id"`
	loadedRelations                        map[string]struct{}
	meeting                                *Meeting
	projections                            *Projection
	usedAsListOfSpeakersCountdownMeeting   *Meeting
	usedAsPollCountdownMeeting             *Meeting
}

func (m ProjectorCountdown) CollectionName() string {
	return "projector_countdown"
}

func (m *ProjectorCountdown) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ProjectorCountdown which was not loaded.")
	}

	return *m.meeting
}

func (m *ProjectorCountdown) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ProjectorCountdown which was not loaded.")
	}

	return m.projections
}

func (m *ProjectorCountdown) UsedAsListOfSpeakersCountdownMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_list_of_speakers_countdown_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsListOfSpeakersCountdownMeeting relation of ProjectorCountdown which was not loaded.")
	}

	return m.usedAsListOfSpeakersCountdownMeeting
}

func (m *ProjectorCountdown) UsedAsPollCountdownMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_poll_countdown_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsPollCountdownMeeting relation of ProjectorCountdown which was not loaded.")
	}

	return m.usedAsPollCountdownMeeting
}

func (m ProjectorCountdown) Get(field string) interface{} {
	switch field {
	case "countdown_time":
		return m.CountdownTime
	case "default_time":
		return m.DefaultTime
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "running":
		return m.Running
	case "title":
		return m.Title
	case "used_as_list_of_speakers_countdown_meeting_id":
		return m.UsedAsListOfSpeakersCountdownMeetingID
	case "used_as_poll_countdown_meeting_id":
		return m.UsedAsPollCountdownMeetingID
	}

	return nil
}

func (m ProjectorCountdown) Update(data map[string]string) error {
	if val, ok := data["countdown_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CountdownTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultTime)
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

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
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

	if val, ok := data["running"]; ok {
		err := json.Unmarshal([]byte(val), &m.Running)
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

	if val, ok := data["used_as_list_of_speakers_countdown_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsListOfSpeakersCountdownMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_poll_countdown_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsPollCountdownMeetingID)
		if err != nil {
			return err
		}
	}

	return nil
}
