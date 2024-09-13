package models

type Assignment struct {
	TagIDs                        []int
	ProjectionIDs                 []int
	SequentialNumber              int
	Description                   *string
	ID                            *int
	MeetingID                     int
	CandidateIDs                  []int
	DefaultPollDescription        *string
	OpenPosts                     *int
	Phase                         *string
	PollIDs                       []int
	AttachmentMeetingMediafileIDs []int
	ListOfSpeakersID              int
	Title                         string
	AgendaItemID                  *int
	NumberPollCandidates          *bool
}

func (m Assignment) CollectionName() string {
	return "assignment"
}
