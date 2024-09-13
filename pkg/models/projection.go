package models

import "encoding/json"

type Projection struct {
	Content            json.RawMessage
	ContentObjectID    string
	CurrentProjectorID *int
	HistoryProjectorID *int
	ID                 *int
	MeetingID          int
	Options            json.RawMessage
	PreviewProjectorID *int
	Stable             *bool
	Type               *string
	Weight             *int
}

func (m Projection) CollectionName() string {
	return "projection"
}
