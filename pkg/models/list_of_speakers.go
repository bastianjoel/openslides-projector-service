package models

type ListOfSpeakers struct {
	Closed                          *bool  `json:"closed"`
	ContentObjectID                 string `json:"content_object_id"`
	ID                              int    `json:"id"`
	MeetingID                       int    `json:"meeting_id"`
	ProjectionIDs                   []int  `json:"projection_ids"`
	SequentialNumber                int    `json:"sequential_number"`
	SpeakerIDs                      []int  `json:"speaker_ids"`
	StructureLevelListOfSpeakersIDs []int  `json:"structure_level_list_of_speakers_ids"`
}

func (m ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}

func (m ListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "closed":
		return m.Closed
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}
