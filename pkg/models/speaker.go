package models

type Speaker struct {
	ID                             *int
	MeetingUserID                  *int
	PauseTime                      *int
	PointOfOrderCategoryID         *int
	UnpauseTime                    *int
	BeginTime                      *int
	MeetingID                      int
	SpeechState                    *string
	Weight                         *int
	EndTime                        *int
	ListOfSpeakersID               int
	Note                           *string
	StructureLevelListOfSpeakersID *int
	TotalPause                     *int
	PointOfOrder                   *bool
}

func (m Speaker) CollectionName() string {
	return "speaker"
}
