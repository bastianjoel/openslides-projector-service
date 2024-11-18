package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID        *int   `json:"default_amendment_workflow_meeting_id"`
	DefaultStatuteAmendmentWorkflowMeetingID *int   `json:"default_statute_amendment_workflow_meeting_id"`
	DefaultWorkflowMeetingID                 *int   `json:"default_workflow_meeting_id"`
	FirstStateID                             int    `json:"first_state_id"`
	ID                                       int    `json:"id"`
	MeetingID                                int    `json:"meeting_id"`
	Name                                     string `json:"name"`
	SequentialNumber                         int    `json:"sequential_number"`
	StateIDs                                 []int  `json:"state_ids"`
	loadedRelations                          map[string]struct{}
	defaultAmendmentWorkflowMeeting          *Meeting
	defaultWorkflowMeeting                   *Meeting
	firstState                               *MotionState
	meeting                                  *Meeting
	states                                   []MotionState
	defaultStatuteAmendmentWorkflowMeeting   *Meeting
}

func (m *MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}

func (m *MotionWorkflow) DefaultAmendmentWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_amendment_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultAmendmentWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultAmendmentWorkflowMeeting
}

func (m *MotionWorkflow) DefaultWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultWorkflowMeeting
}

func (m *MotionWorkflow) FirstState() MotionState {
	if _, ok := m.loadedRelations["first_state_id"]; !ok {
		log.Panic().Msg("Tried to access FirstState relation of MotionWorkflow which was not loaded.")
	}

	return *m.firstState
}

func (m *MotionWorkflow) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkflow which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionWorkflow) States() []MotionState {
	if _, ok := m.loadedRelations["state_ids"]; !ok {
		log.Panic().Msg("Tried to access States relation of MotionWorkflow which was not loaded.")
	}

	return m.states
}

func (m *MotionWorkflow) DefaultStatuteAmendmentWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_statute_amendment_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultStatuteAmendmentWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultStatuteAmendmentWorkflowMeeting
}

func (m *MotionWorkflow) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "default_amendment_workflow_meeting_id":
			m.defaultAmendmentWorkflowMeeting = content.(*Meeting)
		case "default_workflow_meeting_id":
			m.defaultWorkflowMeeting = content.(*Meeting)
		case "first_state_id":
			m.firstState = content.(*MotionState)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "state_ids":
			m.states = content.([]MotionState)
		case "default_statute_amendment_workflow_meeting_id":
			m.defaultStatuteAmendmentWorkflowMeeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionWorkflow) SetRelatedJSON(field string, content []byte) error {
	switch field {
	case "default_amendment_workflow_meeting_id":
		err := json.Unmarshal(content, &m.defaultAmendmentWorkflowMeeting)
		if err != nil {
			return err
		}
	case "default_workflow_meeting_id":
		err := json.Unmarshal(content, &m.defaultWorkflowMeeting)
		if err != nil {
			return err
		}
	case "first_state_id":
		err := json.Unmarshal(content, &m.firstState)
		if err != nil {
			return err
		}
	case "meeting_id":
		err := json.Unmarshal(content, &m.meeting)
		if err != nil {
			return err
		}
	case "state_ids":
		err := json.Unmarshal(content, &m.states)
		if err != nil {
			return err
		}
	case "default_statute_amendment_workflow_meeting_id":
		err := json.Unmarshal(content, &m.defaultStatuteAmendmentWorkflowMeeting)
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

func (m *MotionWorkflow) Get(field string) interface{} {
	switch field {
	case "default_amendment_workflow_meeting_id":
		return m.DefaultAmendmentWorkflowMeetingID
	case "default_statute_amendment_workflow_meeting_id":
		return m.DefaultStatuteAmendmentWorkflowMeetingID
	case "default_workflow_meeting_id":
		return m.DefaultWorkflowMeetingID
	case "first_state_id":
		return m.FirstStateID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "sequential_number":
		return m.SequentialNumber
	case "state_ids":
		return m.StateIDs
	}

	return nil
}

func (m *MotionWorkflow) GetFqids(field string) []string {
	switch field {
	case "default_amendment_workflow_meeting_id":
		if m.DefaultAmendmentWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultAmendmentWorkflowMeetingID)}
		}

	case "default_workflow_meeting_id":
		if m.DefaultWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultWorkflowMeetingID)}
		}

	case "first_state_id":
		return []string{"motion_state/" + strconv.Itoa(m.FirstStateID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "state_ids":
		r := make([]string, len(m.StateIDs))
		for i, id := range m.StateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "default_statute_amendment_workflow_meeting_id":
		if m.DefaultStatuteAmendmentWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultStatuteAmendmentWorkflowMeetingID)}
		}
	}
	return []string{}
}

func (m *MotionWorkflow) Update(data map[string]string) error {
	if val, ok := data["default_amendment_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultAmendmentWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_statute_amendment_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultStatuteAmendmentWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstStateID)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
