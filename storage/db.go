package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

type DBOptions struct {
	ConnectionURL string
}

type DB struct {
	db *sql.DB
}

func NewDB(ctx context.Context, opts *DBOptions) (*DB, error) {
	slog.Info("creating new db connection")
	db, err := sql.Open("postgres", opts.ConnectionURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to postgres: %v", err)
	}

	go func() {
		<-ctx.Done()

		slog.Info("shutting down sql connection")
		db.Close()
	}()

	return &DB{
		db: db,
	}, nil
}
