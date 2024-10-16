package models

type Option struct {
	Abstain                    *string `json:"abstain"`
	ContentObjectID            *string `json:"content_object_id"`
	ID                         int     `json:"id"`
	MeetingID                  int     `json:"meeting_id"`
	No                         *string `json:"no"`
	PollID                     *int    `json:"poll_id"`
	Text                       *string `json:"text"`
	UsedAsGlobalOptionInPollID *int    `json:"used_as_global_option_in_poll_id"`
	VoteIDs                    []int   `json:"vote_ids"`
	Weight                     *int    `json:"weight"`
	Yes                        *string `json:"yes"`
}

func (m Option) CollectionName() string {
	return "option"
}

func (m Option) Get(field string) interface{} {
	switch field {
	case "abstain":
		return m.Abstain
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "no":
		return m.No
	case "poll_id":
		return m.PollID
	case "text":
		return m.Text
	case "used_as_global_option_in_poll_id":
		return m.UsedAsGlobalOptionInPollID
	case "vote_ids":
		return m.VoteIDs
	case "weight":
		return m.Weight
	case "yes":
		return m.Yes
	}

	return nil
}
