package datastore

import (
	"encoding/json"
	"fmt"
)

type subscription[V any] struct {
	db            *Datastore
	updateChannel chan map[string]map[string]string
	Channel       V
	Unsubscribe   func()
}

func (q *query[T]) Subscribe() *subscription[<-chan map[string]map[string]interface{}] {
	notifyChannel := make(chan map[string]map[string]interface{})
	updateChannel := make(chan map[string]map[string]string)
	listener := queryChangeListener{
		fqids:   q.fqids,
		fields:  q.Fields,
		channel: updateChannel,
	}
	q.datastore.change.AddListener <- &listener
	go func() {
		for update := range updateChannel {
			next := make(map[string]map[string]interface{})
			for fqid, fields := range update {
				if fields == nil {
					next[fqid] = nil
					continue
				}

				next[fqid] = make(map[string]interface{})
				for key, val := range fields {
					var fieldData interface{}
					err := json.Unmarshal([]byte(val), &fieldData)
					if err != nil {
						fmt.Println(err, key, fqid, val, next[fqid])
					}

					next[fqid][key] = fieldData
				}
			}
			notifyChannel <- next
		}

		q.datastore.change.RemoveListener <- updateChannel
	}()

	return &subscription[<-chan map[string]map[string]interface{}]{q.datastore, updateChannel, notifyChannel, func() {
		close(notifyChannel)
	}}
}

func (q *query[T]) SubscribeField(field interface{}) *subscription[<-chan struct{}] {
	notifyChannel := make(chan struct{})
	updateChannel := make(chan map[string]map[string]string)
	listener := queryChangeListener{
		fqids:   q.fqids,
		fields:  q.Fields,
		channel: updateChannel,
	}
	q.datastore.change.AddListener <- &listener

	go func() {
		for update := range updateChannel {
			if obj, ok := update[q.fqids[0]]; ok {
				if val, ok := obj[q.Fields[0]]; ok {
					err := json.Unmarshal([]byte(val), field)
					if err != nil {
						fmt.Println(err)
					}

					notifyChannel <- struct{}{}
				}
			}
		}

		q.datastore.change.RemoveListener <- updateChannel
	}()

	return &subscription[<-chan struct{}]{q.datastore, updateChannel, notifyChannel, func() {
		close(notifyChannel)
	}}
}
