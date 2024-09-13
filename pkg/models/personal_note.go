package models

type PersonalNote struct {
	ID              *int
	MeetingID       int
	MeetingUserID   int
	Note            *string
	Star            *bool
	ContentObjectID *string
}

func (m PersonalNote) CollectionName() string {
	return "personal_note"
}
