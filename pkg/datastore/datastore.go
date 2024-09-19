package datastore

import (
	"context"
	"fmt"

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

func (ds *Datastore) Collection(coll interface{}) *query {
	return &query{
		collection: coll,
		datastore:  ds,
		Single:     false,
	}
}

func (ds *Datastore) getKeys(fqids []string, fields []string) (map[string][]byte, error) {
	sql := fmt.Sprintf(`SELECT t.fqid, row_to_json(t)::jsonb - 'fqid' data FROM (SELECT %s from models where fqid = ANY ($1) AND deleted=false) t`, selectStringFromFields(fields))

	rows, err := ds.pool.Query(ds.ctx, sql, fqids)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	defer rows.Close()

	results := map[string][]byte{}
	for rows.Next() {
		raw := rows.RawValues()
		results[string(raw[0])] = raw[1]
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
	result := "fqid"
	for _, field := range fqids {
		result += ", data->'" + field + "' as " + field
	}

	return result
}
