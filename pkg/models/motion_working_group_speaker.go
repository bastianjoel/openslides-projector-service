package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionWorkingGroupSpeaker struct {
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

func (m *MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}

func (m *MotionWorkingGroupSpeaker) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.meeting
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

func (m *MotionWorkingGroupSpeaker) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "motion_id":
			m.motion = content.(*Motion)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionWorkingGroupSpeaker) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "meeting_user_id":
		err := json.Unmarshal(content, &m.meetingUser)
		if err != nil {
			return err
		}
	case "motion_id":
		err := json.Unmarshal(content, &m.motion)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return nil
}

func (m *MotionWorkingGroupSpeaker) Get(field string) interface{} {
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

func (m *MotionWorkingGroupSpeaker) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		return []string{"meeting_user/" + strconv.Itoa(m.MeetingUserID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}
	}
	return []string{}
}

func (m *MotionWorkingGroupSpeaker) Update(data map[string]string) error {
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
