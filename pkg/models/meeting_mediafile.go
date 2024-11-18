package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

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
	loadedRelations                        map[string]struct{}
	listOfSpeakers                         *ListOfSpeakers
	usedAsFontMonospaceInMeeting           *Meeting
	usedAsLogoPdfBallotPaperInMeeting      *Meeting
	usedAsLogoProjectorMainInMeeting       *Meeting
	accessGroups                           *Group
	usedAsFontBoldInMeeting                *Meeting
	usedAsFontChyronSpeakerNameInMeeting   *Meeting
	usedAsFontItalicInMeeting              *Meeting
	usedAsFontProjectorH1InMeeting         *Meeting
	usedAsFontRegularInMeeting             *Meeting
	inheritedAccessGroups                  *Group
	mediafile                              *Mediafile
	projections                            *Projection
	usedAsFontBoldItalicInMeeting          *Meeting
	usedAsFontProjectorH2InMeeting         *Meeting
	usedAsLogoPdfFooterRInMeeting          *Meeting
	usedAsLogoPdfHeaderRInMeeting          *Meeting
	usedAsLogoProjectorHeaderInMeeting     *Meeting
	usedAsLogoWebHeaderInMeeting           *Meeting
	meeting                                *Meeting
	usedAsLogoPdfFooterLInMeeting          *Meeting
	usedAsLogoPdfHeaderLInMeeting          *Meeting
}

func (m *MeetingMediafile) CollectionName() string {
	return "meeting_mediafile"
}

func (m *MeetingMediafile) ListOfSpeakers() *ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of MeetingMediafile which was not loaded.")
	}

	return m.listOfSpeakers
}

func (m *MeetingMediafile) UsedAsFontMonospaceInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_monospace_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontMonospaceInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontMonospaceInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfBallotPaperInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_ballot_paper_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfBallotPaperInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfBallotPaperInMeeting
}

func (m *MeetingMediafile) UsedAsLogoProjectorMainInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_projector_main_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoProjectorMainInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoProjectorMainInMeeting
}

func (m *MeetingMediafile) AccessGroups() *Group {
	if _, ok := m.loadedRelations["access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access AccessGroups relation of MeetingMediafile which was not loaded.")
	}

	return m.accessGroups
}

func (m *MeetingMediafile) UsedAsFontBoldInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_bold_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontBoldInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontBoldInMeeting
}

func (m *MeetingMediafile) UsedAsFontChyronSpeakerNameInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_chyron_speaker_name_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontChyronSpeakerNameInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontChyronSpeakerNameInMeeting
}

func (m *MeetingMediafile) UsedAsFontItalicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_italic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontItalicInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontItalicInMeeting
}

func (m *MeetingMediafile) UsedAsFontProjectorH1InMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_projector_h1_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontProjectorH1InMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontProjectorH1InMeeting
}

func (m *MeetingMediafile) UsedAsFontRegularInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_regular_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontRegularInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontRegularInMeeting
}

func (m *MeetingMediafile) InheritedAccessGroups() *Group {
	if _, ok := m.loadedRelations["inherited_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access InheritedAccessGroups relation of MeetingMediafile which was not loaded.")
	}

	return m.inheritedAccessGroups
}

func (m *MeetingMediafile) Mediafile() Mediafile {
	if _, ok := m.loadedRelations["mediafile_id"]; !ok {
		log.Panic().Msg("Tried to access Mediafile relation of MeetingMediafile which was not loaded.")
	}

	return *m.mediafile
}

func (m *MeetingMediafile) Projections() *Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of MeetingMediafile which was not loaded.")
	}

	return m.projections
}

func (m *MeetingMediafile) UsedAsFontBoldItalicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_bold_italic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontBoldItalicInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontBoldItalicInMeeting
}

func (m *MeetingMediafile) UsedAsFontProjectorH2InMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_projector_h2_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontProjectorH2InMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontProjectorH2InMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfFooterRInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_footer_r_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfFooterRInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfFooterRInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfHeaderRInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_header_r_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfHeaderRInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfHeaderRInMeeting
}

func (m *MeetingMediafile) UsedAsLogoProjectorHeaderInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_projector_header_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoProjectorHeaderInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoProjectorHeaderInMeeting
}

func (m *MeetingMediafile) UsedAsLogoWebHeaderInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_web_header_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoWebHeaderInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoWebHeaderInMeeting
}

func (m *MeetingMediafile) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MeetingMediafile which was not loaded.")
	}

	return *m.meeting
}

func (m *MeetingMediafile) UsedAsLogoPdfFooterLInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_footer_l_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfFooterLInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfFooterLInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfHeaderLInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_header_l_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfHeaderLInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfHeaderLInMeeting
}

func (m *MeetingMediafile) Get(field string) interface{} {
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

func (m *MeetingMediafile) Update(data map[string]string) error {
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
