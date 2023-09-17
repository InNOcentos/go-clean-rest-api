package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func (postgres *Postgres) Ping(ctx context.Context) error {
	err := postgres.pool.Ping(ctx)

	if err != nil {
		return fmt.Errorf("Error while ping: %w", err)
	}

	return nil
}

func NewClient(url string) (*Postgres, error) {
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing database connection url: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	return &Postgres{
		pool: pool,
	}, nil
}
