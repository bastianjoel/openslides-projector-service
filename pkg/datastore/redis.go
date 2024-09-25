package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type queryChangeListener struct {
	q       *query
	channel chan map[string]map[string]interface{}
}

type changeListenerServer struct {
	AddListener    chan *queryChangeListener
	RemoveListener chan (<-chan map[string]map[string]interface{})
}

func (ds *Datastore) setupRedisListener() {
	ds.change = &changeListenerServer{
		AddListener:    make(chan *queryChangeListener),
		RemoveListener: make(chan (<-chan map[string]map[string]interface{})),
	}
	c := ds.change

	changeSource := make(chan map[string]map[string]interface{})
	go setupChangeListener(ds.ctx, ds.redis, changeSource)

	listeners := make([]*queryChangeListener, 0)
	defer func() {
		for _, listener := range listeners {
			close(listener.channel)
		}
		close(changeSource)
	}()

	for {
		select {
		case <-ds.ctx.Done():
			return
		case listener := <-c.AddListener:
			listeners = append(listeners, listener)
		case channel := <-c.RemoveListener:
			for i, listener := range listeners {
				if listener.channel == channel {
					listeners[i] = listeners[len(listeners)-1]
					listeners = listeners[:len(listeners)-1]
					close(listener.channel)
					break
				}
			}
		case changeMap, ok := <-changeSource:
			if !ok {
				return
			}

			for _, listener := range listeners {
				listenerChanged := map[string]map[string]interface{}{}
				for _, fqid := range listener.q.Fqids {
					if val, ok := changeMap[fqid]; ok {
						listenerChanged[fqid] = val
					}
				}
				if len(listenerChanged) > 0 {
					listener.channel <- listenerChanged
				}
			}
		}
	}
}

func setupChangeListener(ctx context.Context, conn *redis.Client, channel chan map[string]map[string]interface{}) {
	id := "$"

	for ctx.Err() == nil {
		nextID, data, err := getNextRedisUpdate(ctx, conn, id)
		if err != nil {
			if nextID != "" {
				id = nextID
			}

			fmt.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		if data != nil {
			channel <- data
		}
		id = nextID
	}
}

func getNextRedisUpdate(ctx context.Context, conn *redis.Client, id string) (string, map[string]map[string]interface{}, error) {
	reply, err := conn.XRead(ctx, &redis.XReadArgs{
		Block:   0,
		Streams: []string{"ModifiedFields", id},
		Count:   10,
	}).Result()
	if err != nil {
		return "", nil, fmt.Errorf("redis `XREAD %s BLOCK 0 STREAMS %s: %w", "ModifiedFields", id, err)
	}

	if reply == nil {
		// This happens, when the redis command times out.
		return id, nil, nil
	}

	id, data, err := parseMessageBus(reply)
	if err != nil {
		return id, nil, fmt.Errorf("parsing message bus: %w", err)
	}

	return id, data, nil
}

func parseMessageBus(streams []redis.XStream) (string, map[string]map[string]interface{}, error) {
	data := map[string]map[string]interface{}{}

	nextID := ""
	for _, stream := range streams {
		for _, message := range stream.Messages {
			nextID = message.ID
			for k, v := range message.Values {
				sepIndex := strings.LastIndex(k, "/")
				if sepIndex < 0 {
					fmt.Println(fmt.Errorf("invalid fqid %s", k))
				}

				var val interface{}
				err := json.Unmarshal([]byte(v.(string)), &val)
				if err != nil {
					fmt.Println(fmt.Errorf("error parsing message bus field: %w", err))
				}

				if _, ok := data[k[:sepIndex]]; !ok {
					data[k[:sepIndex]] = map[string]interface{}{}
				}

				// Element is marked as deleted
				if data[k[:sepIndex]] == nil {
					continue
				}

				// Mark item as deleted
				if k[sepIndex+1:] == "id" && val == nil {
					data[k[:sepIndex]] = nil
				} else {
					data[k[:sepIndex]][k[sepIndex+1:]] = val
				}
			}
		}
	}

	if nextID == "" {
		return "", nil, fmt.Errorf("did not receive new id")
	}

	return nextID, data, nil
}
