package models

type Gender struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	UserIDs        []int  `json:"user_ids"`
}

func (m Gender) CollectionName() string {
	return "gender"
}
