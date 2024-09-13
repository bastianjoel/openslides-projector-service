package models

type Topic struct {
	ProjectionIDs                 []int
	SequentialNumber              int
	Text                          *string
	Title                         string
	AttachmentMeetingMediafileIDs []int
	ID                            *int
	ListOfSpeakersID              int
	AgendaItemID                  int
	MeetingID                     int
	PollIDs                       []int
}

func (m Topic) CollectionName() string {
	return "topic"
}
