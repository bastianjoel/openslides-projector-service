package models

type MotionBlock struct {
	AgendaItemID     *int
	ID               *int
	Internal         *bool
	ListOfSpeakersID int
	MeetingID        int
	MotionIDs        []int
	ProjectionIDs    []int
	SequentialNumber int
	Title            string
}

func (m MotionBlock) CollectionName() string {
	return "motion_block"
}
