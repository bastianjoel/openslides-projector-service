package models

import "encoding/json"

type Motion struct {
	StateExtension                               *string
	OriginID                                     *int
	SortChildIDs                                 []int
	StateID                                      int
	WorkingGroupSpeakerIDs                       []int
	Created                                      *int
	Reason                                       *string
	RecommendationExtensionReferenceIDs          []string
	NumberValue                                  *int
	AmendmentParagraphs                          json.RawMessage
	BlockID                                      *int
	ReferencedInMotionRecommendationExtensionIDs []int
	MeetingID                                    int
	RecommendationID                             *int
	StatuteParagraphID                           *int
	TextHash                                     *string
	AllOriginIDs                                 []int
	DerivedMotionIDs                             []int
	LeadMotionID                                 *int
	PersonalNoteIDs                              []int
	RecommendationExtension                      *string
	EditorIDs                                    []int
	Forwarded                                    *int
	LastModified                                 *int
	Text                                         *string
	AmendmentIDs                                 []int
	CategoryID                                   *int
	CategoryWeight                               *int
	ID                                           *int
	SequentialNumber                             int
	StateExtensionReferenceIDs                   []string
	ChangeRecommendationIDs                      []int
	CommentIDs                                   []int
	OriginMeetingID                              *int
	SortWeight                                   *int
	ModifiedFinalVersion                         *string
	PollIDs                                      []int
	SortParentID                                 *int
	SupporterMeetingUserIDs                      []int
	TagIDs                                       []int
	WorkflowTimestamp                            *int
	AgendaItemID                                 *int
	AttachmentMeetingMediafileIDs                []int
	ListOfSpeakersID                             int
	SubmitterIDs                                 []int
	AllDerivedMotionIDs                          []int
	IDenticalMotionIDs                           []int
	Number                                       *string
	StartLineNumber                              *int
	ReferencedInMotionStateExtensionIDs          []int
	AdditionalSubmitter                          *string
	OptionIDs                                    []int
	ProjectionIDs                                []int
	Title                                        string
}

func (m Motion) CollectionName() string {
	return "motion"
}
