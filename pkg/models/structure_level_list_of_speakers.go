package models

type StructureLevelListOfSpeakers struct {
	CurrentStartTime *int
	ID               int
	ListOfSpeakersID int
	MeetingID        int
	RemainingTime    float32
	AdditionalTime   *float32
	InitialTime      int
	SpeakerIDs       []int
	StructureLevelID int
}

func (m StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
}
