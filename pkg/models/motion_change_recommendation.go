package models

type MotionChangeRecommendation struct {
	ID               *int
	LineFrom         *int
	LineTo           *int
	Text             *string
	Type             *string
	CreationTime     *int
	Internal         *bool
	MeetingID        int
	MotionID         int
	OtherDescription *string
	Rejected         *bool
}

func (m MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}
