package models

type MeetingMediafile struct {
	AccessGroupIDs                         []int
	AttachmentIDs                          []string
	ID                                     *int
	InheritedAccessGroupIDs                []int
	IsPublic                               bool
	ListOfSpeakersID                       *int
	MediafileID                            int
	MeetingID                              int
	ProjectionIDs                          []int
	UsedAsFontBoldInMeetingID              *int
	UsedAsFontBoldItalicInMeetingID        *int
	UsedAsFontChyronSpeakerNameInMeetingID *int
	UsedAsFontItalicInMeetingID            *int
	UsedAsFontMonospaceInMeetingID         *int
	UsedAsFontProjectorH1InMeetingID       *int
	UsedAsFontProjectorH2InMeetingID       *int
	UsedAsFontRegularInMeetingID           *int
	UsedAsLogoPdfBallotPaperInMeetingID    *int
	UsedAsLogoPdfFooterLInMeetingID        *int
	UsedAsLogoPdfFooterRInMeetingID        *int
	UsedAsLogoPdfHeaderLInMeetingID        *int
	UsedAsLogoPdfHeaderRInMeetingID        *int
	UsedAsLogoProjectorHeaderInMeetingID   *int
	UsedAsLogoProjectorMainInMeetingID     *int
	UsedAsLogoWebHeaderInMeetingID         *int
}

func (m MeetingMediafile) CollectionName() string {
	return "meeting_mediafile"
}
