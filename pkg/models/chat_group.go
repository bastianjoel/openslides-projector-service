package models

type ChatGroup struct {
	ChatMessageIDs []int
	ID             *int
	MeetingID      int
	Name           string
	ReadGroupIDs   []int
	Weight         *int
	WriteGroupIDs  []int
}

func (m ChatGroup) CollectionName() string {
	return "chat_group"
}
