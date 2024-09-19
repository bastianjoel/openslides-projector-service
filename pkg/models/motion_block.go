package models

type MotionBlock struct {
	AgendaItemID     *int   `json:"agenda_item_id"`
	ID               *int   `json:"id"`
	Internal         *bool  `json:"internal"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	MeetingID        int    `json:"meeting_id"`
	MotionIDs        []int  `json:"motion_ids"`
	ProjectionIDs    []int  `json:"projection_ids"`
	SequentialNumber int    `json:"sequential_number"`
	Title            string `json:"title"`
}

func (m MotionBlock) CollectionName() string {
	return "motion_block"
}
