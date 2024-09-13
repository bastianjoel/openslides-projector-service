package models

type MotionState struct {
	ShowStateExtensionField          *bool
	AllowMotionForwarding            *bool
	AllowSubmitterEdit               *bool
	CssClass                         string
	ID                               *int
	PreviousStateIDs                 []int
	RecommendationLabel              *string
	ShowRecommendationExtensionField *bool
	SubmitterWithdrawStateID         *int
	WorkflowID                       int
	AllowSupport                     *bool
	MotionRecommendationIDs          []int
	Name                             string
	NextStateIDs                     []int
	SetNumber                        *bool
	AllowCreatePoll                  *bool
	FirstStateOfWorkflowID           *int
	MotionIDs                        []int
	Restrictions                     []string
	SubmitterWithdrawBackIDs         []int
	IsInternal                       *bool
	MeetingID                        int
	MergeAmendmentIntoFinal          *string
	SetWorkflowTimestamp             *bool
	Weight                           int
}

func (m MotionState) CollectionName() string {
	return "motion_state"
}
