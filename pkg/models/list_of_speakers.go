package models

type ListOfSpeakers struct {
	SpeakerIDs                      []int
	StructureLevelListOfSpeakersIDs []int
	Closed                          *bool
	ContentObjectID                 string
	ID                              *int
	MeetingID                       int
	ProjectionIDs                   []int
	SequentialNumber                int
}

func (m ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}
