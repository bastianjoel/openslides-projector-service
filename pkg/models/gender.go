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

func (m Gender) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "user_ids":
		return m.UserIDs
	}

	return nil
}
