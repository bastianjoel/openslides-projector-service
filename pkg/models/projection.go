package models

import "encoding/json"

type Projection struct {
	ID                 *int
	Stable             *bool
	Weight             *int
	ContentObjectID    string
	CurrentProjectorID *int
	HistoryProjectorID *int
	MeetingID          int
	Options            json.RawMessage
	PreviewProjectorID *int
	Type               *string
	Content            json.RawMessage
}

func (m Projection) CollectionName() string {
	return "projection"
}
