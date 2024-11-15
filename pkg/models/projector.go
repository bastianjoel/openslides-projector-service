package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

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
	loadedRelations                                           map[string]struct{}
	usedAsDefaultProjectorForCountdownInMeeting               *Meeting
	usedAsDefaultProjectorForMotionPollInMeeting              *Meeting
	usedAsReferenceProjectorMeeting                           *Meeting
	meeting                                                   *Meeting
	previewProjections                                        *Projection
	usedAsDefaultProjectorForMotionBlockInMeeting             *Meeting
	usedAsDefaultProjectorForAssignmentPollInMeeting          *Meeting
	usedAsDefaultProjectorForPollInMeeting                    *Meeting
	usedAsDefaultProjectorForAmendmentInMeeting               *Meeting
	usedAsDefaultProjectorForMessageInMeeting                 *Meeting
	currentProjections                                        *Projection
	usedAsDefaultProjectorForMediafileInMeeting               *Meeting
	usedAsDefaultProjectorForMotionInMeeting                  *Meeting
	usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting   *Meeting
	usedAsDefaultProjectorForListOfSpeakersInMeeting          *Meeting
	historyProjections                                        *Projection
	usedAsDefaultProjectorForAgendaItemListInMeeting          *Meeting
	usedAsDefaultProjectorForAssignmentInMeeting              *Meeting
	usedAsDefaultProjectorForTopicInMeeting                   *Meeting
}

func (m Projector) CollectionName() string {
	return "projector"
}

func (m *Projector) UsedAsDefaultProjectorForCountdownInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_countdown_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForCountdownInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForCountdownInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMotionPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionPollInMeeting
}

func (m *Projector) UsedAsReferenceProjectorMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_reference_projector_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsReferenceProjectorMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsReferenceProjectorMeeting
}

func (m *Projector) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Projector which was not loaded.")
	}

	return *m.meeting
}

func (m *Projector) PreviewProjections() *Projection {
	if _, ok := m.loadedRelations["preview_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access PreviewProjections relation of Projector which was not loaded.")
	}

	return m.previewProjections
}

func (m *Projector) UsedAsDefaultProjectorForMotionBlockInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_block_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionBlockInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionBlockInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAssignmentPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_assignment_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAssignmentPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAssignmentPollInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForPollInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAmendmentInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_amendment_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAmendmentInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAmendmentInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMessageInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_message_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMessageInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMessageInMeeting
}

func (m *Projector) CurrentProjections() *Projection {
	if _, ok := m.loadedRelations["current_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access CurrentProjections relation of Projector which was not loaded.")
	}

	return m.currentProjections
}

func (m *Projector) UsedAsDefaultProjectorForMediafileInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_mediafile_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMediafileInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMediafileInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMotionInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForCurrentListOfSpeakersInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_current_list_of_speakers_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForCurrentListOfSpeakersInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForListOfSpeakersInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_list_of_speakers_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForListOfSpeakersInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForListOfSpeakersInMeeting
}

func (m *Projector) HistoryProjections() *Projection {
	if _, ok := m.loadedRelations["history_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access HistoryProjections relation of Projector which was not loaded.")
	}

	return m.historyProjections
}

func (m *Projector) UsedAsDefaultProjectorForAgendaItemListInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_agenda_item_list_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAgendaItemListInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAgendaItemListInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAssignmentInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_assignment_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAssignmentInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAssignmentInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForTopicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_topic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForTopicInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForTopicInMeeting
}

func (m Projector) Get(field string) interface{} {
	switch field {
	case "aspect_ratio_denominator":
		return m.AspectRatioDenominator
	case "aspect_ratio_numerator":
		return m.AspectRatioNumerator
	case "background_color":
		return m.BackgroundColor
	case "chyron_background_color":
		return m.ChyronBackgroundColor
	case "chyron_background_color_2":
		return m.ChyronBackgroundColor2
	case "chyron_font_color":
		return m.ChyronFontColor
	case "chyron_font_color_2":
		return m.ChyronFontColor2
	case "color":
		return m.Color
	case "current_projection_ids":
		return m.CurrentProjectionIDs
	case "header_background_color":
		return m.HeaderBackgroundColor
	case "header_font_color":
		return m.HeaderFontColor
	case "header_h1_color":
		return m.HeaderH1Color
	case "history_projection_ids":
		return m.HistoryProjectionIDs
	case "id":
		return m.ID
	case "is_internal":
		return m.IsInternal
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "preview_projection_ids":
		return m.PreviewProjectionIDs
	case "scale":
		return m.Scale
	case "scroll":
		return m.Scroll
	case "sequential_number":
		return m.SequentialNumber
	case "show_clock":
		return m.ShowClock
	case "show_header_footer":
		return m.ShowHeaderFooter
	case "show_logo":
		return m.ShowLogo
	case "show_title":
		return m.ShowTitle
	case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
		return m.UsedAsDefaultProjectorForAgendaItemListInMeetingID
	case "used_as_default_projector_for_amendment_in_meeting_id":
		return m.UsedAsDefaultProjectorForAmendmentInMeetingID
	case "used_as_default_projector_for_assignment_in_meeting_id":
		return m.UsedAsDefaultProjectorForAssignmentInMeetingID
	case "used_as_default_projector_for_assignment_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForAssignmentPollInMeetingID
	case "used_as_default_projector_for_countdown_in_meeting_id":
		return m.UsedAsDefaultProjectorForCountdownInMeetingID
	case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		return m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID
	case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
		return m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID
	case "used_as_default_projector_for_mediafile_in_meeting_id":
		return m.UsedAsDefaultProjectorForMediafileInMeetingID
	case "used_as_default_projector_for_message_in_meeting_id":
		return m.UsedAsDefaultProjectorForMessageInMeetingID
	case "used_as_default_projector_for_motion_block_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionBlockInMeetingID
	case "used_as_default_projector_for_motion_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionInMeetingID
	case "used_as_default_projector_for_motion_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionPollInMeetingID
	case "used_as_default_projector_for_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForPollInMeetingID
	case "used_as_default_projector_for_topic_in_meeting_id":
		return m.UsedAsDefaultProjectorForTopicInMeetingID
	case "used_as_reference_projector_meeting_id":
		return m.UsedAsReferenceProjectorMeetingID
	case "width":
		return m.Width
	}

	return nil
}

func (m Projector) Update(data map[string]string) error {
	if val, ok := data["aspect_ratio_denominator"]; ok {
		err := json.Unmarshal([]byte(val), &m.AspectRatioDenominator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["aspect_ratio_numerator"]; ok {
		err := json.Unmarshal([]byte(val), &m.AspectRatioNumerator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.BackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronBackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_background_color_2"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronBackgroundColor2)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_font_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronFontColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_font_color_2"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronFontColor2)
		if err != nil {
			return err
		}
	}

	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["current_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["header_background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderBackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["header_font_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderFontColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["header_h1_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderH1Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["history_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.HistoryProjectionIDs)
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

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["preview_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviewProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["scale"]; ok {
		err := json.Unmarshal([]byte(val), &m.Scale)
		if err != nil {
			return err
		}
	}

	if val, ok := data["scroll"]; ok {
		err := json.Unmarshal([]byte(val), &m.Scroll)
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

	if val, ok := data["show_clock"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowClock)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_header_footer"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowHeaderFooter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_logo"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowLogo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_agenda_item_list_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAgendaItemListInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_amendment_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAmendmentInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_assignment_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAssignmentInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_assignment_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAssignmentPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_countdown_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForCountdownInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_current_list_of_speakers_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_list_of_speakers_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_mediafile_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMediafileInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_message_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMessageInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_block_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionBlockInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_topic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForTopicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_reference_projector_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsReferenceProjectorMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["width"]; ok {
		err := json.Unmarshal([]byte(val), &m.Width)
		if err != nil {
			return err
		}
	}

	return nil
}
