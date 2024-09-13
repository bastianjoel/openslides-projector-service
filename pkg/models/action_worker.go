package models

import "encoding/json"

type ActionWorker struct {
	Created   int
	ID        *int
	Name      string
	Result    json.RawMessage
	State     string
	Timestamp int
}

func (m ActionWorker) CollectionName() string {
	return "action_worker"
}
