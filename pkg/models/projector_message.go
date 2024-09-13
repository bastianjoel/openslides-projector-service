package models

type ProjectorMessage struct {
	ID            *int
	MeetingID     int
	Message       *string
	ProjectionIDs []int
}

func (m ProjectorMessage) CollectionName() string {
	return "projector_message"
}
