package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MeetingUser struct {
	AboutMe                      *string `json:"about_me"`
	AssignmentCandidateIDs       []int   `json:"assignment_candidate_ids"`
	ChatMessageIDs               []int   `json:"chat_message_ids"`
	Comment                      *string `json:"comment"`
	GroupIDs                     []int   `json:"group_ids"`
	ID                           int     `json:"id"`
	LockedOut                    *bool   `json:"locked_out"`
	MeetingID                    int     `json:"meeting_id"`
	MotionEditorIDs              []int   `json:"motion_editor_ids"`
	MotionSubmitterIDs           []int   `json:"motion_submitter_ids"`
	MotionWorkingGroupSpeakerIDs []int   `json:"motion_working_group_speaker_ids"`
	Number                       *string `json:"number"`
	PersonalNoteIDs              []int   `json:"personal_note_ids"`
	SpeakerIDs                   []int   `json:"speaker_ids"`
	StructureLevelIDs            []int   `json:"structure_level_ids"`
	SupportedMotionIDs           []int   `json:"supported_motion_ids"`
	UserID                       int     `json:"user_id"`
	VoteDelegatedToID            *int    `json:"vote_delegated_to_id"`
	VoteDelegationsFromIDs       []int   `json:"vote_delegations_from_ids"`
	VoteWeight                   *string `json:"vote_weight"`
	loadedRelations              map[string]struct{}
	voteDelegatedTo              *MeetingUser
	groups                       *Group
	motionWorkingGroupSpeakers   *MotionWorkingGroupSpeaker
	supportedMotions             *Motion
	user                         *User
	voteDelegationsFroms         *MeetingUser
	chatMessages                 *ChatMessage
	personalNotes                *PersonalNote
	speakers                     *Speaker
	structureLevels              *StructureLevel
	assignmentCandidates         *AssignmentCandidate
	meeting                      *Meeting
	motionSubmitters             *MotionSubmitter
	motionEditors                *MotionEditor
}

func (m *MeetingUser) CollectionName() string {
	return "meeting_user"
}

func (m *MeetingUser) VoteDelegatedTo() *MeetingUser {
	if _, ok := m.loadedRelations["vote_delegated_to_id"]; !ok {
		log.Panic().Msg("Tried to access VoteDelegatedTo relation of MeetingUser which was not loaded.")
	}

	return m.voteDelegatedTo
}

func (m *MeetingUser) Groups() *Group {
	if _, ok := m.loadedRelations["group_ids"]; !ok {
		log.Panic().Msg("Tried to access Groups relation of MeetingUser which was not loaded.")
	}

	return m.groups
}

func (m *MeetingUser) MotionWorkingGroupSpeakers() *MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkingGroupSpeakers relation of MeetingUser which was not loaded.")
	}

	return m.motionWorkingGroupSpeakers
}

func (m *MeetingUser) SupportedMotions() *Motion {
	if _, ok := m.loadedRelations["supported_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access SupportedMotions relation of MeetingUser which was not loaded.")
	}

	return m.supportedMotions
}

func (m *MeetingUser) User() User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of MeetingUser which was not loaded.")
	}

	return *m.user
}

func (m *MeetingUser) VoteDelegationsFroms() *MeetingUser {
	if _, ok := m.loadedRelations["vote_delegations_from_ids"]; !ok {
		log.Panic().Msg("Tried to access VoteDelegationsFroms relation of MeetingUser which was not loaded.")
	}

	return m.voteDelegationsFroms
}

func (m *MeetingUser) ChatMessages() *ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of MeetingUser which was not loaded.")
	}

	return m.chatMessages
}

func (m *MeetingUser) PersonalNotes() *PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of MeetingUser which was not loaded.")
	}

	return m.personalNotes
}

func (m *MeetingUser) Speakers() *Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of MeetingUser which was not loaded.")
	}

	return m.speakers
}

func (m *MeetingUser) StructureLevels() *StructureLevel {
	if _, ok := m.loadedRelations["structure_level_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevels relation of MeetingUser which was not loaded.")
	}

	return m.structureLevels
}

func (m *MeetingUser) AssignmentCandidates() *AssignmentCandidate {
	if _, ok := m.loadedRelations["assignment_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentCandidates relation of MeetingUser which was not loaded.")
	}

	return m.assignmentCandidates
}

func (m *MeetingUser) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MeetingUser which was not loaded.")
	}

	return *m.meeting
}

func (m *MeetingUser) MotionSubmitters() *MotionSubmitter {
	if _, ok := m.loadedRelations["motion_submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionSubmitters relation of MeetingUser which was not loaded.")
	}

	return m.motionSubmitters
}

func (m *MeetingUser) MotionEditors() *MotionEditor {
	if _, ok := m.loadedRelations["motion_editor_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionEditors relation of MeetingUser which was not loaded.")
	}

	return m.motionEditors
}

func (m *MeetingUser) Get(field string) interface{} {
	switch field {
	case "about_me":
		return m.AboutMe
	case "assignment_candidate_ids":
		return m.AssignmentCandidateIDs
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "comment":
		return m.Comment
	case "group_ids":
		return m.GroupIDs
	case "id":
		return m.ID
	case "locked_out":
		return m.LockedOut
	case "meeting_id":
		return m.MeetingID
	case "motion_editor_ids":
		return m.MotionEditorIDs
	case "motion_submitter_ids":
		return m.MotionSubmitterIDs
	case "motion_working_group_speaker_ids":
		return m.MotionWorkingGroupSpeakerIDs
	case "number":
		return m.Number
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_ids":
		return m.StructureLevelIDs
	case "supported_motion_ids":
		return m.SupportedMotionIDs
	case "user_id":
		return m.UserID
	case "vote_delegated_to_id":
		return m.VoteDelegatedToID
	case "vote_delegations_from_ids":
		return m.VoteDelegationsFromIDs
	case "vote_weight":
		return m.VoteWeight
	}

	return nil
}

func (m *MeetingUser) Update(data map[string]string) error {
	if val, ok := data["about_me"]; ok {
		err := json.Unmarshal([]byte(val), &m.AboutMe)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentCandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GroupIDs)
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

	if val, ok := data["locked_out"]; ok {
		err := json.Unmarshal([]byte(val), &m.LockedOut)
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

	if val, ok := data["motion_editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionEditorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionSubmitterIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["number"]; ok {
		err := json.Unmarshal([]byte(val), &m.Number)
		if err != nil {
			return err
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
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

	if val, ok := data["structure_level_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["supported_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SupportedMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_delegated_to_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteDelegatedToID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_delegations_from_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteDelegationsFromIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteWeight)
		if err != nil {
			return err
		}
	}

	return nil
}
