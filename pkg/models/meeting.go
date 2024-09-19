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
	ID                                           *int            `json:"id"`
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
