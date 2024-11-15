package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type PointOfOrderCategory struct {
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	Rank            int    `json:"rank"`
	SpeakerIDs      []int  `json:"speaker_ids"`
	Text            string `json:"text"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	speakers        *Speaker
}

func (m PointOfOrderCategory) CollectionName() string {
	return "point_of_order_category"
}

func (m *PointOfOrderCategory) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PointOfOrderCategory which was not loaded.")
	}

	return *m.meeting
}

func (m *PointOfOrderCategory) Speakers() *Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of PointOfOrderCategory which was not loaded.")
	}

	return m.speakers
}

func (m PointOfOrderCategory) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "rank":
		return m.Rank
	case "speaker_ids":
		return m.SpeakerIDs
	case "text":
		return m.Text
	}

	return nil
}

func (m PointOfOrderCategory) Update(data map[string]string) error {
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

	if val, ok := data["rank"]; ok {
		err := json.Unmarshal([]byte(val), &m.Rank)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	return nil
}
