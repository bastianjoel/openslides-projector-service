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

func (m Projection) Update(data map[string]string) error {
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
