package models

type Tag struct {
	ID        int      `json:"id"`
	MeetingID int      `json:"meeting_id"`
	Name      string   `json:"name"`
	TaggedIDs []string `json:"tagged_ids"`
}

func (m Tag) CollectionName() string {
	return "tag"
}
