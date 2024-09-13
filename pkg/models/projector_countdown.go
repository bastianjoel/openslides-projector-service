package models

type ProjectorCountdown struct {
	CountdownTime                          *float32
	DefaultTime                            *int
	Description                            *string
	ID                                     *int
	MeetingID                              int
	ProjectionIDs                          []int
	Running                                *bool
	Title                                  string
	UsedAsListOfSpeakersCountdownMeetingID *int
	UsedAsPollCountdownMeetingID           *int
}

func (m ProjectorCountdown) CollectionName() string {
	return "projector_countdown"
}
