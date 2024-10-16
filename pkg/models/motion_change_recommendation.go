package models

type MotionChangeRecommendation struct {
	CreationTime     *int    `json:"creation_time"`
	ID               int     `json:"id"`
	Internal         *bool   `json:"internal"`
	LineFrom         *int    `json:"line_from"`
	LineTo           *int    `json:"line_to"`
	MeetingID        int     `json:"meeting_id"`
	MotionID         int     `json:"motion_id"`
	OtherDescription *string `json:"other_description"`
	Rejected         *bool   `json:"rejected"`
	Text             *string `json:"text"`
	Type             *string `json:"type"`
}

func (m MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}

func (m MotionChangeRecommendation) Get(field string) interface{} {
	switch field {
	case "creation_time":
		return m.CreationTime
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "line_from":
		return m.LineFrom
	case "line_to":
		return m.LineTo
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "other_description":
		return m.OtherDescription
	case "rejected":
		return m.Rejected
	case "text":
		return m.Text
	case "type":
		return m.Type
	}

	return nil
}
