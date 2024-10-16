package models

import "encoding/json"

type Motion struct {
	AdditionalSubmitter                          *string         `json:"additional_submitter"`
	AgendaItemID                                 *int            `json:"agenda_item_id"`
	AllDerivedMotionIDs                          []int           `json:"all_derived_motion_ids"`
	AllOriginIDs                                 []int           `json:"all_origin_ids"`
	AmendmentIDs                                 []int           `json:"amendment_ids"`
	AmendmentParagraphs                          json.RawMessage `json:"amendment_paragraphs"`
	AttachmentMeetingMediafileIDs                []int           `json:"attachment_meeting_mediafile_ids"`
	BlockID                                      *int            `json:"block_id"`
	CategoryID                                   *int            `json:"category_id"`
	CategoryWeight                               *int            `json:"category_weight"`
	ChangeRecommendationIDs                      []int           `json:"change_recommendation_ids"`
	CommentIDs                                   []int           `json:"comment_ids"`
	Created                                      *int            `json:"created"`
	DerivedMotionIDs                             []int           `json:"derived_motion_ids"`
	EditorIDs                                    []int           `json:"editor_ids"`
	Forwarded                                    *int            `json:"forwarded"`
	ID                                           int             `json:"id"`
	IDenticalMotionIDs                           []int           `json:"identical_motion_ids"`
	LastModified                                 *int            `json:"last_modified"`
	LeadMotionID                                 *int            `json:"lead_motion_id"`
	ListOfSpeakersID                             int             `json:"list_of_speakers_id"`
	MeetingID                                    int             `json:"meeting_id"`
	ModifiedFinalVersion                         *string         `json:"modified_final_version"`
	Number                                       *string         `json:"number"`
	NumberValue                                  *int            `json:"number_value"`
	OptionIDs                                    []int           `json:"option_ids"`
	OriginID                                     *int            `json:"origin_id"`
	OriginMeetingID                              *int            `json:"origin_meeting_id"`
	PersonalNoteIDs                              []int           `json:"personal_note_ids"`
	PollIDs                                      []int           `json:"poll_ids"`
	ProjectionIDs                                []int           `json:"projection_ids"`
	Reason                                       *string         `json:"reason"`
	RecommendationExtension                      *string         `json:"recommendation_extension"`
	RecommendationExtensionReferenceIDs          []string        `json:"recommendation_extension_reference_ids"`
	RecommendationID                             *int            `json:"recommendation_id"`
	ReferencedInMotionRecommendationExtensionIDs []int           `json:"referenced_in_motion_recommendation_extension_ids"`
	ReferencedInMotionStateExtensionIDs          []int           `json:"referenced_in_motion_state_extension_ids"`
	SequentialNumber                             int             `json:"sequential_number"`
	SortChildIDs                                 []int           `json:"sort_child_ids"`
	SortParentID                                 *int            `json:"sort_parent_id"`
	SortWeight                                   *int            `json:"sort_weight"`
	StartLineNumber                              *int            `json:"start_line_number"`
	StateExtension                               *string         `json:"state_extension"`
	StateExtensionReferenceIDs                   []string        `json:"state_extension_reference_ids"`
	StateID                                      int             `json:"state_id"`
	StatuteParagraphID                           *int            `json:"statute_paragraph_id"`
	SubmitterIDs                                 []int           `json:"submitter_ids"`
	SupporterMeetingUserIDs                      []int           `json:"supporter_meeting_user_ids"`
	TagIDs                                       []int           `json:"tag_ids"`
	Text                                         *string         `json:"text"`
	TextHash                                     *string         `json:"text_hash"`
	Title                                        string          `json:"title"`
	WorkflowTimestamp                            *int            `json:"workflow_timestamp"`
	WorkingGroupSpeakerIDs                       []int           `json:"working_group_speaker_ids"`
}

func (m Motion) CollectionName() string {
	return "motion"
}

func (m Motion) Get(field string) interface{} {
	switch field {
	case "additional_submitter":
		return m.AdditionalSubmitter
	case "agenda_item_id":
		return m.AgendaItemID
	case "all_derived_motion_ids":
		return m.AllDerivedMotionIDs
	case "all_origin_ids":
		return m.AllOriginIDs
	case "amendment_ids":
		return m.AmendmentIDs
	case "amendment_paragraphs":
		return m.AmendmentParagraphs
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "block_id":
		return m.BlockID
	case "category_id":
		return m.CategoryID
	case "category_weight":
		return m.CategoryWeight
	case "change_recommendation_ids":
		return m.ChangeRecommendationIDs
	case "comment_ids":
		return m.CommentIDs
	case "created":
		return m.Created
	case "derived_motion_ids":
		return m.DerivedMotionIDs
	case "editor_ids":
		return m.EditorIDs
	case "forwarded":
		return m.Forwarded
	case "id":
		return m.ID
	case "identical_motion_ids":
		return m.IDenticalMotionIDs
	case "last_modified":
		return m.LastModified
	case "lead_motion_id":
		return m.LeadMotionID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "modified_final_version":
		return m.ModifiedFinalVersion
	case "number":
		return m.Number
	case "number_value":
		return m.NumberValue
	case "option_ids":
		return m.OptionIDs
	case "origin_id":
		return m.OriginID
	case "origin_meeting_id":
		return m.OriginMeetingID
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "reason":
		return m.Reason
	case "recommendation_extension":
		return m.RecommendationExtension
	case "recommendation_extension_reference_ids":
		return m.RecommendationExtensionReferenceIDs
	case "recommendation_id":
		return m.RecommendationID
	case "referenced_in_motion_recommendation_extension_ids":
		return m.ReferencedInMotionRecommendationExtensionIDs
	case "referenced_in_motion_state_extension_ids":
		return m.ReferencedInMotionStateExtensionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "sort_child_ids":
		return m.SortChildIDs
	case "sort_parent_id":
		return m.SortParentID
	case "sort_weight":
		return m.SortWeight
	case "start_line_number":
		return m.StartLineNumber
	case "state_extension":
		return m.StateExtension
	case "state_extension_reference_ids":
		return m.StateExtensionReferenceIDs
	case "state_id":
		return m.StateID
	case "statute_paragraph_id":
		return m.StatuteParagraphID
	case "submitter_ids":
		return m.SubmitterIDs
	case "supporter_meeting_user_ids":
		return m.SupporterMeetingUserIDs
	case "tag_ids":
		return m.TagIDs
	case "text":
		return m.Text
	case "text_hash":
		return m.TextHash
	case "title":
		return m.Title
	case "workflow_timestamp":
		return m.WorkflowTimestamp
	case "working_group_speaker_ids":
		return m.WorkingGroupSpeakerIDs
	}

	return nil
}
