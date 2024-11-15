package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type AgendaItem struct {
	ChildIDs        []int   `json:"child_ids"`
	Closed          *bool   `json:"closed"`
	Comment         *string `json:"comment"`
	ContentObjectID string  `json:"content_object_id"`
	Duration        *int    `json:"duration"`
	ID              int     `json:"id"`
	IsHidden        *bool   `json:"is_hidden"`
	IsInternal      *bool   `json:"is_internal"`
	ItemNumber      *string `json:"item_number"`
	Level           *int    `json:"level"`
	MeetingID       int     `json:"meeting_id"`
	ModeratorNotes  *string `json:"moderator_notes"`
	ParentID        *int    `json:"parent_id"`
	ProjectionIDs   []int   `json:"projection_ids"`
	TagIDs          []int   `json:"tag_ids"`
	Type            *string `json:"type"`
	Weight          *int    `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	childs          *AgendaItem
	parent          *AgendaItem
	projections     *Projection
	tags            *Tag
}

func (m AgendaItem) CollectionName() string {
	return "agenda_item"
}

func (m *AgendaItem) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AgendaItem which was not loaded.")
	}

	return *m.meeting
}

func (m *AgendaItem) Childs() *AgendaItem {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of AgendaItem which was not loaded.")
	}

	return m.childs
}

func (m *AgendaItem) Parent() *AgendaItem {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of AgendaItem which was not loaded.")
	}

	return m.parent
}

func (m *AgendaItem) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of AgendaItem which was not loaded.")
	}

	return m.projections
}

func (m *AgendaItem) Tags() *Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of AgendaItem which was not loaded.")
	}

	return m.tags
}

func (m AgendaItem) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "closed":
		return m.Closed
	case "comment":
		return m.Comment
	case "content_object_id":
		return m.ContentObjectID
	case "duration":
		return m.Duration
	case "id":
		return m.ID
	case "is_hidden":
		return m.IsHidden
	case "is_internal":
		return m.IsInternal
	case "item_number":
		return m.ItemNumber
	case "level":
		return m.Level
	case "meeting_id":
		return m.MeetingID
	case "moderator_notes":
		return m.ModeratorNotes
	case "parent_id":
		return m.ParentID
	case "projection_ids":
		return m.ProjectionIDs
	case "tag_ids":
		return m.TagIDs
	case "type":
		return m.Type
	case "weight":
		return m.Weight
	}

	return nil
}

func (m AgendaItem) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.Closed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content_object_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ContentObjectID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["duration"]; ok {
		err := json.Unmarshal([]byte(val), &m.Duration)
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

	if val, ok := data["is_hidden"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsHidden)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["item_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.ItemNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["level"]; ok {
		err := json.Unmarshal([]byte(val), &m.Level)
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

	if val, ok := data["moderator_notes"]; ok {
		err := json.Unmarshal([]byte(val), &m.ModeratorNotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ParentID)
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

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
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
