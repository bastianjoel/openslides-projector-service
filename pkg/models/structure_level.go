package models

type StructureLevel struct {
	Color                           *string `json:"color"`
	DefaultTime                     *int    `json:"default_time"`
	ID                              int     `json:"id"`
	MeetingID                       int     `json:"meeting_id"`
	MeetingUserIDs                  []int   `json:"meeting_user_ids"`
	Name                            string  `json:"name"`
	StructureLevelListOfSpeakersIDs []int   `json:"structure_level_list_of_speakers_ids"`
}

func (m StructureLevel) CollectionName() string {
	return "structure_level"
}
