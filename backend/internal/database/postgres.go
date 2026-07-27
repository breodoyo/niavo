package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool creates a postreSQL connection pool
func NewPool(DatabaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), DatabaseURL)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
