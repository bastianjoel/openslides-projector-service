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

func (m StructureLevel) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "default_time":
		return m.DefaultTime
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}
