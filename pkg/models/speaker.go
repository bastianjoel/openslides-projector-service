package models

type Speaker struct {
	BeginTime                      *int
	EndTime                        *int
	ID                             *int
	ListOfSpeakersID               int
	MeetingID                      int
	MeetingUserID                  *int
	Note                           *string
	PauseTime                      *int
	PointOfOrder                   *bool
	PointOfOrderCategoryID         *int
	SpeechState                    *string
	StructureLevelListOfSpeakersID *int
	TotalPause                     *int
	UnpauseTime                    *int
	Weight                         *int
}

func (m Speaker) CollectionName() string {
	return "speaker"
}
