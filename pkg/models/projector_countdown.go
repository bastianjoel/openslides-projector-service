package models

type ProjectorCountdown struct {
	DefaultTime                            *int
	Description                            *string
	MeetingID                              int
	Running                                *bool
	UsedAsPollCountdownMeetingID           *int
	CountdownTime                          *float32
	ID                                     *int
	ProjectionIDs                          []int
	Title                                  string
	UsedAsListOfSpeakersCountdownMeetingID *int
}

func (m ProjectorCountdown) CollectionName() string {
	return "projector_countdown"
}
