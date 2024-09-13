package models

type StructureLevelListOfSpeakers struct {
	AdditionalTime   *float32
	CurrentStartTime *int
	ID               int
	InitialTime      int
	ListOfSpeakersID int
	MeetingID        int
	RemainingTime    float32
	SpeakerIDs       []int
	StructureLevelID int
}

func (m StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
}
