package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type PollCandidateList struct {
	ID               int   `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
	loadedRelations  map[string]struct{}
	option           *Option
	pollCandidates   *PollCandidate
	meeting          *Meeting
}

func (m PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}

func (m *PollCandidateList) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of PollCandidateList which was not loaded.")
	}

	return *m.option
}

func (m *PollCandidateList) PollCandidates() *PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of PollCandidateList which was not loaded.")
	}

	return m.pollCandidates
}

func (m *PollCandidateList) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidateList which was not loaded.")
	}

	return *m.meeting
}

func (m PollCandidateList) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	}

	return nil
}

func (m PollCandidateList) Update(data map[string]string) error {
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

	if val, ok := data["option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
