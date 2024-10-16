package models

type OrganizationTag struct {
	Color          string   `json:"color"`
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	OrganizationID int      `json:"organization_id"`
	TaggedIDs      []string `json:"tagged_ids"`
}

func (m OrganizationTag) CollectionName() string {
	return "organization_tag"
}

func (m OrganizationTag) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "tagged_ids":
		return m.TaggedIDs
	}

	return nil
}
