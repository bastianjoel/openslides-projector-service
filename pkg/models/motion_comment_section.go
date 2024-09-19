package models

type MotionCommentSection struct {
	CommentIDs        []int  `json:"comment_ids"`
	ID                *int   `json:"id"`
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
