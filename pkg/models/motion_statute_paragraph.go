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
