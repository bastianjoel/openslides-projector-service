package models

type MeetingMediafile struct {
	ProjectionIDs                          []int
	UsedAsFontMonospaceInMeetingID         *int
	UsedAsFontRegularInMeetingID           *int
	UsedAsLogoPdfBallotPaperInMeetingID    *int
	UsedAsLogoPdfHeaderLInMeetingID        *int
	UsedAsLogoProjectorMainInMeetingID     *int
	IsPublic                               bool
	MeetingID                              int
	InheritedAccessGroupIDs                []int
	UsedAsFontProjectorH1InMeetingID       *int
	UsedAsLogoPdfFooterRInMeetingID        *int
	UsedAsLogoProjectorHeaderInMeetingID   *int
	AttachmentIDs                          []string
	ID                                     *int
	UsedAsFontChyronSpeakerNameInMeetingID *int
	UsedAsFontItalicInMeetingID            *int
	UsedAsFontProjectorH2InMeetingID       *int
	UsedAsLogoPdfFooterLInMeetingID        *int
	UsedAsLogoWebHeaderInMeetingID         *int
	ListOfSpeakersID                       *int
	MediafileID                            int
	UsedAsFontBoldItalicInMeetingID        *int
	UsedAsLogoPdfHeaderRInMeetingID        *int
	AccessGroupIDs                         []int
	UsedAsFontBoldInMeetingID              *int
}

func (m MeetingMediafile) CollectionName() string {
	return "meeting_mediafile"
}
