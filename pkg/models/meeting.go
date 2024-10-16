package models

import "encoding/json"

type Meeting struct {
	AdminGroupID                                 *int            `json:"admin_group_id"`
	AgendaEnableNumbering                        *bool           `json:"agenda_enable_numbering"`
	AgendaItemCreation                           *string         `json:"agenda_item_creation"`
	AgendaItemIDs                                []int           `json:"agenda_item_ids"`
	AgendaNewItemsDefaultVisibility              *string         `json:"agenda_new_items_default_visibility"`
	AgendaNumberPrefix                           *string         `json:"agenda_number_prefix"`
	AgendaNumeralSystem                          *string         `json:"agenda_numeral_system"`
	AgendaShowInternalItemsOnProjector           *bool           `json:"agenda_show_internal_items_on_projector"`
	AgendaShowSubtitles                          *bool           `json:"agenda_show_subtitles"`
	AgendaShowTopicNavigationOnDetailView        *bool           `json:"agenda_show_topic_navigation_on_detail_view"`
	AllProjectionIDs                             []int           `json:"all_projection_ids"`
	AnonymousGroupID                             *int            `json:"anonymous_group_id"`
	ApplauseEnable                               *bool           `json:"applause_enable"`
	ApplauseMaxAmount                            *int            `json:"applause_max_amount"`
	ApplauseMinAmount                            *int            `json:"applause_min_amount"`
	ApplauseParticleImageUrl                     *string         `json:"applause_particle_image_url"`
	ApplauseShowLevel                            *bool           `json:"applause_show_level"`
	ApplauseTimeout                              *int            `json:"applause_timeout"`
	ApplauseType                                 *string         `json:"applause_type"`
	AssignmentCandidateIDs                       []int           `json:"assignment_candidate_ids"`
	AssignmentIDs                                []int           `json:"assignment_ids"`
	AssignmentPollAddCandidatesToListOfSpeakers  *bool           `json:"assignment_poll_add_candidates_to_list_of_speakers"`
	AssignmentPollBallotPaperNumber              *int            `json:"assignment_poll_ballot_paper_number"`
	AssignmentPollBallotPaperSelection           *string         `json:"assignment_poll_ballot_paper_selection"`
	AssignmentPollDefaultBackend                 *string         `json:"assignment_poll_default_backend"`
	AssignmentPollDefaultGroupIDs                []int           `json:"assignment_poll_default_group_ids"`
	AssignmentPollDefaultMethod                  *string         `json:"assignment_poll_default_method"`
	AssignmentPollDefaultOnehundredPercentBase   *string         `json:"assignment_poll_default_onehundred_percent_base"`
	AssignmentPollDefaultType                    *string         `json:"assignment_poll_default_type"`
	AssignmentPollEnableMaxVotesPerOption        *bool           `json:"assignment_poll_enable_max_votes_per_option"`
	AssignmentPollSortPollResultByVotes          *bool           `json:"assignment_poll_sort_poll_result_by_votes"`
	AssignmentsExportPreamble                    *string         `json:"assignments_export_preamble"`
	AssignmentsExportTitle                       *string         `json:"assignments_export_title"`
	ChatGroupIDs                                 []int           `json:"chat_group_ids"`
	ChatMessageIDs                               []int           `json:"chat_message_ids"`
	CommitteeID                                  int             `json:"committee_id"`
	ConferenceAutoConnect                        *bool           `json:"conference_auto_connect"`
	ConferenceAutoConnectNextSpeakers            *int            `json:"conference_auto_connect_next_speakers"`
	ConferenceEnableHelpdesk                     *bool           `json:"conference_enable_helpdesk"`
	ConferenceLosRestriction                     *bool           `json:"conference_los_restriction"`
	ConferenceOpenMicrophone                     *bool           `json:"conference_open_microphone"`
	ConferenceOpenVideo                          *bool           `json:"conference_open_video"`
	ConferenceShow                               *bool           `json:"conference_show"`
	ConferenceStreamPosterUrl                    *string         `json:"conference_stream_poster_url"`
	ConferenceStreamUrl                          *string         `json:"conference_stream_url"`
	CustomTranslations                           json.RawMessage `json:"custom_translations"`
	DefaultGroupID                               int             `json:"default_group_id"`
	DefaultMeetingForCommitteeID                 *int            `json:"default_meeting_for_committee_id"`
	DefaultProjectorAgendaItemListIDs            []int           `json:"default_projector_agenda_item_list_ids"`
	DefaultProjectorAmendmentIDs                 []int           `json:"default_projector_amendment_ids"`
	DefaultProjectorAssignmentIDs                []int           `json:"default_projector_assignment_ids"`
	DefaultProjectorAssignmentPollIDs            []int           `json:"default_projector_assignment_poll_ids"`
	DefaultProjectorCountdownIDs                 []int           `json:"default_projector_countdown_ids"`
	DefaultProjectorCurrentListOfSpeakersIDs     []int           `json:"default_projector_current_list_of_speakers_ids"`
	DefaultProjectorListOfSpeakersIDs            []int           `json:"default_projector_list_of_speakers_ids"`
	DefaultProjectorMediafileIDs                 []int           `json:"default_projector_mediafile_ids"`
	DefaultProjectorMessageIDs                   []int           `json:"default_projector_message_ids"`
	DefaultProjectorMotionBlockIDs               []int           `json:"default_projector_motion_block_ids"`
	DefaultProjectorMotionIDs                    []int           `json:"default_projector_motion_ids"`
	DefaultProjectorMotionPollIDs                []int           `json:"default_projector_motion_poll_ids"`
	DefaultProjectorPollIDs                      []int           `json:"default_projector_poll_ids"`
	DefaultProjectorTopicIDs                     []int           `json:"default_projector_topic_ids"`
	Description                                  *string         `json:"description"`
	EnableAnonymous                              *bool           `json:"enable_anonymous"`
	EndTime                                      *int            `json:"end_time"`
	ExportCsvEncoding                            *string         `json:"export_csv_encoding"`
	ExportCsvSeparator                           *string         `json:"export_csv_separator"`
	ExportPdfFontsize                            *int            `json:"export_pdf_fontsize"`
	ExportPdfLineHeight                          *float32        `json:"export_pdf_line_height"`
	ExportPdfPageMarginBottom                    *int            `json:"export_pdf_page_margin_bottom"`
	ExportPdfPageMarginLeft                      *int            `json:"export_pdf_page_margin_left"`
	ExportPdfPageMarginRight                     *int            `json:"export_pdf_page_margin_right"`
	ExportPdfPageMarginTop                       *int            `json:"export_pdf_page_margin_top"`
	ExportPdfPagenumberAlignment                 *string         `json:"export_pdf_pagenumber_alignment"`
	ExportPdfPagesize                            *string         `json:"export_pdf_pagesize"`
	ExternalID                                   *string         `json:"external_id"`
	FontBoldID                                   *int            `json:"font_bold_id"`
	FontBoldItalicID                             *int            `json:"font_bold_italic_id"`
	FontChyronSpeakerNameID                      *int            `json:"font_chyron_speaker_name_id"`
	FontItalicID                                 *int            `json:"font_italic_id"`
	FontMonospaceID                              *int            `json:"font_monospace_id"`
	FontProjectorH1ID                            *int            `json:"font_projector_h1_id"`
	FontProjectorH2ID                            *int            `json:"font_projector_h2_id"`
	FontRegularID                                *int            `json:"font_regular_id"`
	ForwardedMotionIDs                           []int           `json:"forwarded_motion_ids"`
	GroupIDs                                     []int           `json:"group_ids"`
	ID                                           int             `json:"id"`
	ImportedAt                                   *int            `json:"imported_at"`
	IsActiveInOrganizationID                     *int            `json:"is_active_in_organization_id"`
	IsArchivedInOrganizationID                   *int            `json:"is_archived_in_organization_id"`
	JitsiDomain                                  *string         `json:"jitsi_domain"`
	JitsiRoomName                                *string         `json:"jitsi_room_name"`
	JitsiRoomPassword                            *string         `json:"jitsi_room_password"`
	Language                                     string          `json:"language"`
	ListOfSpeakersAllowMultipleSpeakers          *bool           `json:"list_of_speakers_allow_multiple_speakers"`
	ListOfSpeakersAmountLastOnProjector          *int            `json:"list_of_speakers_amount_last_on_projector"`
	ListOfSpeakersAmountNextOnProjector          *int            `json:"list_of_speakers_amount_next_on_projector"`
	ListOfSpeakersCanCreatePointOfOrderForOthers *bool           `json:"list_of_speakers_can_create_point_of_order_for_others"`
	ListOfSpeakersCanSetContributionSelf         *bool           `json:"list_of_speakers_can_set_contribution_self"`
	ListOfSpeakersClosingDisablesPointOfOrder    *bool           `json:"list_of_speakers_closing_disables_point_of_order"`
	ListOfSpeakersCountdownID                    *int            `json:"list_of_speakers_countdown_id"`
	ListOfSpeakersCoupleCountdown                *bool           `json:"list_of_speakers_couple_countdown"`
	ListOfSpeakersDefaultStructureLevelTime      *int            `json:"list_of_speakers_default_structure_level_time"`
	ListOfSpeakersEnableInterposedQuestion       *bool           `json:"list_of_speakers_enable_interposed_question"`
	ListOfSpeakersEnablePointOfOrderCategories   *bool           `json:"list_of_speakers_enable_point_of_order_categories"`
	ListOfSpeakersEnablePointOfOrderSpeakers     *bool           `json:"list_of_speakers_enable_point_of_order_speakers"`
	ListOfSpeakersEnableProContraSpeech          *bool           `json:"list_of_speakers_enable_pro_contra_speech"`
	ListOfSpeakersHideContributionCount          *bool           `json:"list_of_speakers_hide_contribution_count"`
	ListOfSpeakersIDs                            []int           `json:"list_of_speakers_ids"`
	ListOfSpeakersInitiallyClosed                *bool           `json:"list_of_speakers_initially_closed"`
	ListOfSpeakersInterventionTime               *int            `json:"list_of_speakers_intervention_time"`
	ListOfSpeakersPresentUsersOnly               *bool           `json:"list_of_speakers_present_users_only"`
	ListOfSpeakersShowAmountOfSpeakersOnSlide    *bool           `json:"list_of_speakers_show_amount_of_speakers_on_slide"`
	ListOfSpeakersShowFirstContribution          *bool           `json:"list_of_speakers_show_first_contribution"`
	ListOfSpeakersSpeakerNoteForEveryone         *bool           `json:"list_of_speakers_speaker_note_for_everyone"`
	Location                                     *string         `json:"location"`
	LockedFromInside                             *bool           `json:"locked_from_inside"`
	LogoPdfBallotPaperID                         *int            `json:"logo_pdf_ballot_paper_id"`
	LogoPdfFooterLID                             *int            `json:"logo_pdf_footer_l_id"`
	LogoPdfFooterRID                             *int            `json:"logo_pdf_footer_r_id"`
	LogoPdfHeaderLID                             *int            `json:"logo_pdf_header_l_id"`
	LogoPdfHeaderRID                             *int            `json:"logo_pdf_header_r_id"`
	LogoProjectorHeaderID                        *int            `json:"logo_projector_header_id"`
	LogoProjectorMainID                          *int            `json:"logo_projector_main_id"`
	LogoWebHeaderID                              *int            `json:"logo_web_header_id"`
	MediafileIDs                                 []int           `json:"mediafile_ids"`
	MeetingMediafileIDs                          []int           `json:"meeting_mediafile_ids"`
	MeetingUserIDs                               []int           `json:"meeting_user_ids"`
	MotionBlockIDs                               []int           `json:"motion_block_ids"`
	MotionCategoryIDs                            []int           `json:"motion_category_ids"`
	MotionChangeRecommendationIDs                []int           `json:"motion_change_recommendation_ids"`
	MotionCommentIDs                             []int           `json:"motion_comment_ids"`
	MotionCommentSectionIDs                      []int           `json:"motion_comment_section_ids"`
	MotionEditorIDs                              []int           `json:"motion_editor_ids"`
	MotionIDs                                    []int           `json:"motion_ids"`
	MotionPollBallotPaperNumber                  *int            `json:"motion_poll_ballot_paper_number"`
	MotionPollBallotPaperSelection               *string         `json:"motion_poll_ballot_paper_selection"`
	MotionPollDefaultBackend                     *string         `json:"motion_poll_default_backend"`
	MotionPollDefaultGroupIDs                    []int           `json:"motion_poll_default_group_ids"`
	MotionPollDefaultMethod                      *string         `json:"motion_poll_default_method"`
	MotionPollDefaultOnehundredPercentBase       *string         `json:"motion_poll_default_onehundred_percent_base"`
	MotionPollDefaultType                        *string         `json:"motion_poll_default_type"`
	MotionStateIDs                               []int           `json:"motion_state_ids"`
	MotionStatuteParagraphIDs                    []int           `json:"motion_statute_paragraph_ids"`
	MotionSubmitterIDs                           []int           `json:"motion_submitter_ids"`
	MotionWorkflowIDs                            []int           `json:"motion_workflow_ids"`
	MotionWorkingGroupSpeakerIDs                 []int           `json:"motion_working_group_speaker_ids"`
	MotionsAmendmentsEnabled                     *bool           `json:"motions_amendments_enabled"`
	MotionsAmendmentsInMainList                  *bool           `json:"motions_amendments_in_main_list"`
	MotionsAmendmentsMultipleParagraphs          *bool           `json:"motions_amendments_multiple_paragraphs"`
	MotionsAmendmentsOfAmendments                *bool           `json:"motions_amendments_of_amendments"`
	MotionsAmendmentsPrefix                      *string         `json:"motions_amendments_prefix"`
	MotionsAmendmentsTextMode                    *string         `json:"motions_amendments_text_mode"`
	MotionsBlockSlideColumns                     *int            `json:"motions_block_slide_columns"`
	MotionsDefaultAmendmentWorkflowID            int             `json:"motions_default_amendment_workflow_id"`
	MotionsDefaultLineNumbering                  *string         `json:"motions_default_line_numbering"`
	MotionsDefaultSorting                        *string         `json:"motions_default_sorting"`
	MotionsDefaultStatuteAmendmentWorkflowID     int             `json:"motions_default_statute_amendment_workflow_id"`
	MotionsDefaultWorkflowID                     int             `json:"motions_default_workflow_id"`
	MotionsEnableEditor                          *bool           `json:"motions_enable_editor"`
	MotionsEnableReasonOnProjector               *bool           `json:"motions_enable_reason_on_projector"`
	MotionsEnableRecommendationOnProjector       *bool           `json:"motions_enable_recommendation_on_projector"`
	MotionsEnableSideboxOnProjector              *bool           `json:"motions_enable_sidebox_on_projector"`
	MotionsEnableTextOnProjector                 *bool           `json:"motions_enable_text_on_projector"`
	MotionsEnableWorkingGroupSpeaker             *bool           `json:"motions_enable_working_group_speaker"`
	MotionsExportFollowRecommendation            *bool           `json:"motions_export_follow_recommendation"`
	MotionsExportPreamble                        *string         `json:"motions_export_preamble"`
	MotionsExportSubmitterRecommendation         *bool           `json:"motions_export_submitter_recommendation"`
	MotionsExportTitle                           *string         `json:"motions_export_title"`
	MotionsHideMetadataBackground                *bool           `json:"motions_hide_metadata_background"`
	MotionsLineLength                            *int            `json:"motions_line_length"`
	MotionsNumberMinDigits                       *int            `json:"motions_number_min_digits"`
	MotionsNumberType                            *string         `json:"motions_number_type"`
	MotionsNumberWithBlank                       *bool           `json:"motions_number_with_blank"`
	MotionsPreamble                              *string         `json:"motions_preamble"`
	MotionsReasonRequired                        *bool           `json:"motions_reason_required"`
	MotionsRecommendationTextMode                *string         `json:"motions_recommendation_text_mode"`
	MotionsRecommendationsBy                     *string         `json:"motions_recommendations_by"`
	MotionsShowReferringMotions                  *bool           `json:"motions_show_referring_motions"`
	MotionsShowSequentialNumber                  *bool           `json:"motions_show_sequential_number"`
	MotionsStatuteRecommendationsBy              *string         `json:"motions_statute_recommendations_by"`
	MotionsStatutesEnabled                       *bool           `json:"motions_statutes_enabled"`
	MotionsSupportersMinAmount                   *int            `json:"motions_supporters_min_amount"`
	Name                                         string          `json:"name"`
	OptionIDs                                    []int           `json:"option_ids"`
	OrganizationTagIDs                           []int           `json:"organization_tag_ids"`
	PersonalNoteIDs                              []int           `json:"personal_note_ids"`
	PointOfOrderCategoryIDs                      []int           `json:"point_of_order_category_ids"`
	PollBallotPaperNumber                        *int            `json:"poll_ballot_paper_number"`
	PollBallotPaperSelection                     *string         `json:"poll_ballot_paper_selection"`
	PollCandidateIDs                             []int           `json:"poll_candidate_ids"`
	PollCandidateListIDs                         []int           `json:"poll_candidate_list_ids"`
	PollCountdownID                              *int            `json:"poll_countdown_id"`
	PollCoupleCountdown                          *bool           `json:"poll_couple_countdown"`
	PollDefaultBackend                           *string         `json:"poll_default_backend"`
	PollDefaultGroupIDs                          []int           `json:"poll_default_group_ids"`
	PollDefaultMethod                            *string         `json:"poll_default_method"`
	PollDefaultOnehundredPercentBase             *string         `json:"poll_default_onehundred_percent_base"`
	PollDefaultType                              *string         `json:"poll_default_type"`
	PollIDs                                      []int           `json:"poll_ids"`
	PollSortPollResultByVotes                    *bool           `json:"poll_sort_poll_result_by_votes"`
	PresentUserIDs                               []int           `json:"present_user_ids"`
	ProjectionIDs                                []int           `json:"projection_ids"`
	ProjectorCountdownDefaultTime                int             `json:"projector_countdown_default_time"`
	ProjectorCountdownIDs                        []int           `json:"projector_countdown_ids"`
	ProjectorCountdownWarningTime                int             `json:"projector_countdown_warning_time"`
	ProjectorIDs                                 []int           `json:"projector_ids"`
	ProjectorMessageIDs                          []int           `json:"projector_message_ids"`
	ReferenceProjectorID                         int             `json:"reference_projector_id"`
	SpeakerIDs                                   []int           `json:"speaker_ids"`
	StartTime                                    *int            `json:"start_time"`
	StructureLevelIDs                            []int           `json:"structure_level_ids"`
	StructureLevelListOfSpeakersIDs              []int           `json:"structure_level_list_of_speakers_ids"`
	TagIDs                                       []int           `json:"tag_ids"`
	TemplateForOrganizationID                    *int            `json:"template_for_organization_id"`
	TopicIDs                                     []int           `json:"topic_ids"`
	TopicPollDefaultGroupIDs                     []int           `json:"topic_poll_default_group_ids"`
	UserIDs                                      []int           `json:"user_ids"`
	UsersAllowSelfSetPresent                     *bool           `json:"users_allow_self_set_present"`
	UsersEmailBody                               *string         `json:"users_email_body"`
	UsersEmailReplyto                            *string         `json:"users_email_replyto"`
	UsersEmailSender                             *string         `json:"users_email_sender"`
	UsersEmailSubject                            *string         `json:"users_email_subject"`
	UsersEnablePresenceView                      *bool           `json:"users_enable_presence_view"`
	UsersEnableVoteDelegations                   *bool           `json:"users_enable_vote_delegations"`
	UsersEnableVoteWeight                        *bool           `json:"users_enable_vote_weight"`
	UsersForbidDelegatorAsSubmitter              *bool           `json:"users_forbid_delegator_as_submitter"`
	UsersForbidDelegatorAsSupporter              *bool           `json:"users_forbid_delegator_as_supporter"`
	UsersForbidDelegatorInListOfSpeakers         *bool           `json:"users_forbid_delegator_in_list_of_speakers"`
	UsersForbidDelegatorToVote                   *bool           `json:"users_forbid_delegator_to_vote"`
	UsersPdfWelcometext                          *string         `json:"users_pdf_welcometext"`
	UsersPdfWelcometitle                         *string         `json:"users_pdf_welcometitle"`
	UsersPdfWlanEncryption                       *string         `json:"users_pdf_wlan_encryption"`
	UsersPdfWlanPassword                         *string         `json:"users_pdf_wlan_password"`
	UsersPdfWlanSsid                             *string         `json:"users_pdf_wlan_ssid"`
	VoteIDs                                      []int           `json:"vote_ids"`
	WelcomeText                                  *string         `json:"welcome_text"`
	WelcomeTitle                                 *string         `json:"welcome_title"`
}

func (m Meeting) CollectionName() string {
	return "meeting"
}

func (m Meeting) Get(field string) interface{} {
	switch field {
	case "admin_group_id":
		return m.AdminGroupID
	case "agenda_enable_numbering":
		return m.AgendaEnableNumbering
	case "agenda_item_creation":
		return m.AgendaItemCreation
	case "agenda_item_ids":
		return m.AgendaItemIDs
	case "agenda_new_items_default_visibility":
		return m.AgendaNewItemsDefaultVisibility
	case "agenda_number_prefix":
		return m.AgendaNumberPrefix
	case "agenda_numeral_system":
		return m.AgendaNumeralSystem
	case "agenda_show_internal_items_on_projector":
		return m.AgendaShowInternalItemsOnProjector
	case "agenda_show_subtitles":
		return m.AgendaShowSubtitles
	case "agenda_show_topic_navigation_on_detail_view":
		return m.AgendaShowTopicNavigationOnDetailView
	case "all_projection_ids":
		return m.AllProjectionIDs
	case "anonymous_group_id":
		return m.AnonymousGroupID
	case "applause_enable":
		return m.ApplauseEnable
	case "applause_max_amount":
		return m.ApplauseMaxAmount
	case "applause_min_amount":
		return m.ApplauseMinAmount
	case "applause_particle_image_url":
		return m.ApplauseParticleImageUrl
	case "applause_show_level":
		return m.ApplauseShowLevel
	case "applause_timeout":
		return m.ApplauseTimeout
	case "applause_type":
		return m.ApplauseType
	case "assignment_candidate_ids":
		return m.AssignmentCandidateIDs
	case "assignment_ids":
		return m.AssignmentIDs
	case "assignment_poll_add_candidates_to_list_of_speakers":
		return m.AssignmentPollAddCandidatesToListOfSpeakers
	case "assignment_poll_ballot_paper_number":
		return m.AssignmentPollBallotPaperNumber
	case "assignment_poll_ballot_paper_selection":
		return m.AssignmentPollBallotPaperSelection
	case "assignment_poll_default_backend":
		return m.AssignmentPollDefaultBackend
	case "assignment_poll_default_group_ids":
		return m.AssignmentPollDefaultGroupIDs
	case "assignment_poll_default_method":
		return m.AssignmentPollDefaultMethod
	case "assignment_poll_default_onehundred_percent_base":
		return m.AssignmentPollDefaultOnehundredPercentBase
	case "assignment_poll_default_type":
		return m.AssignmentPollDefaultType
	case "assignment_poll_enable_max_votes_per_option":
		return m.AssignmentPollEnableMaxVotesPerOption
	case "assignment_poll_sort_poll_result_by_votes":
		return m.AssignmentPollSortPollResultByVotes
	case "assignments_export_preamble":
		return m.AssignmentsExportPreamble
	case "assignments_export_title":
		return m.AssignmentsExportTitle
	case "chat_group_ids":
		return m.ChatGroupIDs
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "committee_id":
		return m.CommitteeID
	case "conference_auto_connect":
		return m.ConferenceAutoConnect
	case "conference_auto_connect_next_speakers":
		return m.ConferenceAutoConnectNextSpeakers
	case "conference_enable_helpdesk":
		return m.ConferenceEnableHelpdesk
	case "conference_los_restriction":
		return m.ConferenceLosRestriction
	case "conference_open_microphone":
		return m.ConferenceOpenMicrophone
	case "conference_open_video":
		return m.ConferenceOpenVideo
	case "conference_show":
		return m.ConferenceShow
	case "conference_stream_poster_url":
		return m.ConferenceStreamPosterUrl
	case "conference_stream_url":
		return m.ConferenceStreamUrl
	case "custom_translations":
		return m.CustomTranslations
	case "default_group_id":
		return m.DefaultGroupID
	case "default_meeting_for_committee_id":
		return m.DefaultMeetingForCommitteeID
	case "default_projector_agenda_item_list_ids":
		return m.DefaultProjectorAgendaItemListIDs
	case "default_projector_amendment_ids":
		return m.DefaultProjectorAmendmentIDs
	case "default_projector_assignment_ids":
		return m.DefaultProjectorAssignmentIDs
	case "default_projector_assignment_poll_ids":
		return m.DefaultProjectorAssignmentPollIDs
	case "default_projector_countdown_ids":
		return m.DefaultProjectorCountdownIDs
	case "default_projector_current_list_of_speakers_ids":
		return m.DefaultProjectorCurrentListOfSpeakersIDs
	case "default_projector_list_of_speakers_ids":
		return m.DefaultProjectorListOfSpeakersIDs
	case "default_projector_mediafile_ids":
		return m.DefaultProjectorMediafileIDs
	case "default_projector_message_ids":
		return m.DefaultProjectorMessageIDs
	case "default_projector_motion_block_ids":
		return m.DefaultProjectorMotionBlockIDs
	case "default_projector_motion_ids":
		return m.DefaultProjectorMotionIDs
	case "default_projector_motion_poll_ids":
		return m.DefaultProjectorMotionPollIDs
	case "default_projector_poll_ids":
		return m.DefaultProjectorPollIDs
	case "default_projector_topic_ids":
		return m.DefaultProjectorTopicIDs
	case "description":
		return m.Description
	case "enable_anonymous":
		return m.EnableAnonymous
	case "end_time":
		return m.EndTime
	case "export_csv_encoding":
		return m.ExportCsvEncoding
	case "export_csv_separator":
		return m.ExportCsvSeparator
	case "export_pdf_fontsize":
		return m.ExportPdfFontsize
	case "export_pdf_line_height":
		return m.ExportPdfLineHeight
	case "export_pdf_page_margin_bottom":
		return m.ExportPdfPageMarginBottom
	case "export_pdf_page_margin_left":
		return m.ExportPdfPageMarginLeft
	case "export_pdf_page_margin_right":
		return m.ExportPdfPageMarginRight
	case "export_pdf_page_margin_top":
		return m.ExportPdfPageMarginTop
	case "export_pdf_pagenumber_alignment":
		return m.ExportPdfPagenumberAlignment
	case "export_pdf_pagesize":
		return m.ExportPdfPagesize
	case "external_id":
		return m.ExternalID
	case "font_bold_id":
		return m.FontBoldID
	case "font_bold_italic_id":
		return m.FontBoldItalicID
	case "font_chyron_speaker_name_id":
		return m.FontChyronSpeakerNameID
	case "font_italic_id":
		return m.FontItalicID
	case "font_monospace_id":
		return m.FontMonospaceID
	case "font_projector_h1_id":
		return m.FontProjectorH1ID
	case "font_projector_h2_id":
		return m.FontProjectorH2ID
	case "font_regular_id":
		return m.FontRegularID
	case "forwarded_motion_ids":
		return m.ForwardedMotionIDs
	case "group_ids":
		return m.GroupIDs
	case "id":
		return m.ID
	case "imported_at":
		return m.ImportedAt
	case "is_active_in_organization_id":
		return m.IsActiveInOrganizationID
	case "is_archived_in_organization_id":
		return m.IsArchivedInOrganizationID
	case "jitsi_domain":
		return m.JitsiDomain
	case "jitsi_room_name":
		return m.JitsiRoomName
	case "jitsi_room_password":
		return m.JitsiRoomPassword
	case "language":
		return m.Language
	case "list_of_speakers_allow_multiple_speakers":
		return m.ListOfSpeakersAllowMultipleSpeakers
	case "list_of_speakers_amount_last_on_projector":
		return m.ListOfSpeakersAmountLastOnProjector
	case "list_of_speakers_amount_next_on_projector":
		return m.ListOfSpeakersAmountNextOnProjector
	case "list_of_speakers_can_create_point_of_order_for_others":
		return m.ListOfSpeakersCanCreatePointOfOrderForOthers
	case "list_of_speakers_can_set_contribution_self":
		return m.ListOfSpeakersCanSetContributionSelf
	case "list_of_speakers_closing_disables_point_of_order":
		return m.ListOfSpeakersClosingDisablesPointOfOrder
	case "list_of_speakers_countdown_id":
		return m.ListOfSpeakersCountdownID
	case "list_of_speakers_couple_countdown":
		return m.ListOfSpeakersCoupleCountdown
	case "list_of_speakers_default_structure_level_time":
		return m.ListOfSpeakersDefaultStructureLevelTime
	case "list_of_speakers_enable_interposed_question":
		return m.ListOfSpeakersEnableInterposedQuestion
	case "list_of_speakers_enable_point_of_order_categories":
		return m.ListOfSpeakersEnablePointOfOrderCategories
	case "list_of_speakers_enable_point_of_order_speakers":
		return m.ListOfSpeakersEnablePointOfOrderSpeakers
	case "list_of_speakers_enable_pro_contra_speech":
		return m.ListOfSpeakersEnableProContraSpeech
	case "list_of_speakers_hide_contribution_count":
		return m.ListOfSpeakersHideContributionCount
	case "list_of_speakers_ids":
		return m.ListOfSpeakersIDs
	case "list_of_speakers_initially_closed":
		return m.ListOfSpeakersInitiallyClosed
	case "list_of_speakers_intervention_time":
		return m.ListOfSpeakersInterventionTime
	case "list_of_speakers_present_users_only":
		return m.ListOfSpeakersPresentUsersOnly
	case "list_of_speakers_show_amount_of_speakers_on_slide":
		return m.ListOfSpeakersShowAmountOfSpeakersOnSlide
	case "list_of_speakers_show_first_contribution":
		return m.ListOfSpeakersShowFirstContribution
	case "list_of_speakers_speaker_note_for_everyone":
		return m.ListOfSpeakersSpeakerNoteForEveryone
	case "location":
		return m.Location
	case "locked_from_inside":
		return m.LockedFromInside
	case "logo_pdf_ballot_paper_id":
		return m.LogoPdfBallotPaperID
	case "logo_pdf_footer_l_id":
		return m.LogoPdfFooterLID
	case "logo_pdf_footer_r_id":
		return m.LogoPdfFooterRID
	case "logo_pdf_header_l_id":
		return m.LogoPdfHeaderLID
	case "logo_pdf_header_r_id":
		return m.LogoPdfHeaderRID
	case "logo_projector_header_id":
		return m.LogoProjectorHeaderID
	case "logo_projector_main_id":
		return m.LogoProjectorMainID
	case "logo_web_header_id":
		return m.LogoWebHeaderID
	case "mediafile_ids":
		return m.MediafileIDs
	case "meeting_mediafile_ids":
		return m.MeetingMediafileIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "motion_block_ids":
		return m.MotionBlockIDs
	case "motion_category_ids":
		return m.MotionCategoryIDs
	case "motion_change_recommendation_ids":
		return m.MotionChangeRecommendationIDs
	case "motion_comment_ids":
		return m.MotionCommentIDs
	case "motion_comment_section_ids":
		return m.MotionCommentSectionIDs
	case "motion_editor_ids":
		return m.MotionEditorIDs
	case "motion_ids":
		return m.MotionIDs
	case "motion_poll_ballot_paper_number":
		return m.MotionPollBallotPaperNumber
	case "motion_poll_ballot_paper_selection":
		return m.MotionPollBallotPaperSelection
	case "motion_poll_default_backend":
		return m.MotionPollDefaultBackend
	case "motion_poll_default_group_ids":
		return m.MotionPollDefaultGroupIDs
	case "motion_poll_default_method":
		return m.MotionPollDefaultMethod
	case "motion_poll_default_onehundred_percent_base":
		return m.MotionPollDefaultOnehundredPercentBase
	case "motion_poll_default_type":
		return m.MotionPollDefaultType
	case "motion_state_ids":
		return m.MotionStateIDs
	case "motion_statute_paragraph_ids":
		return m.MotionStatuteParagraphIDs
	case "motion_submitter_ids":
		return m.MotionSubmitterIDs
	case "motion_workflow_ids":
		return m.MotionWorkflowIDs
	case "motion_working_group_speaker_ids":
		return m.MotionWorkingGroupSpeakerIDs
	case "motions_amendments_enabled":
		return m.MotionsAmendmentsEnabled
	case "motions_amendments_in_main_list":
		return m.MotionsAmendmentsInMainList
	case "motions_amendments_multiple_paragraphs":
		return m.MotionsAmendmentsMultipleParagraphs
	case "motions_amendments_of_amendments":
		return m.MotionsAmendmentsOfAmendments
	case "motions_amendments_prefix":
		return m.MotionsAmendmentsPrefix
	case "motions_amendments_text_mode":
		return m.MotionsAmendmentsTextMode
	case "motions_block_slide_columns":
		return m.MotionsBlockSlideColumns
	case "motions_default_amendment_workflow_id":
		return m.MotionsDefaultAmendmentWorkflowID
	case "motions_default_line_numbering":
		return m.MotionsDefaultLineNumbering
	case "motions_default_sorting":
		return m.MotionsDefaultSorting
	case "motions_default_statute_amendment_workflow_id":
		return m.MotionsDefaultStatuteAmendmentWorkflowID
	case "motions_default_workflow_id":
		return m.MotionsDefaultWorkflowID
	case "motions_enable_editor":
		return m.MotionsEnableEditor
	case "motions_enable_reason_on_projector":
		return m.MotionsEnableReasonOnProjector
	case "motions_enable_recommendation_on_projector":
		return m.MotionsEnableRecommendationOnProjector
	case "motions_enable_sidebox_on_projector":
		return m.MotionsEnableSideboxOnProjector
	case "motions_enable_text_on_projector":
		return m.MotionsEnableTextOnProjector
	case "motions_enable_working_group_speaker":
		return m.MotionsEnableWorkingGroupSpeaker
	case "motions_export_follow_recommendation":
		return m.MotionsExportFollowRecommendation
	case "motions_export_preamble":
		return m.MotionsExportPreamble
	case "motions_export_submitter_recommendation":
		return m.MotionsExportSubmitterRecommendation
	case "motions_export_title":
		return m.MotionsExportTitle
	case "motions_hide_metadata_background":
		return m.MotionsHideMetadataBackground
	case "motions_line_length":
		return m.MotionsLineLength
	case "motions_number_min_digits":
		return m.MotionsNumberMinDigits
	case "motions_number_type":
		return m.MotionsNumberType
	case "motions_number_with_blank":
		return m.MotionsNumberWithBlank
	case "motions_preamble":
		return m.MotionsPreamble
	case "motions_reason_required":
		return m.MotionsReasonRequired
	case "motions_recommendation_text_mode":
		return m.MotionsRecommendationTextMode
	case "motions_recommendations_by":
		return m.MotionsRecommendationsBy
	case "motions_show_referring_motions":
		return m.MotionsShowReferringMotions
	case "motions_show_sequential_number":
		return m.MotionsShowSequentialNumber
	case "motions_statute_recommendations_by":
		return m.MotionsStatuteRecommendationsBy
	case "motions_statutes_enabled":
		return m.MotionsStatutesEnabled
	case "motions_supporters_min_amount":
		return m.MotionsSupportersMinAmount
	case "name":
		return m.Name
	case "option_ids":
		return m.OptionIDs
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "point_of_order_category_ids":
		return m.PointOfOrderCategoryIDs
	case "poll_ballot_paper_number":
		return m.PollBallotPaperNumber
	case "poll_ballot_paper_selection":
		return m.PollBallotPaperSelection
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	case "poll_candidate_list_ids":
		return m.PollCandidateListIDs
	case "poll_countdown_id":
		return m.PollCountdownID
	case "poll_couple_countdown":
		return m.PollCoupleCountdown
	case "poll_default_backend":
		return m.PollDefaultBackend
	case "poll_default_group_ids":
		return m.PollDefaultGroupIDs
	case "poll_default_method":
		return m.PollDefaultMethod
	case "poll_default_onehundred_percent_base":
		return m.PollDefaultOnehundredPercentBase
	case "poll_default_type":
		return m.PollDefaultType
	case "poll_ids":
		return m.PollIDs
	case "poll_sort_poll_result_by_votes":
		return m.PollSortPollResultByVotes
	case "present_user_ids":
		return m.PresentUserIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "projector_countdown_default_time":
		return m.ProjectorCountdownDefaultTime
	case "projector_countdown_ids":
		return m.ProjectorCountdownIDs
	case "projector_countdown_warning_time":
		return m.ProjectorCountdownWarningTime
	case "projector_ids":
		return m.ProjectorIDs
	case "projector_message_ids":
		return m.ProjectorMessageIDs
	case "reference_projector_id":
		return m.ReferenceProjectorID
	case "speaker_ids":
		return m.SpeakerIDs
	case "start_time":
		return m.StartTime
	case "structure_level_ids":
		return m.StructureLevelIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	case "tag_ids":
		return m.TagIDs
	case "template_for_organization_id":
		return m.TemplateForOrganizationID
	case "topic_ids":
		return m.TopicIDs
	case "topic_poll_default_group_ids":
		return m.TopicPollDefaultGroupIDs
	case "user_ids":
		return m.UserIDs
	case "users_allow_self_set_present":
		return m.UsersAllowSelfSetPresent
	case "users_email_body":
		return m.UsersEmailBody
	case "users_email_replyto":
		return m.UsersEmailReplyto
	case "users_email_sender":
		return m.UsersEmailSender
	case "users_email_subject":
		return m.UsersEmailSubject
	case "users_enable_presence_view":
		return m.UsersEnablePresenceView
	case "users_enable_vote_delegations":
		return m.UsersEnableVoteDelegations
	case "users_enable_vote_weight":
		return m.UsersEnableVoteWeight
	case "users_forbid_delegator_as_submitter":
		return m.UsersForbidDelegatorAsSubmitter
	case "users_forbid_delegator_as_supporter":
		return m.UsersForbidDelegatorAsSupporter
	case "users_forbid_delegator_in_list_of_speakers":
		return m.UsersForbidDelegatorInListOfSpeakers
	case "users_forbid_delegator_to_vote":
		return m.UsersForbidDelegatorToVote
	case "users_pdf_welcometext":
		return m.UsersPdfWelcometext
	case "users_pdf_welcometitle":
		return m.UsersPdfWelcometitle
	case "users_pdf_wlan_encryption":
		return m.UsersPdfWlanEncryption
	case "users_pdf_wlan_password":
		return m.UsersPdfWlanPassword
	case "users_pdf_wlan_ssid":
		return m.UsersPdfWlanSsid
	case "vote_ids":
		return m.VoteIDs
	case "welcome_text":
		return m.WelcomeText
	case "welcome_title":
		return m.WelcomeTitle
	}

	return nil
}

func (m Meeting) Update(data map[string]string) error {
	if val, ok := data["admin_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdminGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_enable_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaEnableNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_creation"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemCreation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_new_items_default_visibility"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNewItemsDefaultVisibility)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_number_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumberPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_numeral_system"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumeralSystem)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_internal_items_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowInternalItemsOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_subtitles"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowSubtitles)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_topic_navigation_on_detail_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowTopicNavigationOnDetailView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["anonymous_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AnonymousGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_enable"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseEnable)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_max_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMaxAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMinAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_particle_image_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseParticleImageUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_show_level"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseShowLevel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_timeout"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseTimeout)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentCandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_add_candidates_to_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollAddCandidatesToListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_enable_max_votes_per_option"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollEnableMaxVotesPerOption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnect)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect_next_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnectNextSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_enable_helpdesk"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceEnableHelpdesk)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_los_restriction"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceLosRestriction)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_microphone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenMicrophone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_video"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenVideo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_show"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceShow)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_poster_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamPosterUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["custom_translations"]; ok {
		err := json.Unmarshal([]byte(val), &m.CustomTranslations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_meeting_for_committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultMeetingForCommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_agenda_item_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAgendaItemListIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_amendment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAmendmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_assignment_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCountdownIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_current_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCurrentListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionBlockIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorTopicIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
		if err != nil {
			return err
		}
	}

	if val, ok := data["enable_anonymous"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableAnonymous)
		if err != nil {
			return err
		}
	}

	if val, ok := data["end_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.EndTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_csv_encoding"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvEncoding)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_csv_separator"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvSeparator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_fontsize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfFontsize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_line_height"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfLineHeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_bottom"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginBottom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_left"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginLeft)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_right"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginRight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_top"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginTop)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagenumber_alignment"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagenumberAlignment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagesize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagesize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_chyron_speaker_name_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontChyronSpeakerNameID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_monospace_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontMonospaceID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h1_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH1ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h2_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH2ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_regular_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontRegularID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarded_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardedMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GroupIDs)
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

	if val, ok := data["imported_at"]; ok {
		err := json.Unmarshal([]byte(val), &m.ImportedAt)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_active_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsActiveInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_archived_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsArchivedInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_domain"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiDomain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_name"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomName)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["language"]; ok {
		err := json.Unmarshal([]byte(val), &m.Language)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_allow_multiple_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAllowMultipleSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_last_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountLastOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_next_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountNextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_create_point_of_order_for_others"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanCreatePointOfOrderForOthers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_set_contribution_self"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanSetContributionSelf)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_closing_disables_point_of_order"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersClosingDisablesPointOfOrder)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_default_structure_level_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersDefaultStructureLevelTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_interposed_question"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableInterposedQuestion)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_categories"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderCategories)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_pro_contra_speech"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableProContraSpeech)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_hide_contribution_count"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersHideContributionCount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_initially_closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInitiallyClosed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_intervention_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInterventionTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_present_users_only"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersPresentUsersOnly)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_amount_of_speakers_on_slide"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowAmountOfSpeakersOnSlide)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_first_contribution"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowFirstContribution)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_speaker_note_for_everyone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersSpeakerNoteForEveryone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["location"]; ok {
		err := json.Unmarshal([]byte(val), &m.Location)
		if err != nil {
			return err
		}
	}

	if val, ok := data["locked_from_inside"]; ok {
		err := json.Unmarshal([]byte(val), &m.LockedFromInside)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_ballot_paper_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfBallotPaperID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_main_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorMainID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_web_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoWebHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionBlockIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCategoryIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_change_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionChangeRecommendationIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentSectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionEditorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionStateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_statute_paragraph_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionStatuteParagraphIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionSubmitterIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_workflow_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkflowIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_enabled"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsEnabled)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_in_main_list"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsInMainList)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_multiple_paragraphs"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsMultipleParagraphs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_of_amendments"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsOfAmendments)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_block_slide_columns"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsBlockSlideColumns)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_amendment_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultAmendmentWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_line_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultLineNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_sorting"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultSorting)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_statute_amendment_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultStatuteAmendmentWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_editor"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableEditor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_reason_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableReasonOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_recommendation_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableRecommendationOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_sidebox_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableSideboxOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_text_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableTextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_working_group_speaker"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableWorkingGroupSpeaker)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_follow_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportFollowRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_submitter_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportSubmitterRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_hide_metadata_background"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsHideMetadataBackground)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_line_length"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsLineLength)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_min_digits"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberMinDigits)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_with_blank"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberWithBlank)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_reason_required"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsReasonRequired)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendation_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendations_by"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationsBy)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_referring_motions"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowReferringMotions)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowSequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_statute_recommendations_by"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsStatuteRecommendationsBy)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_statutes_enabled"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsStatutesEnabled)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_supporters_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsSupportersMinAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
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

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
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

	if val, ok := data["point_of_order_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrderCategoryIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateListIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultType)
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

	if val, ok := data["poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["present_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PresentUserIDs)
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

	if val, ok := data["projector_countdown_default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownDefaultTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_warning_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownWarningTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["reference_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferenceProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["start_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.StartTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
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

	if val, ok := data["template_for_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.TemplateForOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["topic_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_allow_self_set_present"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersAllowSelfSetPresent)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_body"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailBody)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_replyto"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailReplyto)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_sender"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSender)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_subject"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSubject)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_presence_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnablePresenceView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_delegations"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteDelegations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_submitter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSubmitter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_supporter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSupporter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_in_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorInListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_to_vote"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorToVote)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometext"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometext)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometitle"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_encryption"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanEncryption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_ssid"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanSsid)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["welcome_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["welcome_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeTitle)
		if err != nil {
			return err
		}
	}

	return nil
}
