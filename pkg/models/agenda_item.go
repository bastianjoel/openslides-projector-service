package models

type AgendaItem struct {
	ChildIDs        []int   `json:"child_ids"`
	Closed          *bool   `json:"closed"`
	Comment         *string `json:"comment"`
	ContentObjectID string  `json:"content_object_id"`
	Duration        *int    `json:"duration"`
	ID              int     `json:"id"`
	IsHidden        *bool   `json:"is_hidden"`
	IsInternal      *bool   `json:"is_internal"`
	ItemNumber      *string `json:"item_number"`
	Level           *int    `json:"level"`
	MeetingID       int     `json:"meeting_id"`
	ModeratorNotes  *string `json:"moderator_notes"`
	ParentID        *int    `json:"parent_id"`
	ProjectionIDs   []int   `json:"projection_ids"`
	TagIDs          []int   `json:"tag_ids"`
	Type            *string `json:"type"`
	Weight          *int    `json:"weight"`
}

func (m AgendaItem) CollectionName() string {
	return "agenda_item"
}
