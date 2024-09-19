package models

type ListOfSpeakers struct {
	Closed                          *bool  `json:"closed"`
	ContentObjectID                 string `json:"content_object_id"`
	ID                              *int   `json:"id"`
	MeetingID                       int    `json:"meeting_id"`
	ProjectionIDs                   []int  `json:"projection_ids"`
	SequentialNumber                int    `json:"sequential_number"`
	SpeakerIDs                      []int  `json:"speaker_ids"`
	StructureLevelListOfSpeakersIDs []int  `json:"structure_level_list_of_speakers_ids"`
}

func (m ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}
