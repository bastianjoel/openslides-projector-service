package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type AssignmentCandidate struct {
	AssignmentID    int  `json:"assignment_id"`
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   *int `json:"meeting_user_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	assignment      *Assignment
}

func (m *AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}

func (m *AssignmentCandidate) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AssignmentCandidate which was not loaded.")
	}

	return *m.meeting
}

func (m *AssignmentCandidate) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of AssignmentCandidate which was not loaded.")
	}

	return m.meetingUser
}

func (m *AssignmentCandidate) Assignment() Assignment {
	if _, ok := m.loadedRelations["assignment_id"]; !ok {
		log.Panic().Msg("Tried to access Assignment relation of AssignmentCandidate which was not loaded.")
	}

	return *m.assignment
}

func (m *AssignmentCandidate) Get(field string) interface{} {
	switch field {
	case "assignment_id":
		return m.AssignmentID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *AssignmentCandidate) Update(data map[string]string) error {
	if val, ok := data["assignment_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentID)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
