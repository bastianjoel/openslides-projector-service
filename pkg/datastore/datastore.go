package datastore

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Datastore struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func New(addr string) (*Datastore, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	ds := Datastore{pool: pool, ctx: context.Background()}
	return &ds, nil
}

func (ds *Datastore) Collection(coll interface{}) query {
	return query{
		collection: coll,
		datastore:  ds,
		Single:     false,
	}
}

func (ds *Datastore) getKeys(fields []string, fqids []string) ([]map[string][]byte, error) {
	sql := fmt.Sprintf(`SELECT %s from models where fqid = ANY ($1) AND deleted=false;`, selectStringFromFields(fields))

	rows, err := ds.pool.Query(ds.ctx, sql, fqids)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	defer rows.Close()

	var results []map[string][]byte
	for rows.Next() {
		raw := rows.RawValues()
		result := map[string][]byte{}
		result["fqid"] = raw[0]
		for i, field := range fields {
			result[field] = raw[i+1]
		}
		results = append(results, result)
	}

	return results, nil
}

func (ds *Datastore) getFull(fqids []string) (map[string][]byte, error) {
	sql := `SELECT fqid, data from models where fqid = ANY ($1) AND deleted=false;`

	rows, err := ds.pool.Query(ds.ctx, sql, fqids)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	defer rows.Close()

	result := map[string][]byte{}
	for rows.Next() {
		raw := rows.RawValues()
		result[string(raw[0])] = raw[1]
	}

	return result, nil
}

func selectStringFromFields(fqids []string) string {
	result := make([]string, len(fqids)+1)
	result[0] = "fqid"
	for i, field := range fqids {
		result[i+1] = "data->>'" + field + "'"
	}

	return strings.Join(result, ", ")
}
