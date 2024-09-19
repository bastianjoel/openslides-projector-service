package datastore

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
)

type query struct {
	collection interface{}
	datastore  *Datastore
	Fields     []string
	Fqids      []string
	Single     bool
}

// Sets the fqids of the query by plain ids
func (q *query) SetIds(ids ...int) *query {
	typeName := strcase.ToSnake(reflect.ValueOf(q.collection).Elem().Type().Name())
	for _, id := range ids {
		q.Fqids = append(q.Fqids, fmt.Sprintf("%s/%d", typeName, id))
	}

	return q
}

func (q *query) AsSingle() *query {
	q.Single = true
	return q
}

// Returns a map[string]interface{} if Fields of the query were set otherwise
// the given collection struct
// If single is set to false this will return an array otherwise the first
// found element
func (q *query) Run() (interface{}, error) {
	resultIsMap := q.Fields != nil

	var result interface{}
	var err error
	if resultIsMap {
		result, err = q.resultMaps()
		if err != nil {
			return nil, fmt.Errorf("getting result as map: %w", err)
		}
	} else {
		result, err = q.resultStructs()
		if err != nil {
			return nil, fmt.Errorf("getting result as structs: %w", err)
		}
	}

	if q.Single {
		resultSlice := reflect.ValueOf(result)
		if resultSlice.Len() > 0 {
			return resultSlice.Index(0).Interface(), nil
		}

		return nil, nil
	}

	return result, nil
}

func (q *query) resultMaps() ([]map[string]interface{}, error) {
	data, err := q.datastore.getKeys(q.Fqids, q.Fields)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for _, row := range data {
		result := map[string]interface{}{}
		if err := json.Unmarshal(row, &result); err != nil {
			return nil, err
		}
		// TODO: Fix oversized fields
		results = append(results, result)
	}

	return results, nil
}

func (q *query) resultStructs() (any, error) {
	dsResults, err := q.datastore.getFull(q.Fqids)
	if err != nil {
		return nil, err
	}

	t := reflect.ValueOf(q.collection).Elem().Type()
	results := reflect.MakeSlice(reflect.SliceOf(t), 0, len(dsResults))
	for _, dsResult := range dsResults {
		el := reflect.New(t)
		err := json.Unmarshal(dsResult, el.Interface())
		if err != nil {
			return nil, err
		}
		results = reflect.Append(results, el.Elem())
	}

	return results.Interface(), nil
}
