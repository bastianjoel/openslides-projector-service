package models

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
}

func (m MotionCommentSection) CollectionName() string {
	return "motion_comment_section"
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
