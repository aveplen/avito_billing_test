package repository

import (
	"context"
	"fmt"

	"github.com/aveplen/avito_test/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDBConnectionPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.SSLMode,
		cfg.Password,
	)
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("db connect failed: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}
	return pool, nil
}
