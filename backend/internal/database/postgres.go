package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool creates a postreSQL connection pool
func NewPool(DatabaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(context.Background(), DatabaseURL)
	if err != nil {
		return nil, err
	}

	//verify database connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
