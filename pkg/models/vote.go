package models

type Vote struct {
	DelegatedUserID *int    `json:"delegated_user_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	OptionID        int     `json:"option_id"`
	UserID          *int    `json:"user_id"`
	UserToken       string  `json:"user_token"`
	Value           *string `json:"value"`
	Weight          *string `json:"weight"`
}

func (m Vote) CollectionName() string {
	return "vote"
}
