package models

type MotionCommentSection struct {
	CommentIDs        []int
	ID                *int
	MeetingID         int
	Name              string
	ReadGroupIDs      []int
	SequentialNumber  int
	SubmitterCanWrite *bool
	Weight            *int
	WriteGroupIDs     []int
}

func (m MotionCommentSection) CollectionName() string {
	return "motion_comment_section"
}
