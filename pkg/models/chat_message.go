package models

type ChatMessage struct {
	ChatGroupID   int
	Content       string
	Created       int
	ID            *int
	MeetingID     int
	MeetingUserID *int
}

func (m ChatMessage) CollectionName() string {
	return "chat_message"
}
