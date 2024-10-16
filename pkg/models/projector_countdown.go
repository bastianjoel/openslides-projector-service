package models

type ProjectorCountdown struct {
	CountdownTime                          *float32 `json:"countdown_time"`
	DefaultTime                            *int     `json:"default_time"`
	Description                            *string  `json:"description"`
	ID                                     int      `json:"id"`
	MeetingID                              int      `json:"meeting_id"`
	ProjectionIDs                          []int    `json:"projection_ids"`
	Running                                *bool    `json:"running"`
	Title                                  string   `json:"title"`
	UsedAsListOfSpeakersCountdownMeetingID *int     `json:"used_as_list_of_speakers_countdown_meeting_id"`
	UsedAsPollCountdownMeetingID           *int     `json:"used_as_poll_countdown_meeting_id"`
}

func (m ProjectorCountdown) CollectionName() string {
	return "projector_countdown"
}

func (m ProjectorCountdown) Get(field string) interface{} {
	switch field {
	case "countdown_time":
		return m.CountdownTime
	case "default_time":
		return m.DefaultTime
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "running":
		return m.Running
	case "title":
		return m.Title
	case "used_as_list_of_speakers_countdown_meeting_id":
		return m.UsedAsListOfSpeakersCountdownMeetingID
	case "used_as_poll_countdown_meeting_id":
		return m.UsedAsPollCountdownMeetingID
	}

	return nil
}
