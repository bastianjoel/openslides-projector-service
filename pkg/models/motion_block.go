package models

type MotionBlock struct {
	ID               *int
	Internal         *bool
	ListOfSpeakersID int
	ProjectionIDs    []int
	Title            string
	AgendaItemID     *int
	MeetingID        int
	MotionIDs        []int
	SequentialNumber int
}

func (m MotionBlock) CollectionName() string {
	return "motion_block"
}
