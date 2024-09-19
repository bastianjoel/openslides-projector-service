package models

type Topic struct {
	AgendaItemID                  int     `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	ID                            *int    `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	Text                          *string `json:"text"`
	Title                         string  `json:"title"`
}

func (m Topic) CollectionName() string {
	return "topic"
}
