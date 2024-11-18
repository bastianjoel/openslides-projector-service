package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type PollCandidateList struct {
	ID               int   `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	option           *Option
	pollCandidates   []PollCandidate
}

func (m *PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}

func (m *PollCandidateList) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidateList which was not loaded.")
	}

	return *m.meeting
}

func (m *PollCandidateList) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of PollCandidateList which was not loaded.")
	}

	return *m.option
}

func (m *PollCandidateList) PollCandidates() []PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of PollCandidateList which was not loaded.")
	}

	return m.pollCandidates
}

func (m *PollCandidateList) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_id":
			m.option = content.(*Option)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]PollCandidate)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PollCandidateList) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "option_id":
		err := json.Unmarshal(content, &m.option)
		if err != nil {
			return err
		}
	case "poll_candidate_ids":
		err := json.Unmarshal(content, &m.pollCandidates)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return nil
}

func (m *PollCandidateList) Get(field string) interface{} {
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

func (m *PollCandidateList) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_id":
		return []string{"option/" + strconv.Itoa(m.OptionID)}

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *PollCandidateList) Update(data map[string]string) error {
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
