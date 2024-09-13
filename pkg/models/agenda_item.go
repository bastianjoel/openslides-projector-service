package models

type AgendaItem struct {
	Type            *string
	MeetingID       int
	ModeratorNotes  *string
	TagIDs          []int
	ContentObjectID string
	Duration        *int
	IsHidden        *bool
	IsInternal      *bool
	ItemNumber      *string
	ChildIDs        []int
	Closed          *bool
	Comment         *string
	ProjectionIDs   []int
	Weight          *int
	Level           *int
	ParentID        *int
	ID              *int
}

func (m AgendaItem) CollectionName() string {
	return "agenda_item"
}
