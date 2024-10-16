package datastore

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
)

type baseModel interface {
	CollectionName() string
	Get(field string) interface{}
	Update(data map[string]string) error
}

type query[T baseModel] struct {
	collection *T
	datastore  *Datastore
	Fields     []string
	fqids      []string
}

func (q *query[T]) SetFqids(ids ...string) *query[T] {
	q.fqids = append(q.fqids, ids...)

	return q
}

// Sets the fqids of the query by plain ids
func (q *query[T]) SetIds(ids ...int) *query[T] {
	typeName := strcase.ToSnake(reflect.ValueOf(q.collection).Elem().Type().Name())
	for _, id := range ids {
		q.fqids = append(q.fqids, fmt.Sprintf("%s/%d", typeName, id))
	}

	return q
}

// Sets the fqids of the query by plain ids
func (q *query[T]) SetFields(fields ...string) *query[T] {
	q.Fields = append(q.Fields, fields...)

	return q
}

func (q *query[T]) GetOne() (*T, error) {
	result, err := q.resultStructs()
	if err != nil {
		return nil, fmt.Errorf("getting single result as struct: %w", err)
	}

	for _, fqid := range q.fqids {
		if val, ok := result[fqid]; ok {
			return val, nil
		}
	}

	return nil, nil
}

func (q *query[T]) Get() (map[string]*T, error) {
	result, err := q.resultStructs()
	if err != nil {
		return nil, fmt.Errorf("getting results as structs: %w", err)
	}

	return result, nil
}

func (q *query[T]) GetMaps() (map[string]map[string]interface{}, error) {
	result, err := q.resultMaps()
	if err != nil {
		return nil, fmt.Errorf("getting results as maps: %w", err)
	}

	return result, nil
}

func (q *query[T]) resultMaps() (map[string]map[string]interface{}, error) {
	data, err := q.datastore.getKeys(q.fqids, q.Fields)
	if err != nil {
		return nil, err
	}

	results := map[string]map[string]interface{}{}
	for fqid, row := range data {
		result := map[string]interface{}{}
		if err := json.Unmarshal(row, &result); err != nil {
			return nil, err
		}
		// TODO: Fix oversized fields
		results[fqid] = result
	}

	return results, nil
}

func (q *query[T]) resultStructs() (map[string]*T, error) {
	dsResults, err := q.datastore.getFull(q.fqids)
	if err != nil {
		return nil, err
	}

	results := map[string]*T{}
	for fqid, dsResult := range dsResults {
		el := new(T)
		err := json.Unmarshal(dsResult, el)
		if err != nil {
			return nil, err
		}
		results[fqid] = el
	}

	return results, nil
}
