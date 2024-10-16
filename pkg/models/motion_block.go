package models

type MotionBlock struct {
	AgendaItemID     *int   `json:"agenda_item_id"`
	ID               int    `json:"id"`
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

func (m MotionBlock) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "title":
		return m.Title
	}

	return nil
}
