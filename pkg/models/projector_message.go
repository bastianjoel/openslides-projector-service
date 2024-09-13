package models

type ProjectorMessage struct {
	ProjectionIDs []int
	ID            *int
	MeetingID     int
	Message       *string
}

func (m ProjectorMessage) CollectionName() string {
	return "projector_message"
}
