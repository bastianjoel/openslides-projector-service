package models

type PointOfOrderCategory struct {
	ID         *int
	MeetingID  int
	Rank       int
	SpeakerIDs []int
	Text       string
}

func (m PointOfOrderCategory) CollectionName() string {
	return "point_of_order_category"
}
