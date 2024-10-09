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
