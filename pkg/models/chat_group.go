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

func (m ChatGroup) Get(field string) interface{} {
	switch field {
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "read_group_ids":
		return m.ReadGroupIDs
	case "weight":
		return m.Weight
	case "write_group_ids":
		return m.WriteGroupIDs
	}

	return nil
}
