package models

import (
	"encoding/json"

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
	anonymousGroupForMeeting                *Meeting
	meeting                                 *Meeting
	meetingMediafileInheritedAccessGroups   *MeetingMediafile
	meetingUsers                            *MeetingUser
	readChatGroups                          *ChatGroup
	usedAsPollDefault                       *Meeting
	adminGroupForMeeting                    *Meeting
	defaultGroupForMeeting                  *Meeting
	meetingMediafileAccessGroups            *MeetingMediafile
	readCommentSections                     *MotionCommentSection
	polls                                   *Poll
	usedAsAssignmentPollDefault             *Meeting
	usedAsMotionPollDefault                 *Meeting
	usedAsTopicPollDefault                  *Meeting
	writeChatGroups                         *ChatGroup
	writeCommentSections                    *MotionCommentSection
}

func (m *Group) CollectionName() string {
	return "group"
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

func (m *Group) MeetingMediafileInheritedAccessGroups() *MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_inherited_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileInheritedAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileInheritedAccessGroups
}

func (m *Group) MeetingUsers() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of Group which was not loaded.")
	}

	return m.meetingUsers
}

func (m *Group) ReadChatGroups() *ChatGroup {
	if _, ok := m.loadedRelations["read_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadChatGroups relation of Group which was not loaded.")
	}

	return m.readChatGroups
}

func (m *Group) UsedAsPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsPollDefault
}

func (m *Group) AdminGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["admin_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access AdminGroupForMeeting relation of Group which was not loaded.")
	}

	return m.adminGroupForMeeting
}

func (m *Group) DefaultGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultGroupForMeeting relation of Group which was not loaded.")
	}

	return m.defaultGroupForMeeting
}

func (m *Group) MeetingMediafileAccessGroups() *MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileAccessGroups
}

func (m *Group) ReadCommentSections() *MotionCommentSection {
	if _, ok := m.loadedRelations["read_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadCommentSections relation of Group which was not loaded.")
	}

	return m.readCommentSections
}

func (m *Group) Polls() *Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Group which was not loaded.")
	}

	return m.polls
}

func (m *Group) UsedAsAssignmentPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_assignment_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsAssignmentPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsAssignmentPollDefault
}

func (m *Group) UsedAsMotionPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_motion_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsMotionPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsMotionPollDefault
}

func (m *Group) UsedAsTopicPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_topic_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsTopicPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsTopicPollDefault
}

func (m *Group) WriteChatGroups() *ChatGroup {
	if _, ok := m.loadedRelations["write_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteChatGroups relation of Group which was not loaded.")
	}

	return m.writeChatGroups
}

func (m *Group) WriteCommentSections() *MotionCommentSection {
	if _, ok := m.loadedRelations["write_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteCommentSections relation of Group which was not loaded.")
	}

	return m.writeCommentSections
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
