package datastore

import (
	"encoding/json"
	"reflect"

	"github.com/rs/zerolog/log"
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
						log.Error().Err(err).Msgf("parsing subscription field `%s` for fqid `%s` with value %s", key, fqid, val)
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
		val := reflect.ValueOf(field)
		if val.Kind() != reflect.Ptr {
			log.Fatal().Msg("value passed to SubscribeField must be a pointer")
		}

		data, err := q.GetOne()
		if err != nil {
			log.Warn().Msgf("failed initial data fetch for subscription on %s", q.fqids[0])
		} else if data != nil {
			val.Elem().Set(reflect.ValueOf((*data).Get(q.Fields[0])))
			notifyChannel <- struct{}{}
		}

		for update := range updateChannel {
			if obj, ok := update[q.fqids[0]]; ok {
				if val, ok := obj[q.Fields[0]]; ok {
					err := json.Unmarshal([]byte(val), field)
					if err != nil {
						log.Error().Err(err).Msg("parsing subscription field")
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
