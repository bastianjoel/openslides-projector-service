package models

type Group struct {
	MeetingMediafileAccessGroupIDs          []int
	ID                                      *int
	ExternalID                              *string
	MeetingUserIDs                          []int
	ReadCommentSectionIDs                   []int
	UsedAsMotionPollDefaultID               *int
	UsedAsPollDefaultID                     *int
	WriteChatGroupIDs                       []int
	AdminGroupForMeetingID                  *int
	DefaultGroupForMeetingID                *int
	MeetingID                               int
	Name                                    string
	PollIDs                                 []int
	ReadChatGroupIDs                        []int
	Weight                                  *int
	WriteCommentSectionIDs                  []int
	AnonymousGroupForMeetingID              *int
	Permissions                             []string
	UsedAsAssignmentPollDefaultID           *int
	UsedAsTopicPollDefaultID                *int
	MeetingMediafileInheritedAccessGroupIDs []int
}

func (m Group) CollectionName() string {
	return "group"
}
