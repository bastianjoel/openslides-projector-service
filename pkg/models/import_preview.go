package models

import "encoding/json"

type ImportPreview struct {
	Created int             `json:"created"`
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Result  json.RawMessage `json:"result"`
	State   string          `json:"state"`
}

func (m ImportPreview) CollectionName() string {
	return "import_preview"
}

func (m ImportPreview) Get(field string) interface{} {
	switch field {
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "result":
		return m.Result
	case "state":
		return m.State
	}

	return nil
}
