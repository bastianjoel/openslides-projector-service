package models

type MotionChangeRecommendation struct {
	CreationTime     *int    `json:"creation_time"`
	ID               *int    `json:"id"`
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
