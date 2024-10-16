package models

import "encoding/json"

type ActionWorker struct {
	Created   int             `json:"created"`
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Result    json.RawMessage `json:"result"`
	State     string          `json:"state"`
	Timestamp int             `json:"timestamp"`
}

func (m ActionWorker) CollectionName() string {
	return "action_worker"
}

func (m ActionWorker) Get(field string) interface{} {
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
	case "timestamp":
		return m.Timestamp
	}

	return nil
}
