package models

import "encoding/json"

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
}

func (m Projection) CollectionName() string {
	return "projection"
}

func (m Projection) Get(field string) interface{} {
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
