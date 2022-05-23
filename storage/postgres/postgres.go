package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct {
	sqlxdb *sqlx.DB
}

func NewFromConn(db *sql.DB) DB {
	sqlxDB := sqlx.NewDb(db, "postgres")
	return DB{sqlxdb: sqlxDB}
}

func NewFromConnectionString(connStr string) (DB, error) {
	cfg, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		return DB{}, fmt.Errorf("connetion string error: %w", err)
	}
	stdDB := stdlib.OpenDB(cfg)
	sqlxDB := sqlx.NewDb(stdDB, "postgres")

	return DB{sqlxdb: sqlxDB}, nil
}

func (s *DB) RunMigrations() error {

	usr, err := os.Getwd()
	if err != nil {
		return err
	}
	src := fmt.Sprintf("file://%s/storage/postgres/migrations", path.Dir(usr))

	driver, err := postgres.WithInstance(s.sqlxdb.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating pg driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(src, "postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}
	return m.Up()
}

func (s DB) DB() *sqlx.DB {
	return s.sqlxdb
}

func (s DB) Ping() error {
	return s.sqlxdb.Ping()
}

func (s DB) Close() error {
	return s.sqlxdb.Close()
}
