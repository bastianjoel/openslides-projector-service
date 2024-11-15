package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type PersonalNote struct {
	ContentObjectID *string `json:"content_object_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	MeetingUserID   int     `json:"meeting_user_id"`
	Note            *string `json:"note"`
	Star            *bool   `json:"star"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
}

func (m PersonalNote) CollectionName() string {
	return "personal_note"
}

func (m *PersonalNote) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PersonalNote which was not loaded.")
	}

	return *m.meeting
}

func (m *PersonalNote) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of PersonalNote which was not loaded.")
	}

	return *m.meetingUser
}

func (m PersonalNote) Get(field string) interface{} {
	switch field {
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "note":
		return m.Note
	case "star":
		return m.Star
	}

	return nil
}

func (m PersonalNote) Update(data map[string]string) error {
	if val, ok := data["content_object_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ContentObjectID)
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["note"]; ok {
		err := json.Unmarshal([]byte(val), &m.Note)
		if err != nil {
			return err
		}
	}

	if val, ok := data["star"]; ok {
		err := json.Unmarshal([]byte(val), &m.Star)
		if err != nil {
			return err
		}
	}

	return nil
}
