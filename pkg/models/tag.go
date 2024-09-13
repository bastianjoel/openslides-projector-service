package models

type Tag struct {
	ID        *int
	MeetingID int
	Name      string
	TaggedIDs []string
}

func (m Tag) CollectionName() string {
	return "tag"
}
