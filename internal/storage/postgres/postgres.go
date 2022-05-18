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

func NewStore(connStr string) (DB, error) {
	s, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		return DB{}, fmt.Errorf("connetion string error: %w", err)
	}

	db := stdlib.OpenDB(s)
	conn := sqlx.NewDb(db, "postgres")

	return DB{db: conn}, nil
}

// TODO:
func (s DB) RunMigrations() error {
	return nil
}

func (s DB) DB() *sqlx.DB {
	return s.db
}

func (s DB) Close() error {
	return s.db.Close()
}
