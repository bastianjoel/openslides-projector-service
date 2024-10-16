package models

import "encoding/json"

type ProjectorMessage struct {
	ID            int     `json:"id"`
	MeetingID     int     `json:"meeting_id"`
	Message       *string `json:"message"`
	ProjectionIDs []int   `json:"projection_ids"`
}

func (m ProjectorMessage) CollectionName() string {
	return "projector_message"
}

func (m ProjectorMessage) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "message":
		return m.Message
	case "projection_ids":
		return m.ProjectionIDs
	}

	return nil
}

func (m ProjectorMessage) Update(data map[string]string) error {
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

	if val, ok := data["message"]; ok {
		err := json.Unmarshal([]byte(val), &m.Message)
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

	return nil
}
