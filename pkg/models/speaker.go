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

func (m Speaker) Get(field string) interface{} {
	switch field {
	case "begin_time":
		return m.BeginTime
	case "end_time":
		return m.EndTime
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "note":
		return m.Note
	case "pause_time":
		return m.PauseTime
	case "point_of_order":
		return m.PointOfOrder
	case "point_of_order_category_id":
		return m.PointOfOrderCategoryID
	case "speech_state":
		return m.SpeechState
	case "structure_level_list_of_speakers_id":
		return m.StructureLevelListOfSpeakersID
	case "total_pause":
		return m.TotalPause
	case "unpause_time":
		return m.UnpauseTime
	case "weight":
		return m.Weight
	}

	return nil
}
