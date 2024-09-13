package models

type Projector struct {
	AspectRatioDenominator                                    *int
	AspectRatioNumerator                                      *int
	BackgroundColor                                           *string
	ChyronBackgroundColor                                     *string
	ChyronBackgroundColor2                                    *string
	ChyronFontColor                                           *string
	ChyronFontColor2                                          *string
	Color                                                     *string
	CurrentProjectionIDs                                      []int
	HeaderBackgroundColor                                     *string
	HeaderFontColor                                           *string
	HeaderH1Color                                             *string
	HistoryProjectionIDs                                      []int
	ID                                                        *int
	IsInternal                                                *bool
	MeetingID                                                 int
	Name                                                      *string
	PreviewProjectionIDs                                      []int
	Scale                                                     *int
	Scroll                                                    *int
	SequentialNumber                                          int
	ShowClock                                                 *bool
	ShowHeaderFooter                                          *bool
	ShowLogo                                                  *bool
	ShowTitle                                                 *bool
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        *int
	UsedAsDefaultProjectorForAmendmentInMeetingID             *int
	UsedAsDefaultProjectorForAssignmentInMeetingID            *int
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        *int
	UsedAsDefaultProjectorForCountdownInMeetingID             *int
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID *int
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        *int
	UsedAsDefaultProjectorForMediafileInMeetingID             *int
	UsedAsDefaultProjectorForMessageInMeetingID               *int
	UsedAsDefaultProjectorForMotionBlockInMeetingID           *int
	UsedAsDefaultProjectorForMotionInMeetingID                *int
	UsedAsDefaultProjectorForMotionPollInMeetingID            *int
	UsedAsDefaultProjectorForPollInMeetingID                  *int
	UsedAsDefaultProjectorForTopicInMeetingID                 *int
	UsedAsReferenceProjectorMeetingID                         *int
	Width                                                     *int
}

func (m Projector) CollectionName() string {
	return "projector"
}
