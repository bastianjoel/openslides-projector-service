package models

type StructureLevelListOfSpeakers struct {
	AdditionalTime   *float32 `json:"additional_time"`
	CurrentStartTime *int     `json:"current_start_time"`
	ID               int      `json:"id"`
	InitialTime      int      `json:"initial_time"`
	ListOfSpeakersID int      `json:"list_of_speakers_id"`
	MeetingID        int      `json:"meeting_id"`
	RemainingTime    float32  `json:"remaining_time"`
	SpeakerIDs       []int    `json:"speaker_ids"`
	StructureLevelID int      `json:"structure_level_id"`
}

func (m StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
}

func (m StructureLevelListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "additional_time":
		return m.AdditionalTime
	case "current_start_time":
		return m.CurrentStartTime
	case "id":
		return m.ID
	case "initial_time":
		return m.InitialTime
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "remaining_time":
		return m.RemainingTime
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_id":
		return m.StructureLevelID
	}

	return nil
}
