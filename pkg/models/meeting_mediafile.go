package models

type MeetingMediafile struct {
	AccessGroupIDs                         []int    `json:"access_group_ids"`
	AttachmentIDs                          []string `json:"attachment_ids"`
	ID                                     int      `json:"id"`
	InheritedAccessGroupIDs                []int    `json:"inherited_access_group_ids"`
	IsPublic                               bool     `json:"is_public"`
	ListOfSpeakersID                       *int     `json:"list_of_speakers_id"`
	MediafileID                            int      `json:"mediafile_id"`
	MeetingID                              int      `json:"meeting_id"`
	ProjectionIDs                          []int    `json:"projection_ids"`
	UsedAsFontBoldInMeetingID              *int     `json:"used_as_font_bold_in_meeting_id"`
	UsedAsFontBoldItalicInMeetingID        *int     `json:"used_as_font_bold_italic_in_meeting_id"`
	UsedAsFontChyronSpeakerNameInMeetingID *int     `json:"used_as_font_chyron_speaker_name_in_meeting_id"`
	UsedAsFontItalicInMeetingID            *int     `json:"used_as_font_italic_in_meeting_id"`
	UsedAsFontMonospaceInMeetingID         *int     `json:"used_as_font_monospace_in_meeting_id"`
	UsedAsFontProjectorH1InMeetingID       *int     `json:"used_as_font_projector_h1_in_meeting_id"`
	UsedAsFontProjectorH2InMeetingID       *int     `json:"used_as_font_projector_h2_in_meeting_id"`
	UsedAsFontRegularInMeetingID           *int     `json:"used_as_font_regular_in_meeting_id"`
	UsedAsLogoPdfBallotPaperInMeetingID    *int     `json:"used_as_logo_pdf_ballot_paper_in_meeting_id"`
	UsedAsLogoPdfFooterLInMeetingID        *int     `json:"used_as_logo_pdf_footer_l_in_meeting_id"`
	UsedAsLogoPdfFooterRInMeetingID        *int     `json:"used_as_logo_pdf_footer_r_in_meeting_id"`
	UsedAsLogoPdfHeaderLInMeetingID        *int     `json:"used_as_logo_pdf_header_l_in_meeting_id"`
	UsedAsLogoPdfHeaderRInMeetingID        *int     `json:"used_as_logo_pdf_header_r_in_meeting_id"`
	UsedAsLogoProjectorHeaderInMeetingID   *int     `json:"used_as_logo_projector_header_in_meeting_id"`
	UsedAsLogoProjectorMainInMeetingID     *int     `json:"used_as_logo_projector_main_in_meeting_id"`
	UsedAsLogoWebHeaderInMeetingID         *int     `json:"used_as_logo_web_header_in_meeting_id"`
}

func (m MeetingMediafile) CollectionName() string {
	return "meeting_mediafile"
}
