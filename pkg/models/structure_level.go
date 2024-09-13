package models

type StructureLevel struct {
	ID                              int
	MeetingID                       int
	MeetingUserIDs                  []int
	Name                            string
	StructureLevelListOfSpeakersIDs []int
	Color                           *string
	DefaultTime                     *int
}

func (m StructureLevel) CollectionName() string {
	return "structure_level"
}
