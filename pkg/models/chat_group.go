package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ChatGroup struct {
	ChatMessageIDs  []int  `json:"chat_message_ids"`
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	Name            string `json:"name"`
	ReadGroupIDs    []int  `json:"read_group_ids"`
	Weight          *int   `json:"weight"`
	WriteGroupIDs   []int  `json:"write_group_ids"`
	loadedRelations map[string]struct{}
	readGroups      []Group
	writeGroups     []Group
	chatMessages    []ChatMessage
	meeting         *Meeting
}

func (m *ChatGroup) CollectionName() string {
	return "chat_group"
}

func (m *ChatGroup) ReadGroups() []Group {
	if _, ok := m.loadedRelations["read_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadGroups relation of ChatGroup which was not loaded.")
	}

	return m.readGroups
}

func (m *ChatGroup) WriteGroups() []Group {
	if _, ok := m.loadedRelations["write_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteGroups relation of ChatGroup which was not loaded.")
	}

	return m.writeGroups
}

func (m *ChatGroup) ChatMessages() []ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of ChatGroup which was not loaded.")
	}

	return m.chatMessages
}

func (m *ChatGroup) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ChatGroup which was not loaded.")
	}

	return *m.meeting
}

func (m *ChatGroup) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "read_group_ids":
			m.readGroups = content.([]Group)
		case "write_group_ids":
			m.writeGroups = content.([]Group)
		case "chat_message_ids":
			m.chatMessages = content.([]ChatMessage)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *ChatGroup) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "read_group_ids":
		err := json.Unmarshal(content, &m.readGroups)
		if err != nil {
			return err
		}
	case "write_group_ids":
		err := json.Unmarshal(content, &m.writeGroups)
		if err != nil {
			return err
		}
	case "chat_message_ids":
		err := json.Unmarshal(content, &m.chatMessages)
		if err != nil {
			return err
		}
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
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

func (m *ChatGroup) Get(field string) interface{} {
	switch field {
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "read_group_ids":
		return m.ReadGroupIDs
	case "weight":
		return m.Weight
	case "write_group_ids":
		return m.WriteGroupIDs
	}

	return nil
}

func (m *ChatGroup) GetFqids(field string) []string {
	switch field {
	case "read_group_ids":
		r := make([]string, len(m.ReadGroupIDs))
		for i, id := range m.ReadGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "write_group_ids":
		r := make([]string, len(m.WriteGroupIDs))
		for i, id := range m.WriteGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "chat_message_ids":
		r := make([]string, len(m.ChatMessageIDs))
		for i, id := range m.ChatMessageIDs {
			r[i] = "chat_message/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}
	}
	return []string{}
}

func (m *ChatGroup) Update(data map[string]string) error {
	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["read_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadGroupIDs)
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

	if val, ok := data["write_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteGroupIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
