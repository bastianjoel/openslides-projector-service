package models

type Vote struct {
	ID              *int
	MeetingID       int
	OptionID        int
	UserID          *int
	UserToken       string
	Value           *string
	Weight          *string
	DelegatedUserID *int
}

func (m Vote) CollectionName() string {
	return "vote"
}
