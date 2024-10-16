package models

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
