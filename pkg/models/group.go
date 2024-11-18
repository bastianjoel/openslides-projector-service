package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Group struct {
	AdminGroupForMeetingID                  *int     `json:"admin_group_for_meeting_id"`
	AnonymousGroupForMeetingID              *int     `json:"anonymous_group_for_meeting_id"`
	DefaultGroupForMeetingID                *int     `json:"default_group_for_meeting_id"`
	ExternalID                              *string  `json:"external_id"`
	ID                                      int      `json:"id"`
	MeetingID                               int      `json:"meeting_id"`
	MeetingMediafileAccessGroupIDs          []int    `json:"meeting_mediafile_access_group_ids"`
	MeetingMediafileInheritedAccessGroupIDs []int    `json:"meeting_mediafile_inherited_access_group_ids"`
	MeetingUserIDs                          []int    `json:"meeting_user_ids"`
	Name                                    string   `json:"name"`
	Permissions                             []string `json:"permissions"`
	PollIDs                                 []int    `json:"poll_ids"`
	ReadChatGroupIDs                        []int    `json:"read_chat_group_ids"`
	ReadCommentSectionIDs                   []int    `json:"read_comment_section_ids"`
	UsedAsAssignmentPollDefaultID           *int     `json:"used_as_assignment_poll_default_id"`
	UsedAsMotionPollDefaultID               *int     `json:"used_as_motion_poll_default_id"`
	UsedAsPollDefaultID                     *int     `json:"used_as_poll_default_id"`
	UsedAsTopicPollDefaultID                *int     `json:"used_as_topic_poll_default_id"`
	Weight                                  *int     `json:"weight"`
	WriteChatGroupIDs                       []int    `json:"write_chat_group_ids"`
	WriteCommentSectionIDs                  []int    `json:"write_comment_section_ids"`
	loadedRelations                         map[string]struct{}
	writeChatGroups                         []ChatGroup
	writeCommentSections                    []MotionCommentSection
	adminGroupForMeeting                    *Meeting
	usedAsTopicPollDefault                  *Meeting
	defaultGroupForMeeting                  *Meeting
	meetingUsers                            []MeetingUser
	usedAsAssignmentPollDefault             *Meeting
	meetingMediafileAccessGroups            []MeetingMediafile
	usedAsPollDefault                       *Meeting
	readChatGroups                          []ChatGroup
	readCommentSections                     []MotionCommentSection
	usedAsMotionPollDefault                 *Meeting
	anonymousGroupForMeeting                *Meeting
	meeting                                 *Meeting
	meetingMediafileInheritedAccessGroups   []MeetingMediafile
	polls                                   []Poll
}

func (m *Group) CollectionName() string {
	return "group"
}

func (m *Group) WriteChatGroups() []ChatGroup {
	if _, ok := m.loadedRelations["write_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteChatGroups relation of Group which was not loaded.")
	}

	return m.writeChatGroups
}

func (m *Group) WriteCommentSections() []MotionCommentSection {
	if _, ok := m.loadedRelations["write_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteCommentSections relation of Group which was not loaded.")
	}

	return m.writeCommentSections
}

func (m *Group) AdminGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["admin_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access AdminGroupForMeeting relation of Group which was not loaded.")
	}

	return m.adminGroupForMeeting
}

func (m *Group) UsedAsTopicPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_topic_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsTopicPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsTopicPollDefault
}

func (m *Group) DefaultGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultGroupForMeeting relation of Group which was not loaded.")
	}

	return m.defaultGroupForMeeting
}

func (m *Group) MeetingUsers() []MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of Group which was not loaded.")
	}

	return m.meetingUsers
}

func (m *Group) UsedAsAssignmentPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_assignment_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsAssignmentPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsAssignmentPollDefault
}

func (m *Group) MeetingMediafileAccessGroups() []MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileAccessGroups
}

func (m *Group) UsedAsPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsPollDefault
}

func (m *Group) ReadChatGroups() []ChatGroup {
	if _, ok := m.loadedRelations["read_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadChatGroups relation of Group which was not loaded.")
	}

	return m.readChatGroups
}

func (m *Group) ReadCommentSections() []MotionCommentSection {
	if _, ok := m.loadedRelations["read_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadCommentSections relation of Group which was not loaded.")
	}

	return m.readCommentSections
}

func (m *Group) UsedAsMotionPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_motion_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsMotionPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsMotionPollDefault
}

func (m *Group) AnonymousGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["anonymous_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access AnonymousGroupForMeeting relation of Group which was not loaded.")
	}

	return m.anonymousGroupForMeeting
}

func (m *Group) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Group which was not loaded.")
	}

	return *m.meeting
}

func (m *Group) MeetingMediafileInheritedAccessGroups() []MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_inherited_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileInheritedAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileInheritedAccessGroups
}

func (m *Group) Polls() []Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Group which was not loaded.")
	}

	return m.polls
}

func (m *Group) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "write_chat_group_ids":
			m.writeChatGroups = content.([]ChatGroup)
		case "write_comment_section_ids":
			m.writeCommentSections = content.([]MotionCommentSection)
		case "admin_group_for_meeting_id":
			m.adminGroupForMeeting = content.(*Meeting)
		case "used_as_topic_poll_default_id":
			m.usedAsTopicPollDefault = content.(*Meeting)
		case "default_group_for_meeting_id":
			m.defaultGroupForMeeting = content.(*Meeting)
		case "meeting_user_ids":
			m.meetingUsers = content.([]MeetingUser)
		case "used_as_assignment_poll_default_id":
			m.usedAsAssignmentPollDefault = content.(*Meeting)
		case "meeting_mediafile_access_group_ids":
			m.meetingMediafileAccessGroups = content.([]MeetingMediafile)
		case "used_as_poll_default_id":
			m.usedAsPollDefault = content.(*Meeting)
		case "read_chat_group_ids":
			m.readChatGroups = content.([]ChatGroup)
		case "read_comment_section_ids":
			m.readCommentSections = content.([]MotionCommentSection)
		case "used_as_motion_poll_default_id":
			m.usedAsMotionPollDefault = content.(*Meeting)
		case "anonymous_group_for_meeting_id":
			m.anonymousGroupForMeeting = content.(*Meeting)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_mediafile_inherited_access_group_ids":
			m.meetingMediafileInheritedAccessGroups = content.([]MeetingMediafile)
		case "poll_ids":
			m.polls = content.([]Poll)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Group) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "write_chat_group_ids":
		err := json.Unmarshal(content, &m.writeChatGroups)
		if err != nil {
			return err
		}
	case "write_comment_section_ids":
		err := json.Unmarshal(content, &m.writeCommentSections)
		if err != nil {
			return err
		}
	case "admin_group_for_meeting_id":
		err := json.Unmarshal(content, &m.adminGroupForMeeting)
		if err != nil {
			return err
		}
	case "used_as_topic_poll_default_id":
		err := json.Unmarshal(content, &m.usedAsTopicPollDefault)
		if err != nil {
			return err
		}
	case "default_group_for_meeting_id":
		err := json.Unmarshal(content, &m.defaultGroupForMeeting)
		if err != nil {
			return err
		}
	case "meeting_user_ids":
		err := json.Unmarshal(content, &m.meetingUsers)
		if err != nil {
			return err
		}
	case "used_as_assignment_poll_default_id":
		err := json.Unmarshal(content, &m.usedAsAssignmentPollDefault)
		if err != nil {
			return err
		}
	case "meeting_mediafile_access_group_ids":
		err := json.Unmarshal(content, &m.meetingMediafileAccessGroups)
		if err != nil {
			return err
		}
	case "used_as_poll_default_id":
		err := json.Unmarshal(content, &m.usedAsPollDefault)
		if err != nil {
			return err
		}
	case "read_chat_group_ids":
		err := json.Unmarshal(content, &m.readChatGroups)
		if err != nil {
			return err
		}
	case "read_comment_section_ids":
		err := json.Unmarshal(content, &m.readCommentSections)
		if err != nil {
			return err
		}
	case "used_as_motion_poll_default_id":
		err := json.Unmarshal(content, &m.usedAsMotionPollDefault)
		if err != nil {
			return err
		}
	case "anonymous_group_for_meeting_id":
		err := json.Unmarshal(content, &m.anonymousGroupForMeeting)
		if err != nil {
			return err
		}
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "meeting_mediafile_inherited_access_group_ids":
		err := json.Unmarshal(content, &m.meetingMediafileInheritedAccessGroups)
		if err != nil {
			return err
		}
	case "poll_ids":
		err := json.Unmarshal(content, &m.polls)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return nil
}

func (m *Group) Get(field string) interface{} {
	switch field {
	case "admin_group_for_meeting_id":
		return m.AdminGroupForMeetingID
	case "anonymous_group_for_meeting_id":
		return m.AnonymousGroupForMeetingID
	case "default_group_for_meeting_id":
		return m.DefaultGroupForMeetingID
	case "external_id":
		return m.ExternalID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_mediafile_access_group_ids":
		return m.MeetingMediafileAccessGroupIDs
	case "meeting_mediafile_inherited_access_group_ids":
		return m.MeetingMediafileInheritedAccessGroupIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "permissions":
		return m.Permissions
	case "poll_ids":
		return m.PollIDs
	case "read_chat_group_ids":
		return m.ReadChatGroupIDs
	case "read_comment_section_ids":
		return m.ReadCommentSectionIDs
	case "used_as_assignment_poll_default_id":
		return m.UsedAsAssignmentPollDefaultID
	case "used_as_motion_poll_default_id":
		return m.UsedAsMotionPollDefaultID
	case "used_as_poll_default_id":
		return m.UsedAsPollDefaultID
	case "used_as_topic_poll_default_id":
		return m.UsedAsTopicPollDefaultID
	case "weight":
		return m.Weight
	case "write_chat_group_ids":
		return m.WriteChatGroupIDs
	case "write_comment_section_ids":
		return m.WriteCommentSectionIDs
	}

	return nil
}

func (m *Group) GetFqids(field string) []string {
	switch field {
	case "write_chat_group_ids":
		r := make([]string, len(m.WriteChatGroupIDs))
		for i, id := range m.WriteChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "write_comment_section_ids":
		r := make([]string, len(m.WriteCommentSectionIDs))
		for i, id := range m.WriteCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r

	case "admin_group_for_meeting_id":
		if m.AdminGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.AdminGroupForMeetingID)}
		}

	case "used_as_topic_poll_default_id":
		if m.UsedAsTopicPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsTopicPollDefaultID)}
		}

	case "default_group_for_meeting_id":
		if m.DefaultGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultGroupForMeetingID)}
		}

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "used_as_assignment_poll_default_id":
		if m.UsedAsAssignmentPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsAssignmentPollDefaultID)}
		}

	case "meeting_mediafile_access_group_ids":
		r := make([]string, len(m.MeetingMediafileAccessGroupIDs))
		for i, id := range m.MeetingMediafileAccessGroupIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "used_as_poll_default_id":
		if m.UsedAsPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsPollDefaultID)}
		}

	case "read_chat_group_ids":
		r := make([]string, len(m.ReadChatGroupIDs))
		for i, id := range m.ReadChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "read_comment_section_ids":
		r := make([]string, len(m.ReadCommentSectionIDs))
		for i, id := range m.ReadCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r

	case "used_as_motion_poll_default_id":
		if m.UsedAsMotionPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsMotionPollDefaultID)}
		}

	case "anonymous_group_for_meeting_id":
		if m.AnonymousGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.AnonymousGroupForMeetingID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_mediafile_inherited_access_group_ids":
		r := make([]string, len(m.MeetingMediafileInheritedAccessGroupIDs))
		for i, id := range m.MeetingMediafileInheritedAccessGroupIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "poll_ids":
		r := make([]string, len(m.PollIDs))
		for i, id := range m.PollIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Group) Update(data map[string]string) error {
	if val, ok := data["admin_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdminGroupForMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["anonymous_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AnonymousGroupForMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultGroupForMeetingID)
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

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
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

	if val, ok := data["meeting_mediafile_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileAccessGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_mediafile_inherited_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileInheritedAccessGroupIDs)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["permissions"]; ok {
		err := json.Unmarshal([]byte(val), &m.Permissions)
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

	if val, ok := data["read_chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadChatGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["read_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadCommentSectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_assignment_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsAssignmentPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_motion_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsMotionPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_topic_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsTopicPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["write_chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteChatGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["write_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteCommentSectionIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
