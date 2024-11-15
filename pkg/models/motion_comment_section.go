package models

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MotionCommentSection struct {
	CommentIDs        []int  `json:"comment_ids"`
	ID                int    `json:"id"`
	MeetingID         int    `json:"meeting_id"`
	Name              string `json:"name"`
	ReadGroupIDs      []int  `json:"read_group_ids"`
	SequentialNumber  int    `json:"sequential_number"`
	SubmitterCanWrite *bool  `json:"submitter_can_write"`
	Weight            *int   `json:"weight"`
	WriteGroupIDs     []int  `json:"write_group_ids"`
	loadedRelations   map[string]struct{}
	readGroups        *Group
	writeGroups       *Group
	comments          *MotionComment
	meeting           *Meeting
}

func (m MotionCommentSection) CollectionName() string {
	return "motion_comment_section"
}

func (m *MotionCommentSection) ReadGroups() *Group {
	if _, ok := m.loadedRelations["read_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadGroups relation of MotionCommentSection which was not loaded.")
	}

	return m.readGroups
}

func (m *MotionCommentSection) WriteGroups() *Group {
	if _, ok := m.loadedRelations["write_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteGroups relation of MotionCommentSection which was not loaded.")
	}

	return m.writeGroups
}

func (m *MotionCommentSection) Comments() *MotionComment {
	if _, ok := m.loadedRelations["comment_ids"]; !ok {
		log.Panic().Msg("Tried to access Comments relation of MotionCommentSection which was not loaded.")
	}

	return m.comments
}

func (m *MotionCommentSection) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionCommentSection which was not loaded.")
	}

	return *m.meeting
}

func (m MotionCommentSection) Get(field string) interface{} {
	switch field {
	case "comment_ids":
		return m.CommentIDs
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "read_group_ids":
		return m.ReadGroupIDs
	case "sequential_number":
		return m.SequentialNumber
	case "submitter_can_write":
		return m.SubmitterCanWrite
	case "weight":
		return m.Weight
	case "write_group_ids":
		return m.WriteGroupIDs
	}

	return nil
}

func (m MotionCommentSection) Update(data map[string]string) error {
	if val, ok := data["comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommentIDs)
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

	if val, ok := data["read_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadGroupIDs)
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

	if val, ok := data["submitter_can_write"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterCanWrite)
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

	if val, ok := data["write_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteGroupIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
