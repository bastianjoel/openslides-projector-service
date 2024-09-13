package models

type Assignment struct {
	AgendaItemID                  *int
	AttachmentMeetingMediafileIDs []int
	CandidateIDs                  []int
	DefaultPollDescription        *string
	Description                   *string
	ID                            *int
	ListOfSpeakersID              int
	MeetingID                     int
	NumberPollCandidates          *bool
	OpenPosts                     *int
	Phase                         *string
	PollIDs                       []int
	ProjectionIDs                 []int
	SequentialNumber              int
	TagIDs                        []int
	Title                         string
}

func (m Assignment) CollectionName() string {
	return "assignment"
}
