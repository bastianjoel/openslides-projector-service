package models

type ChatGroup struct {
	ChatMessageIDs []int  `json:"chat_message_ids"`
	ID             int    `json:"id"`
	MeetingID      int    `json:"meeting_id"`
	Name           string `json:"name"`
	ReadGroupIDs   []int  `json:"read_group_ids"`
	Weight         *int   `json:"weight"`
	WriteGroupIDs  []int  `json:"write_group_ids"`
}

func (m ChatGroup) CollectionName() string {
	return "chat_group"
}
