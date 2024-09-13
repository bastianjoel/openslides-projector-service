package models

type MotionStatuteParagraph struct {
	ID               *int
	MeetingID        int
	MotionIDs        []int
	SequentialNumber int
	Text             *string
	Title            string
	Weight           *int
}

func (m MotionStatuteParagraph) CollectionName() string {
	return "motion_statute_paragraph"
}
