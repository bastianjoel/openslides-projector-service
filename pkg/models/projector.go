package models

type Projector struct {
	AspectRatioDenominator                                    *int    `json:"aspect_ratio_denominator"`
	AspectRatioNumerator                                      *int    `json:"aspect_ratio_numerator"`
	BackgroundColor                                           *string `json:"background_color"`
	ChyronBackgroundColor                                     *string `json:"chyron_background_color"`
	ChyronBackgroundColor2                                    *string `json:"chyron_background_color_2"`
	ChyronFontColor                                           *string `json:"chyron_font_color"`
	ChyronFontColor2                                          *string `json:"chyron_font_color_2"`
	Color                                                     *string `json:"color"`
	CurrentProjectionIDs                                      []int   `json:"current_projection_ids"`
	HeaderBackgroundColor                                     *string `json:"header_background_color"`
	HeaderFontColor                                           *string `json:"header_font_color"`
	HeaderH1Color                                             *string `json:"header_h1_color"`
	HistoryProjectionIDs                                      []int   `json:"history_projection_ids"`
	ID                                                        int     `json:"id"`
	IsInternal                                                *bool   `json:"is_internal"`
	MeetingID                                                 int     `json:"meeting_id"`
	Name                                                      *string `json:"name"`
	PreviewProjectionIDs                                      []int   `json:"preview_projection_ids"`
	Scale                                                     *int    `json:"scale"`
	Scroll                                                    *int    `json:"scroll"`
	SequentialNumber                                          int     `json:"sequential_number"`
	ShowClock                                                 *bool   `json:"show_clock"`
	ShowHeaderFooter                                          *bool   `json:"show_header_footer"`
	ShowLogo                                                  *bool   `json:"show_logo"`
	ShowTitle                                                 *bool   `json:"show_title"`
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        *int    `json:"used_as_default_projector_for_agenda_item_list_in_meeting_id"`
	UsedAsDefaultProjectorForAmendmentInMeetingID             *int    `json:"used_as_default_projector_for_amendment_in_meeting_id"`
	UsedAsDefaultProjectorForAssignmentInMeetingID            *int    `json:"used_as_default_projector_for_assignment_in_meeting_id"`
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        *int    `json:"used_as_default_projector_for_assignment_poll_in_meeting_id"`
	UsedAsDefaultProjectorForCountdownInMeetingID             *int    `json:"used_as_default_projector_for_countdown_in_meeting_id"`
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID *int    `json:"used_as_default_projector_for_current_list_of_speakers_in_meeting_id"`
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        *int    `json:"used_as_default_projector_for_list_of_speakers_in_meeting_id"`
	UsedAsDefaultProjectorForMediafileInMeetingID             *int    `json:"used_as_default_projector_for_mediafile_in_meeting_id"`
	UsedAsDefaultProjectorForMessageInMeetingID               *int    `json:"used_as_default_projector_for_message_in_meeting_id"`
	UsedAsDefaultProjectorForMotionBlockInMeetingID           *int    `json:"used_as_default_projector_for_motion_block_in_meeting_id"`
	UsedAsDefaultProjectorForMotionInMeetingID                *int    `json:"used_as_default_projector_for_motion_in_meeting_id"`
	UsedAsDefaultProjectorForMotionPollInMeetingID            *int    `json:"used_as_default_projector_for_motion_poll_in_meeting_id"`
	UsedAsDefaultProjectorForPollInMeetingID                  *int    `json:"used_as_default_projector_for_poll_in_meeting_id"`
	UsedAsDefaultProjectorForTopicInMeetingID                 *int    `json:"used_as_default_projector_for_topic_in_meeting_id"`
	UsedAsReferenceProjectorMeetingID                         *int    `json:"used_as_reference_projector_meeting_id"`
	Width                                                     *int    `json:"width"`
}

func (m Projector) CollectionName() string {
	return "projector"
}
