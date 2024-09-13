package datastore

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/models"
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

func (ds *Datastore) getKeys(fields []string, fqids []string) (map[string][]byte, error) {
	sql := fmt.Sprintf(`SELECT %s from models where fqid = ANY ($1) AND deleted=false;`, selectStringFromFields(fields))

	rows, err := ds.pool.Query(ds.ctx, sql, fqids)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	defer rows.Close()

	projectors, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (interface{}, error) {
		var projector models.Projector
		return projector, nil
	})
	fmt.Println(err)
	fmt.Println(projectors)
	/*
		println(sql)
		for it.Next() {
			r := it.RawValues()
			err := it.Scan(&projector)
			fmt.Println(err)
			fmt.Println(projector)
			println(string(r[0]))
		}
	*/

	// TODO: Implement data fetch from db
	return map[string][]byte{}, nil
}

func selectStringFromFields(fqids []string) string {
	result := "fqid"
	for _, field := range fqids {
		result += ", data->>'" + field + "' as " + field
	}

	return result
}
