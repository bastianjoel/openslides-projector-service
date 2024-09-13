package models

type MeetingUser struct {
	VoteWeight                   *string
	AboutMe                      *string
	MotionEditorIDs              []int
	MotionSubmitterIDs           []int
	Number                       *string
	PersonalNoteIDs              []int
	VoteDelegationsFromIDs       []int
	GroupIDs                     []int
	LockedOut                    *bool
	MotionWorkingGroupSpeakerIDs []int
	SupportedMotionIDs           []int
	VoteDelegatedToID            *int
	AssignmentCandidateIDs       []int
	ChatMessageIDs               []int
	Comment                      *string
	SpeakerIDs                   []int
	StructureLevelIDs            []int
	ID                           int
	MeetingID                    int
	UserID                       int
}

func (m MeetingUser) CollectionName() string {
	return "meeting_user"
}
