package models

type MotionChangeRecommendation struct {
	CreationTime     *int
	ID               *int
	Internal         *bool
	LineFrom         *int
	LineTo           *int
	MeetingID        int
	MotionID         int
	OtherDescription *string
	Rejected         *bool
	Text             *string
	Type             *string
}

func (m MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}
