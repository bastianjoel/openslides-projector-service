package models

type Assignment struct {
	AgendaItemID                  *int    `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	CandidateIDs                  []int   `json:"candidate_ids"`
	DefaultPollDescription        *string `json:"default_poll_description"`
	Description                   *string `json:"description"`
	ID                            int     `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	NumberPollCandidates          *bool   `json:"number_poll_candidates"`
	OpenPosts                     *int    `json:"open_posts"`
	Phase                         *string `json:"phase"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	TagIDs                        []int   `json:"tag_ids"`
	Title                         string  `json:"title"`
}

func (m Assignment) CollectionName() string {
	return "assignment"
}

func (m Assignment) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "candidate_ids":
		return m.CandidateIDs
	case "default_poll_description":
		return m.DefaultPollDescription
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "number_poll_candidates":
		return m.NumberPollCandidates
	case "open_posts":
		return m.OpenPosts
	case "phase":
		return m.Phase
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "tag_ids":
		return m.TagIDs
	case "title":
		return m.Title
	}

	return nil
}
