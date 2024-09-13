package models

type MeetingUser struct {
	AboutMe                      *string
	AssignmentCandidateIDs       []int
	ChatMessageIDs               []int
	Comment                      *string
	GroupIDs                     []int
	ID                           int
	LockedOut                    *bool
	MeetingID                    int
	MotionEditorIDs              []int
	MotionSubmitterIDs           []int
	MotionWorkingGroupSpeakerIDs []int
	Number                       *string
	PersonalNoteIDs              []int
	SpeakerIDs                   []int
	StructureLevelIDs            []int
	SupportedMotionIDs           []int
	UserID                       int
	VoteDelegatedToID            *int
	VoteDelegationsFromIDs       []int
	VoteWeight                   *string
}

func (m MeetingUser) CollectionName() string {
	return "meeting_user"
}
