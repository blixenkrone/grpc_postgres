package postgres

import (
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(connStr string) (DB, error) {
	s, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		return DB{}, fmt.Errorf("connetion string error: %w", err)
	}

	stdDB := stdlib.OpenDB(s)
	sqlxDB := sqlx.NewDb(stdDB, "postgres")

	return DB{db: sqlxDB}, nil
}

// TODO:
func (s DB) RunMigrations() error {
	return nil
}

func (s DB) DB() *sqlx.DB {
	return s.db
}

func (s DB) Ping() error {
	return s.db.Ping()
}

func (s DB) Close() error {
	return s.db.Close()
}
