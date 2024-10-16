package models

type ChatMessage struct {
	ChatGroupID   int    `json:"chat_group_id"`
	Content       string `json:"content"`
	Created       int    `json:"created"`
	ID            int    `json:"id"`
	MeetingID     int    `json:"meeting_id"`
	MeetingUserID *int   `json:"meeting_user_id"`
}

func (m ChatMessage) CollectionName() string {
	return "chat_message"
}

func (m ChatMessage) Get(field string) interface{} {
	switch field {
	case "chat_group_id":
		return m.ChatGroupID
	case "content":
		return m.Content
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	}

	return nil
}
