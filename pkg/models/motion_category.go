package models

type MotionCategory struct {
	ParentID         *int
	SequentialNumber int
	ChildIDs         []int
	ID               *int
	Level            *int
	MeetingID        int
	MotionIDs        []int
	Name             string
	Prefix           *string
	Weight           *int
}

func (m MotionCategory) CollectionName() string {
	return "motion_category"
}
