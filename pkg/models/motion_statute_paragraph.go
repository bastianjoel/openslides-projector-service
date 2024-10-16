package models

type MotionStatuteParagraph struct {
	ID               int     `json:"id"`
	MeetingID        int     `json:"meeting_id"`
	MotionIDs        []int   `json:"motion_ids"`
	SequentialNumber int     `json:"sequential_number"`
	Text             *string `json:"text"`
	Title            string  `json:"title"`
	Weight           *int    `json:"weight"`
}

func (m MotionStatuteParagraph) CollectionName() string {
	return "motion_statute_paragraph"
}

func (m MotionStatuteParagraph) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "text":
		return m.Text
	case "title":
		return m.Title
	case "weight":
		return m.Weight
	}

	return nil
}
