package models

type ChatMessage struct {
	ChatGroupID   int    `json:"chat_group_id"`
	Content       string `json:"content"`
	Created       int    `json:"created"`
	ID            *int   `json:"id"`
	MeetingID     int    `json:"meeting_id"`
	MeetingUserID *int   `json:"meeting_user_id"`
}

func (m ChatMessage) CollectionName() string {
	return "chat_message"
}
