package models

type AgendaItem struct {
	ChildIDs        []int
	Closed          *bool
	Comment         *string
	ContentObjectID string
	Duration        *int
	ID              *int
	IsHidden        *bool
	IsInternal      *bool
	ItemNumber      *string
	Level           *int
	MeetingID       int
	ModeratorNotes  *string
	ParentID        *int
	ProjectionIDs   []int
	TagIDs          []int
	Type            *string
	Weight          *int
}

func (m AgendaItem) CollectionName() string {
	return "agenda_item"
}
