package models

type Topic struct {
	AgendaItemID                  int
	AttachmentMeetingMediafileIDs []int
	ID                            *int
	ListOfSpeakersID              int
	MeetingID                     int
	PollIDs                       []int
	ProjectionIDs                 []int
	SequentialNumber              int
	Text                          *string
	Title                         string
}

func (m Topic) CollectionName() string {
	return "topic"
}
