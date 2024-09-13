package models

type MotionCategory struct {
	ChildIDs         []int
	ID               *int
	Level            *int
	MeetingID        int
	MotionIDs        []int
	Name             string
	ParentID         *int
	Prefix           *string
	SequentialNumber int
	Weight           *int
}

func (m MotionCategory) CollectionName() string {
	return "motion_category"
}
