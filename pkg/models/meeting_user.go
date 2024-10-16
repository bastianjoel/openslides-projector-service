package models

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
}

func (m MeetingUser) CollectionName() string {
	return "meeting_user"
}

func (m MeetingUser) Get(field string) interface{} {
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
