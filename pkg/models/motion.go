package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

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
	loadedRelations                              map[string]struct{}
	allDerivedMotions                            *Motion
	personalNotes                                *PersonalNote
	workingGroupSpeakers                         *MotionWorkingGroupSpeaker
	recommendation                               *MotionState
	submitters                                   *MotionSubmitter
	listOfSpeakers                               *ListOfSpeakers
	originMeeting                                *Meeting
	attachmentMeetingMediafiles                  *MeetingMediafile
	referencedInMotionStateExtensions            *Motion
	sortChilds                                   *Motion
	meeting                                      *Meeting
	editors                                      *MotionEditor
	polls                                        *Poll
	block                                        *MotionBlock
	category                                     *MotionCategory
	supporterMeetingUsers                        *MeetingUser
	agendaItem                                   *AgendaItem
	sortParent                                   *Motion
	statuteParagraph                             *MotionStatuteParagraph
	amendments                                   *Motion
	origin                                       *Motion
	referencedInMotionRecommendationExtensions   *Motion
	allOrigins                                   *Motion
	options                                      *Option
	projections                                  *Projection
	leadMotion                                   *Motion
	comments                                     *MotionComment
	tags                                         *Tag
	changeRecommendations                        *MotionChangeRecommendation
	derivedMotions                               *Motion
	iDenticalMotions                             *Motion
	state                                        *MotionState
}

func (m *Motion) CollectionName() string {
	return "motion"
}

func (m *Motion) AllDerivedMotions() *Motion {
	if _, ok := m.loadedRelations["all_derived_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access AllDerivedMotions relation of Motion which was not loaded.")
	}

	return m.allDerivedMotions
}

func (m *Motion) PersonalNotes() *PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of Motion which was not loaded.")
	}

	return m.personalNotes
}

func (m *Motion) WorkingGroupSpeakers() *MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access WorkingGroupSpeakers relation of Motion which was not loaded.")
	}

	return m.workingGroupSpeakers
}

func (m *Motion) Recommendation() *MotionState {
	if _, ok := m.loadedRelations["recommendation_id"]; !ok {
		log.Panic().Msg("Tried to access Recommendation relation of Motion which was not loaded.")
	}

	return m.recommendation
}

func (m *Motion) Submitters() *MotionSubmitter {
	if _, ok := m.loadedRelations["submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access Submitters relation of Motion which was not loaded.")
	}

	return m.submitters
}

func (m *Motion) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Motion which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Motion) OriginMeeting() *Meeting {
	if _, ok := m.loadedRelations["origin_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access OriginMeeting relation of Motion which was not loaded.")
	}

	return m.originMeeting
}

func (m *Motion) AttachmentMeetingMediafiles() *MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Motion which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Motion) ReferencedInMotionStateExtensions() *Motion {
	if _, ok := m.loadedRelations["referenced_in_motion_state_extension_ids"]; !ok {
		log.Panic().Msg("Tried to access ReferencedInMotionStateExtensions relation of Motion which was not loaded.")
	}

	return m.referencedInMotionStateExtensions
}

func (m *Motion) SortChilds() *Motion {
	if _, ok := m.loadedRelations["sort_child_ids"]; !ok {
		log.Panic().Msg("Tried to access SortChilds relation of Motion which was not loaded.")
	}

	return m.sortChilds
}

func (m *Motion) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Motion which was not loaded.")
	}

	return *m.meeting
}

func (m *Motion) Editors() *MotionEditor {
	if _, ok := m.loadedRelations["editor_ids"]; !ok {
		log.Panic().Msg("Tried to access Editors relation of Motion which was not loaded.")
	}

	return m.editors
}

func (m *Motion) Polls() *Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Motion which was not loaded.")
	}

	return m.polls
}

func (m *Motion) Block() *MotionBlock {
	if _, ok := m.loadedRelations["block_id"]; !ok {
		log.Panic().Msg("Tried to access Block relation of Motion which was not loaded.")
	}

	return m.block
}

func (m *Motion) Category() *MotionCategory {
	if _, ok := m.loadedRelations["category_id"]; !ok {
		log.Panic().Msg("Tried to access Category relation of Motion which was not loaded.")
	}

	return m.category
}

func (m *Motion) SupporterMeetingUsers() *MeetingUser {
	if _, ok := m.loadedRelations["supporter_meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access SupporterMeetingUsers relation of Motion which was not loaded.")
	}

	return m.supporterMeetingUsers
}

func (m *Motion) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Motion which was not loaded.")
	}

	return m.agendaItem
}

func (m *Motion) SortParent() *Motion {
	if _, ok := m.loadedRelations["sort_parent_id"]; !ok {
		log.Panic().Msg("Tried to access SortParent relation of Motion which was not loaded.")
	}

	return m.sortParent
}

func (m *Motion) StatuteParagraph() *MotionStatuteParagraph {
	if _, ok := m.loadedRelations["statute_paragraph_id"]; !ok {
		log.Panic().Msg("Tried to access StatuteParagraph relation of Motion which was not loaded.")
	}

	return m.statuteParagraph
}

func (m *Motion) Amendments() *Motion {
	if _, ok := m.loadedRelations["amendment_ids"]; !ok {
		log.Panic().Msg("Tried to access Amendments relation of Motion which was not loaded.")
	}

	return m.amendments
}

func (m *Motion) Origin() *Motion {
	if _, ok := m.loadedRelations["origin_id"]; !ok {
		log.Panic().Msg("Tried to access Origin relation of Motion which was not loaded.")
	}

	return m.origin
}

func (m *Motion) ReferencedInMotionRecommendationExtensions() *Motion {
	if _, ok := m.loadedRelations["referenced_in_motion_recommendation_extension_ids"]; !ok {
		log.Panic().Msg("Tried to access ReferencedInMotionRecommendationExtensions relation of Motion which was not loaded.")
	}

	return m.referencedInMotionRecommendationExtensions
}

func (m *Motion) AllOrigins() *Motion {
	if _, ok := m.loadedRelations["all_origin_ids"]; !ok {
		log.Panic().Msg("Tried to access AllOrigins relation of Motion which was not loaded.")
	}

	return m.allOrigins
}

func (m *Motion) Options() *Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of Motion which was not loaded.")
	}

	return m.options
}

func (m *Motion) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Motion which was not loaded.")
	}

	return m.projections
}

func (m *Motion) LeadMotion() *Motion {
	if _, ok := m.loadedRelations["lead_motion_id"]; !ok {
		log.Panic().Msg("Tried to access LeadMotion relation of Motion which was not loaded.")
	}

	return m.leadMotion
}

func (m *Motion) Comments() *MotionComment {
	if _, ok := m.loadedRelations["comment_ids"]; !ok {
		log.Panic().Msg("Tried to access Comments relation of Motion which was not loaded.")
	}

	return m.comments
}

func (m *Motion) Tags() *Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Motion which was not loaded.")
	}

	return m.tags
}

func (m *Motion) ChangeRecommendations() *MotionChangeRecommendation {
	if _, ok := m.loadedRelations["change_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access ChangeRecommendations relation of Motion which was not loaded.")
	}

	return m.changeRecommendations
}

func (m *Motion) DerivedMotions() *Motion {
	if _, ok := m.loadedRelations["derived_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access DerivedMotions relation of Motion which was not loaded.")
	}

	return m.derivedMotions
}

func (m *Motion) IDenticalMotions() *Motion {
	if _, ok := m.loadedRelations["identical_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access IDenticalMotions relation of Motion which was not loaded.")
	}

	return m.iDenticalMotions
}

func (m *Motion) State() MotionState {
	if _, ok := m.loadedRelations["state_id"]; !ok {
		log.Panic().Msg("Tried to access State relation of Motion which was not loaded.")
	}

	return *m.state
}

func (m *Motion) Get(field string) interface{} {
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

func (m *Motion) Update(data map[string]string) error {
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
