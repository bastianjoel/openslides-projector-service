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

func (m Motion) Update(data map[string]string) error {
	if val, ok := data["additional_submitter"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdditionalSubmitter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_derived_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllDerivedMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_origin_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllOriginIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["amendment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AmendmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["amendment_paragraphs"]; ok {
		err := json.Unmarshal([]byte(val), &m.AmendmentParagraphs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["attachment_meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AttachmentMeetingMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["block_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.BlockID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["category_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CategoryID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["category_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.CategoryWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["change_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChangeRecommendationIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
		if err != nil {
			return err
		}
	}

	if val, ok := data["derived_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DerivedMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.EditorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarded"]; ok {
		err := json.Unmarshal([]byte(val), &m.Forwarded)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["identical_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.IDenticalMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["last_modified"]; ok {
		err := json.Unmarshal([]byte(val), &m.LastModified)
		if err != nil {
			return err
		}
	}

	if val, ok := data["lead_motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LeadMotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["modified_final_version"]; ok {
		err := json.Unmarshal([]byte(val), &m.ModifiedFinalVersion)
		if err != nil {
			return err
		}
	}

	if val, ok := data["number"]; ok {
		err := json.Unmarshal([]byte(val), &m.Number)
		if err != nil {
			return err
		}
	}

	if val, ok := data["number_value"]; ok {
		err := json.Unmarshal([]byte(val), &m.NumberValue)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["origin_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OriginID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["origin_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OriginMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["reason"]; ok {
		err := json.Unmarshal([]byte(val), &m.Reason)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_extension"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationExtension)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_extension_reference_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationExtensionReferenceIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["referenced_in_motion_recommendation_extension_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferencedInMotionRecommendationExtensionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["referenced_in_motion_state_extension_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferencedInMotionStateExtensionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sort_child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortChildIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sort_parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortParentID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sort_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["start_line_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.StartLineNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_extension"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateExtension)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_extension_reference_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateExtensionReferenceIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["statute_paragraph_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StatuteParagraphID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["supporter_meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SupporterMeetingUserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text_hash"]; ok {
		err := json.Unmarshal([]byte(val), &m.TextHash)
		if err != nil {
			return err
		}
	}

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
		if err != nil {
			return err
		}
	}

	if val, ok := data["workflow_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkflowTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
