package models

type MotionCommentSection struct {
	MeetingID         int
	Weight            *int
	CommentIDs        []int
	ID                *int
	Name              string
	ReadGroupIDs      []int
	SequentialNumber  int
	SubmitterCanWrite *bool
	WriteGroupIDs     []int
}

func (m MotionCommentSection) CollectionName() string {
	return "motion_comment_section"
}
