package models

type MotionWorkflow struct {
	DefaultWorkflowMeetingID                 *int
	FirstStateID                             int
	ID                                       *int
	Name                                     string
	StateIDs                                 []int
	DefaultAmendmentWorkflowMeetingID        *int
	DefaultStatuteAmendmentWorkflowMeetingID *int
	MeetingID                                int
	SequentialNumber                         int
}

func (m MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}
