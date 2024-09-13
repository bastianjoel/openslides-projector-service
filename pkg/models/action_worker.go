package models

import "encoding/json"

type ActionWorker struct {
	Result    json.RawMessage
	State     string
	Timestamp int
	Created   int
	ID        *int
	Name      string
}

func (m ActionWorker) CollectionName() string {
	return "action_worker"
}
