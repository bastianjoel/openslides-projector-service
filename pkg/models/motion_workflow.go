package models

type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID        *int   `json:"default_amendment_workflow_meeting_id"`
	DefaultStatuteAmendmentWorkflowMeetingID *int   `json:"default_statute_amendment_workflow_meeting_id"`
	DefaultWorkflowMeetingID                 *int   `json:"default_workflow_meeting_id"`
	FirstStateID                             int    `json:"first_state_id"`
	ID                                       int    `json:"id"`
	MeetingID                                int    `json:"meeting_id"`
	Name                                     string `json:"name"`
	SequentialNumber                         int    `json:"sequential_number"`
	StateIDs                                 []int  `json:"state_ids"`
}

func (m MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}
