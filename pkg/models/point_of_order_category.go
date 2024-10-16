package models

type PointOfOrderCategory struct {
	ID         int    `json:"id"`
	MeetingID  int    `json:"meeting_id"`
	Rank       int    `json:"rank"`
	SpeakerIDs []int  `json:"speaker_ids"`
	Text       string `json:"text"`
}

func (m PointOfOrderCategory) CollectionName() string {
	return "point_of_order_category"
}

func (m PointOfOrderCategory) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "rank":
		return m.Rank
	case "speaker_ids":
		return m.SpeakerIDs
	case "text":
		return m.Text
	}

	return nil
}
