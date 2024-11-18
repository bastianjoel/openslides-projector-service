package models

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	parent          *AgendaItem
	projections     []Projection
	tags            []Tag
	childs          []AgendaItem
}

func (m *AgendaItem) CollectionName() string {
	return "agenda_item"
}

func (m *AgendaItem) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AgendaItem which was not loaded.")
	}

	return *m.meeting
}

func (m *AgendaItem) Parent() *AgendaItem {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of AgendaItem which was not loaded.")
	}

	return m.parent
}

func (m *AgendaItem) Projections() []Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of AgendaItem which was not loaded.")
	}

	return m.projections
}

func (m *AgendaItem) Tags() []Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of AgendaItem which was not loaded.")
	}

	return m.tags
}

func (m *AgendaItem) Childs() []AgendaItem {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of AgendaItem which was not loaded.")
	}

	return m.childs
}

func (m *AgendaItem) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "parent_id":
			m.parent = content.(*AgendaItem)
		case "projection_ids":
			m.projections = content.([]Projection)
		case "tag_ids":
			m.tags = content.([]Tag)
		case "child_ids":
			m.childs = content.([]AgendaItem)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *AgendaItem) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "parent_id":
		err := json.Unmarshal(content, &m.parent)
		if err != nil {
			return err
		}
	case "projection_ids":
		err := json.Unmarshal(content, &m.projections)
		if err != nil {
			return err
		}
	case "tag_ids":
		err := json.Unmarshal(content, &m.tags)
		if err != nil {
			return err
		}
	case "child_ids":
		err := json.Unmarshal(content, &m.childs)
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

func (m *AgendaItem) Get(field string) interface{} {
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

func (m *AgendaItem) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "parent_id":
		if m.ParentID != nil {
			return []string{"agenda_item/" + strconv.Itoa(*m.ParentID)}
		}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r

	case "child_ids":
		r := make([]string, len(m.ChildIDs))
		for i, id := range m.ChildIDs {
			r[i] = "agenda_item/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *AgendaItem) Update(data map[string]string) error {
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
