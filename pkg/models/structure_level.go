package models

import "encoding/json"

type StructureLevel struct {
	Color                           *string `json:"color"`
	DefaultTime                     *int    `json:"default_time"`
	ID                              int     `json:"id"`
	MeetingID                       int     `json:"meeting_id"`
	MeetingUserIDs                  []int   `json:"meeting_user_ids"`
	Name                            string  `json:"name"`
	StructureLevelListOfSpeakersIDs []int   `json:"structure_level_list_of_speakers_ids"`
}

func (m StructureLevel) CollectionName() string {
	return "structure_level"
}

func (m StructureLevel) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "default_time":
		return m.DefaultTime
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}

func (m StructureLevel) Update(data map[string]string) error {
	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultTime)
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

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
