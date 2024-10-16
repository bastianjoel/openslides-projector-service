package models

import "encoding/json"

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

func (m MeetingMediafile) Get(field string) interface{} {
	switch field {
	case "access_group_ids":
		return m.AccessGroupIDs
	case "attachment_ids":
		return m.AttachmentIDs
	case "id":
		return m.ID
	case "inherited_access_group_ids":
		return m.InheritedAccessGroupIDs
	case "is_public":
		return m.IsPublic
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "mediafile_id":
		return m.MediafileID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "used_as_font_bold_in_meeting_id":
		return m.UsedAsFontBoldInMeetingID
	case "used_as_font_bold_italic_in_meeting_id":
		return m.UsedAsFontBoldItalicInMeetingID
	case "used_as_font_chyron_speaker_name_in_meeting_id":
		return m.UsedAsFontChyronSpeakerNameInMeetingID
	case "used_as_font_italic_in_meeting_id":
		return m.UsedAsFontItalicInMeetingID
	case "used_as_font_monospace_in_meeting_id":
		return m.UsedAsFontMonospaceInMeetingID
	case "used_as_font_projector_h1_in_meeting_id":
		return m.UsedAsFontProjectorH1InMeetingID
	case "used_as_font_projector_h2_in_meeting_id":
		return m.UsedAsFontProjectorH2InMeetingID
	case "used_as_font_regular_in_meeting_id":
		return m.UsedAsFontRegularInMeetingID
	case "used_as_logo_pdf_ballot_paper_in_meeting_id":
		return m.UsedAsLogoPdfBallotPaperInMeetingID
	case "used_as_logo_pdf_footer_l_in_meeting_id":
		return m.UsedAsLogoPdfFooterLInMeetingID
	case "used_as_logo_pdf_footer_r_in_meeting_id":
		return m.UsedAsLogoPdfFooterRInMeetingID
	case "used_as_logo_pdf_header_l_in_meeting_id":
		return m.UsedAsLogoPdfHeaderLInMeetingID
	case "used_as_logo_pdf_header_r_in_meeting_id":
		return m.UsedAsLogoPdfHeaderRInMeetingID
	case "used_as_logo_projector_header_in_meeting_id":
		return m.UsedAsLogoProjectorHeaderInMeetingID
	case "used_as_logo_projector_main_in_meeting_id":
		return m.UsedAsLogoProjectorMainInMeetingID
	case "used_as_logo_web_header_in_meeting_id":
		return m.UsedAsLogoWebHeaderInMeetingID
	}

	return nil
}

func (m MeetingMediafile) Update(data map[string]string) error {
	if val, ok := data["access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccessGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["attachment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AttachmentIDs)
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

	if val, ok := data["inherited_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.InheritedAccessGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_public"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPublic)
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

	if val, ok := data["mediafile_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileID)
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

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_bold_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontBoldInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_bold_italic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontBoldItalicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_chyron_speaker_name_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontChyronSpeakerNameInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_italic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontItalicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_monospace_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontMonospaceInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_projector_h1_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontProjectorH1InMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_projector_h2_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontProjectorH2InMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_regular_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontRegularInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_ballot_paper_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfBallotPaperInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_footer_l_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfFooterLInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_footer_r_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfFooterRInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_header_l_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfHeaderLInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_header_r_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfHeaderRInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_projector_header_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoProjectorHeaderInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_projector_main_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoProjectorMainInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_web_header_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoWebHeaderInMeetingID)
		if err != nil {
			return err
		}
	}

	return nil
}
