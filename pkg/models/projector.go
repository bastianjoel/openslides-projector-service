package models

type Projector struct {
	UsedAsDefaultProjectorForTopicInMeetingID                 *int
	Width                                                     *int
	CurrentProjectionIDs                                      []int
	ID                                                        *int
	SequentialNumber                                          int
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        *int
	UsedAsDefaultProjectorForCountdownInMeetingID             *int
	UsedAsDefaultProjectorForMessageInMeetingID               *int
	HistoryProjectionIDs                                      []int
	MeetingID                                                 int
	Name                                                      *string
	ShowClock                                                 *bool
	ShowLogo                                                  *bool
	UsedAsDefaultProjectorForMotionInMeetingID                *int
	AspectRatioDenominator                                    *int
	BackgroundColor                                           *string
	ChyronFontColor                                           *string
	Scroll                                                    *int
	UsedAsDefaultProjectorForMotionPollInMeetingID            *int
	HeaderH1Color                                             *string
	AspectRatioNumerator                                      *int
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        *int
	HeaderBackgroundColor                                     *string
	IsInternal                                                *bool
	PreviewProjectionIDs                                      []int
	ShowTitle                                                 *bool
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID *int
	UsedAsDefaultProjectorForPollInMeetingID                  *int
	UsedAsDefaultProjectorForAmendmentInMeetingID             *int
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        *int
	ChyronBackgroundColor                                     *string
	ChyronBackgroundColor2                                    *string
	ChyronFontColor2                                          *string
	Color                                                     *string
	Scale                                                     *int
	ShowHeaderFooter                                          *bool
	HeaderFontColor                                           *string
	UsedAsDefaultProjectorForAssignmentInMeetingID            *int
	UsedAsDefaultProjectorForMediafileInMeetingID             *int
	UsedAsDefaultProjectorForMotionBlockInMeetingID           *int
	UsedAsReferenceProjectorMeetingID                         *int
}

func (m Projector) CollectionName() string {
	return "projector"
}
