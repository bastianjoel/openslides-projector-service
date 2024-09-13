package models

type OrganizationTag struct {
	Color          string
	ID             *int
	Name           string
	OrganizationID int
	TaggedIDs      []string
}

func (m OrganizationTag) CollectionName() string {
	return "organization_tag"
}
