package models

type PointOfOrderCategory struct {
	ID         *int   `json:"id"`
	MeetingID  int    `json:"meeting_id"`
	Rank       int    `json:"rank"`
	SpeakerIDs []int  `json:"speaker_ids"`
	Text       string `json:"text"`
}

func (m PointOfOrderCategory) CollectionName() string {
	return "point_of_order_category"
}
