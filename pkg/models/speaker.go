package models

import "encoding/json"

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

func (m Speaker) Update(data map[string]string) error {
	if val, ok := data["begin_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.BeginTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["end_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.EndTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["note"]; ok {
		err := json.Unmarshal([]byte(val), &m.Note)
		if err != nil {
			return err
		}
	}

	if val, ok := data["pause_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.PauseTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["point_of_order"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrder)
		if err != nil {
			return err
		}
	}

	if val, ok := data["point_of_order_category_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrderCategoryID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speech_state"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeechState)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["total_pause"]; ok {
		err := json.Unmarshal([]byte(val), &m.TotalPause)
		if err != nil {
			return err
		}
	}

	if val, ok := data["unpause_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.UnpauseTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
