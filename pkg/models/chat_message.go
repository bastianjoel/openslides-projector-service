package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type ChatMessage struct {
	ChatGroupID     int    `json:"chat_group_id"`
	Content         string `json:"content"`
	Created         int    `json:"created"`
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	MeetingUserID   *int   `json:"meeting_user_id"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	chatGroup       *ChatGroup
}

func (m ChatMessage) CollectionName() string {
	return "chat_message"
}

func (m *ChatMessage) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ChatMessage which was not loaded.")
	}

	return *m.meeting
}

func (m *ChatMessage) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of ChatMessage which was not loaded.")
	}

	return m.meetingUser
}

func (m *ChatMessage) ChatGroup() ChatGroup {
	if _, ok := m.loadedRelations["chat_group_id"]; !ok {
		log.Panic().Msg("Tried to access ChatGroup relation of ChatMessage which was not loaded.")
	}

	return *m.chatGroup
}

func (m ChatMessage) Get(field string) interface{} {
	switch field {
	case "chat_group_id":
		return m.ChatGroupID
	case "content":
		return m.Content
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	}

	return nil
}

func (m ChatMessage) Update(data map[string]string) error {
	if val, ok := data["chat_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content"]; ok {
		err := json.Unmarshal([]byte(val), &m.Content)
		if err != nil {
			return err
		}
	}

	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
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

	return nil
}
