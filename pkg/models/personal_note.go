package models

type PersonalNote struct {
	ContentObjectID *string
	ID              *int
	MeetingID       int
	MeetingUserID   int
	Note            *string
	Star            *bool
}

func (m PersonalNote) CollectionName() string {
	return "personal_note"
}
