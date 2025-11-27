package storage

import "context"

func (db *DB) CreateUser(ctx context.Context, email string, passwordHash []byte) error {
	_, err := db.db.ExecContext(ctx, "insert into users (email, password_hash) values ($1, $2)", email, passwordHash)
	return err
}
