package models

type Vote struct {
	DelegatedUserID *int
	ID              *int
	MeetingID       int
	OptionID        int
	UserID          *int
	UserToken       string
	Value           *string
	Weight          *string
}

func (m Vote) CollectionName() string {
	return "vote"
}
