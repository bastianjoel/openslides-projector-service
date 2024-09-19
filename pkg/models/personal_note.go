package models

type PersonalNote struct {
	ContentObjectID *string `json:"content_object_id"`
	ID              *int    `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	MeetingUserID   int     `json:"meeting_user_id"`
	Note            *string `json:"note"`
	Star            *bool   `json:"star"`
}

func (m PersonalNote) CollectionName() string {
	return "personal_note"
}
