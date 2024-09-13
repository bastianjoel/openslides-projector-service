package models

type MotionStatuteParagraph struct {
	Weight           *int
	ID               *int
	MeetingID        int
	MotionIDs        []int
	SequentialNumber int
	Text             *string
	Title            string
}

func (m MotionStatuteParagraph) CollectionName() string {
	return "motion_statute_paragraph"
}
