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
