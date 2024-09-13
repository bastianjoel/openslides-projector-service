package models

type Group struct {
	AdminGroupForMeetingID                  *int
	AnonymousGroupForMeetingID              *int
	DefaultGroupForMeetingID                *int
	ExternalID                              *string
	ID                                      *int
	MeetingID                               int
	MeetingMediafileAccessGroupIDs          []int
	MeetingMediafileInheritedAccessGroupIDs []int
	MeetingUserIDs                          []int
	Name                                    string
	Permissions                             []string
	PollIDs                                 []int
	ReadChatGroupIDs                        []int
	ReadCommentSectionIDs                   []int
	UsedAsAssignmentPollDefaultID           *int
	UsedAsMotionPollDefaultID               *int
	UsedAsPollDefaultID                     *int
	UsedAsTopicPollDefaultID                *int
	Weight                                  *int
	WriteChatGroupIDs                       []int
	WriteCommentSectionIDs                  []int
}

func (m Group) CollectionName() string {
	return "group"
}
