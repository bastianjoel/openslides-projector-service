package models

type StructureLevel struct {
	Color                           *string
	DefaultTime                     *int
	ID                              int
	MeetingID                       int
	MeetingUserIDs                  []int
	Name                            string
	StructureLevelListOfSpeakersIDs []int
}

func (m StructureLevel) CollectionName() string {
	return "structure_level"
}
