package models

import "encoding/json"

type ImportPreview struct {
	Created int
	ID      *int
	Name    string
	Result  json.RawMessage
	State   string
}

func (m ImportPreview) CollectionName() string {
	return "import_preview"
}
