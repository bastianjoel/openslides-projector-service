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

func (m MotionWorkflow) Get(field string) interface{} {
	switch field {
	case "default_amendment_workflow_meeting_id":
		return m.DefaultAmendmentWorkflowMeetingID
	case "default_statute_amendment_workflow_meeting_id":
		return m.DefaultStatuteAmendmentWorkflowMeetingID
	case "default_workflow_meeting_id":
		return m.DefaultWorkflowMeetingID
	case "first_state_id":
		return m.FirstStateID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "sequential_number":
		return m.SequentialNumber
	case "state_ids":
		return m.StateIDs
	}

	return nil
}
