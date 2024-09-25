package datastore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Datastore struct {
	ctx    context.Context
	pool   *pgxpool.Pool
	redis  *redis.Client
	change *changeListenerServer
}

func New(addr string, redisAddr string) (*Datastore, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ds := Datastore{pool: pool, redis: rdb, ctx: context.Background()}
	go ds.setupRedisListener()

	return &ds, nil
}

func (ds *Datastore) Collection(coll interface{}) *query {
	return &query{
		collection: coll,
		datastore:  ds,
		Single:     false,
	}
}
