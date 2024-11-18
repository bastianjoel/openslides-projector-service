package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionEditor struct {
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   int  `json:"meeting_user_id"`
	MotionID        int  `json:"motion_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	motion          *Motion
}

func (m *MotionEditor) CollectionName() string {
	return "motion_editor"
}

func (m *MotionEditor) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionEditor which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionEditor) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of MotionEditor which was not loaded.")
	}

	return *m.meetingUser
}

func (m *MotionEditor) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionEditor which was not loaded.")
	}

	return *m.motion
}

func (m *MotionEditor) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "motion_id":
		return m.MotionID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionEditor) Update(data map[string]string) error {
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

	if val, ok := data["motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
