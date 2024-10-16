package models

import "encoding/json"

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
}

func (m Group) CollectionName() string {
	return "group"
}

func (m Group) Get(field string) interface{} {
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

func (m Group) Update(data map[string]string) error {
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
