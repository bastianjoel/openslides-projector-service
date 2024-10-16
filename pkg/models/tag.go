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

func (m Tag) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "tagged_ids":
		return m.TaggedIDs
	}

	return nil
}
