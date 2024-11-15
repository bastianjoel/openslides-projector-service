package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionWorkingGroupSpeaker struct {
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   int  `json:"meeting_user_id"`
	MotionID        int  `json:"motion_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meetingUser     *MeetingUser
	motion          *Motion
	meeting         *Meeting
}

func (m MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}

func (m *MotionWorkingGroupSpeaker) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.meetingUser
}

func (m *MotionWorkingGroupSpeaker) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.motion
}

func (m *MotionWorkingGroupSpeaker) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.meeting
}

func (m MotionWorkingGroupSpeaker) Get(field string) interface{} {
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

func (m MotionWorkingGroupSpeaker) Update(data map[string]string) error {
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
