package models

type MotionState struct {
	AllowCreatePoll                  *bool    `json:"allow_create_poll"`
	AllowMotionForwarding            *bool    `json:"allow_motion_forwarding"`
	AllowSubmitterEdit               *bool    `json:"allow_submitter_edit"`
	AllowSupport                     *bool    `json:"allow_support"`
	CssClass                         string   `json:"css_class"`
	FirstStateOfWorkflowID           *int     `json:"first_state_of_workflow_id"`
	ID                               int      `json:"id"`
	IsInternal                       *bool    `json:"is_internal"`
	MeetingID                        int      `json:"meeting_id"`
	MergeAmendmentIntoFinal          *string  `json:"merge_amendment_into_final"`
	MotionIDs                        []int    `json:"motion_ids"`
	MotionRecommendationIDs          []int    `json:"motion_recommendation_ids"`
	Name                             string   `json:"name"`
	NextStateIDs                     []int    `json:"next_state_ids"`
	PreviousStateIDs                 []int    `json:"previous_state_ids"`
	RecommendationLabel              *string  `json:"recommendation_label"`
	Restrictions                     []string `json:"restrictions"`
	SetNumber                        *bool    `json:"set_number"`
	SetWorkflowTimestamp             *bool    `json:"set_workflow_timestamp"`
	ShowRecommendationExtensionField *bool    `json:"show_recommendation_extension_field"`
	ShowStateExtensionField          *bool    `json:"show_state_extension_field"`
	SubmitterWithdrawBackIDs         []int    `json:"submitter_withdraw_back_ids"`
	SubmitterWithdrawStateID         *int     `json:"submitter_withdraw_state_id"`
	Weight                           int      `json:"weight"`
	WorkflowID                       int      `json:"workflow_id"`
}

func (m MotionState) CollectionName() string {
	return "motion_state"
}

func (m MotionState) Get(field string) interface{} {
	switch field {
	case "allow_create_poll":
		return m.AllowCreatePoll
	case "allow_motion_forwarding":
		return m.AllowMotionForwarding
	case "allow_submitter_edit":
		return m.AllowSubmitterEdit
	case "allow_support":
		return m.AllowSupport
	case "css_class":
		return m.CssClass
	case "first_state_of_workflow_id":
		return m.FirstStateOfWorkflowID
	case "id":
		return m.ID
	case "is_internal":
		return m.IsInternal
	case "meeting_id":
		return m.MeetingID
	case "merge_amendment_into_final":
		return m.MergeAmendmentIntoFinal
	case "motion_ids":
		return m.MotionIDs
	case "motion_recommendation_ids":
		return m.MotionRecommendationIDs
	case "name":
		return m.Name
	case "next_state_ids":
		return m.NextStateIDs
	case "previous_state_ids":
		return m.PreviousStateIDs
	case "recommendation_label":
		return m.RecommendationLabel
	case "restrictions":
		return m.Restrictions
	case "set_number":
		return m.SetNumber
	case "set_workflow_timestamp":
		return m.SetWorkflowTimestamp
	case "show_recommendation_extension_field":
		return m.ShowRecommendationExtensionField
	case "show_state_extension_field":
		return m.ShowStateExtensionField
	case "submitter_withdraw_back_ids":
		return m.SubmitterWithdrawBackIDs
	case "submitter_withdraw_state_id":
		return m.SubmitterWithdrawStateID
	case "weight":
		return m.Weight
	case "workflow_id":
		return m.WorkflowID
	}

	return nil
}
