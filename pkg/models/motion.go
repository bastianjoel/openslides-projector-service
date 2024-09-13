package models

import "encoding/json"

type Motion struct {
	AdditionalSubmitter                          *string
	AgendaItemID                                 *int
	AllDerivedMotionIDs                          []int
	AllOriginIDs                                 []int
	AmendmentIDs                                 []int
	AmendmentParagraphs                          json.RawMessage
	AttachmentMeetingMediafileIDs                []int
	BlockID                                      *int
	CategoryID                                   *int
	CategoryWeight                               *int
	ChangeRecommendationIDs                      []int
	CommentIDs                                   []int
	Created                                      *int
	DerivedMotionIDs                             []int
	EditorIDs                                    []int
	Forwarded                                    *int
	ID                                           *int
	IDenticalMotionIDs                           []int
	LastModified                                 *int
	LeadMotionID                                 *int
	ListOfSpeakersID                             int
	MeetingID                                    int
	ModifiedFinalVersion                         *string
	Number                                       *string
	NumberValue                                  *int
	OptionIDs                                    []int
	OriginID                                     *int
	OriginMeetingID                              *int
	PersonalNoteIDs                              []int
	PollIDs                                      []int
	ProjectionIDs                                []int
	Reason                                       *string
	RecommendationExtension                      *string
	RecommendationExtensionReferenceIDs          []string
	RecommendationID                             *int
	ReferencedInMotionRecommendationExtensionIDs []int
	ReferencedInMotionStateExtensionIDs          []int
	SequentialNumber                             int
	SortChildIDs                                 []int
	SortParentID                                 *int
	SortWeight                                   *int
	StartLineNumber                              *int
	StateExtension                               *string
	StateExtensionReferenceIDs                   []string
	StateID                                      int
	StatuteParagraphID                           *int
	SubmitterIDs                                 []int
	SupporterMeetingUserIDs                      []int
	TagIDs                                       []int
	Text                                         *string
	TextHash                                     *string
	Title                                        string
	WorkflowTimestamp                            *int
	WorkingGroupSpeakerIDs                       []int
}

func (m Motion) CollectionName() string {
	return "motion"
}
