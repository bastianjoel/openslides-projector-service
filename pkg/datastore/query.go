package datastore

import (
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
func (q query) SetIds(ids ...int) query {
	typeName := strcase.ToSnake(reflect.ValueOf(q.collection).Elem().Type().Name())
	for _, id := range ids {
		q.Fqids = append(q.Fqids, fmt.Sprintf("%s/%d", typeName, id))
	}

	return q
}

// Returns a map[string]interface{} if Fields of the query were set otherwise
// the given collection struct
// If single is set to false this will return an array otherwise the first
// found element
func (q query) Run() (interface{}, error) {
	resultIsMap := true
	if q.Fields == nil {
		q.fillFieldsByCollection()
		resultIsMap = false
	}

	val, err := q.datastore.getKeys(q.Fields, q.Fqids)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if resultIsMap {
		result = q.resultMaps(val)
	} else {
		result = q.resultStructs(val)
	}

	if q.Single {
		resultArr := result.([]interface{})
		if len(resultArr) > 0 {
			return resultArr[0], nil
		}

		return nil, nil
	}

	return result, nil
}

func (q *query) fillFieldsByCollection() {
	t := reflect.ValueOf(q.collection).Elem()
	q.Fields = make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		q.Fields[i] = strcase.ToSnake(t.Type().Field(i).Name)
	}
}

func (q *query) resultMaps(keys map[string][]byte) []map[string]interface{} {
	return nil
}

func (q *query) resultStructs(keys map[string][]byte) []interface{} {
	return nil
}
