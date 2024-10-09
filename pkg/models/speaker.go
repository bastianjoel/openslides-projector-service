package models

type Speaker struct {
	BeginTime                      *int    `json:"begin_time"`
	EndTime                        *int    `json:"end_time"`
	ID                             int     `json:"id"`
	ListOfSpeakersID               int     `json:"list_of_speakers_id"`
	MeetingID                      int     `json:"meeting_id"`
	MeetingUserID                  *int    `json:"meeting_user_id"`
	Note                           *string `json:"note"`
	PauseTime                      *int    `json:"pause_time"`
	PointOfOrder                   *bool   `json:"point_of_order"`
	PointOfOrderCategoryID         *int    `json:"point_of_order_category_id"`
	SpeechState                    *string `json:"speech_state"`
	StructureLevelListOfSpeakersID *int    `json:"structure_level_list_of_speakers_id"`
	TotalPause                     *int    `json:"total_pause"`
	UnpauseTime                    *int    `json:"unpause_time"`
	Weight                         *int    `json:"weight"`
}

func (m Speaker) CollectionName() string {
	return "speaker"
}
