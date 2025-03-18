package repository

import (
	"context"
	"log"

	"github.com/iamseki/devices-api/src/repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Queries *queries.Queries
	Pool    *pgxpool.Pool
}

func newPgxPool(connString string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to connect to parse pgxpool config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to parse pgxpool config: %v", err)
	}

	return pool
}

func New(connString string) *Repository {
	pool := newPgxPool(connString)
	queries := queries.New()
	return &Repository{queries, pool}
}
