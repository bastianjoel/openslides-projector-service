package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Projection struct {
	Content            json.RawMessage `json:"content"`
	ContentObjectID    string          `json:"content_object_id"`
	CurrentProjectorID *int            `json:"current_projector_id"`
	HistoryProjectorID *int            `json:"history_projector_id"`
	ID                 int             `json:"id"`
	MeetingID          int             `json:"meeting_id"`
	Options            json.RawMessage `json:"options"`
	PreviewProjectorID *int            `json:"preview_projector_id"`
	Stable             *bool           `json:"stable"`
	Type               *string         `json:"type"`
	Weight             *int            `json:"weight"`
	loadedRelations    map[string]struct{}
	meeting            *Meeting
	previewProjector   *Projector
	currentProjector   *Projector
	historyProjector   *Projector
}

func (m *Projection) CollectionName() string {
	return "projection"
}

func (m *Projection) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Projection which was not loaded.")
	}

	return *m.meeting
}

func (m *Projection) PreviewProjector() *Projector {
	if _, ok := m.loadedRelations["preview_projector_id"]; !ok {
		log.Panic().Msg("Tried to access PreviewProjector relation of Projection which was not loaded.")
	}

	return m.previewProjector
}

func (m *Projection) CurrentProjector() *Projector {
	if _, ok := m.loadedRelations["current_projector_id"]; !ok {
		log.Panic().Msg("Tried to access CurrentProjector relation of Projection which was not loaded.")
	}

	return m.currentProjector
}

func (m *Projection) HistoryProjector() *Projector {
	if _, ok := m.loadedRelations["history_projector_id"]; !ok {
		log.Panic().Msg("Tried to access HistoryProjector relation of Projection which was not loaded.")
	}

	return m.historyProjector
}

func (m *Projection) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "preview_projector_id":
			m.previewProjector = content.(*Projector)
		case "current_projector_id":
			m.currentProjector = content.(*Projector)
		case "history_projector_id":
			m.historyProjector = content.(*Projector)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Projection) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "preview_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.previewProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	case "current_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.currentProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	case "history_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.historyProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *Projection) Get(field string) interface{} {
	switch field {
	case "content":
		return m.Content
	case "content_object_id":
		return m.ContentObjectID
	case "current_projector_id":
		return m.CurrentProjectorID
	case "history_projector_id":
		return m.HistoryProjectorID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "options":
		return m.Options
	case "preview_projector_id":
		return m.PreviewProjectorID
	case "stable":
		return m.Stable
	case "type":
		return m.Type
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *Projection) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "preview_projector_id":
		if m.PreviewProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.PreviewProjectorID)}
		}

	case "current_projector_id":
		if m.CurrentProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.CurrentProjectorID)}
		}

	case "history_projector_id":
		if m.HistoryProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.HistoryProjectorID)}
		}
	}
	return []string{}
}

func (m *Projection) Update(data map[string]string) error {
	if val, ok := data["content"]; ok {
		err := json.Unmarshal([]byte(val), &m.Content)
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

	if val, ok := data["current_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["history_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.HistoryProjectorID)
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

	if val, ok := data["options"]; ok {
		err := json.Unmarshal([]byte(val), &m.Options)
		if err != nil {
			return err
		}
	}

	if val, ok := data["preview_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviewProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["stable"]; ok {
		err := json.Unmarshal([]byte(val), &m.Stable)
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

func (m *Projection) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}
