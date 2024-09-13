package models

type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID        *int
	DefaultStatuteAmendmentWorkflowMeetingID *int
	DefaultWorkflowMeetingID                 *int
	FirstStateID                             int
	ID                                       *int
	MeetingID                                int
	Name                                     string
	SequentialNumber                         int
	StateIDs                                 []int
}

func (m MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}
