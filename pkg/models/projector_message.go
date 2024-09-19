package models

type ProjectorMessage struct {
	ID            *int    `json:"id"`
	MeetingID     int     `json:"meeting_id"`
	Message       *string `json:"message"`
	ProjectionIDs []int   `json:"projection_ids"`
}

func (m ProjectorMessage) CollectionName() string {
	return "projector_message"
}
