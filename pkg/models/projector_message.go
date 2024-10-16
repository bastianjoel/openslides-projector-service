package models

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
