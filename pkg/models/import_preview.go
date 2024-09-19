package models

import "encoding/json"

type ImportPreview struct {
	Created int             `json:"created"`
	ID      *int            `json:"id"`
	Name    string          `json:"name"`
	Result  json.RawMessage `json:"result"`
	State   string          `json:"state"`
}

func (m ImportPreview) CollectionName() string {
	return "import_preview"
}
