package models

type ListOfSpeakers struct {
	Closed                          *bool
	ContentObjectID                 string
	ID                              *int
	MeetingID                       int
	ProjectionIDs                   []int
	SequentialNumber                int
	SpeakerIDs                      []int
	StructureLevelListOfSpeakersIDs []int
}

func (m ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}
