package models

type PersonalNote struct {
	ContentObjectID *string `json:"content_object_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	MeetingUserID   int     `json:"meeting_user_id"`
	Note            *string `json:"note"`
	Star            *bool   `json:"star"`
}

func (m PersonalNote) CollectionName() string {
	return "personal_note"
}

func (m PersonalNote) Get(field string) interface{} {
	switch field {
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "note":
		return m.Note
	case "star":
		return m.Star
	}

	return nil
}
