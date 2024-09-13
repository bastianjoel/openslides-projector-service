package models

type ChatGroup struct {
	ID             *int
	MeetingID      int
	Name           string
	ReadGroupIDs   []int
	Weight         *int
	WriteGroupIDs  []int
	ChatMessageIDs []int
}

func (m ChatGroup) CollectionName() string {
	return "chat_group"
}
