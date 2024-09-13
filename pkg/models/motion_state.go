package models

type MotionState struct {
	AllowCreatePoll                  *bool
	AllowMotionForwarding            *bool
	AllowSubmitterEdit               *bool
	AllowSupport                     *bool
	CssClass                         string
	FirstStateOfWorkflowID           *int
	ID                               *int
	IsInternal                       *bool
	MeetingID                        int
	MergeAmendmentIntoFinal          *string
	MotionIDs                        []int
	MotionRecommendationIDs          []int
	Name                             string
	NextStateIDs                     []int
	PreviousStateIDs                 []int
	RecommendationLabel              *string
	Restrictions                     []string
	SetNumber                        *bool
	SetWorkflowTimestamp             *bool
	ShowRecommendationExtensionField *bool
	ShowStateExtensionField          *bool
	SubmitterWithdrawBackIDs         []int
	SubmitterWithdrawStateID         *int
	Weight                           int
	WorkflowID                       int
}

func (m MotionState) CollectionName() string {
	return "motion_state"
}
