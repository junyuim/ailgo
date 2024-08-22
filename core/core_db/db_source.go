package core_db

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type DbSource struct {
	DriverName     string `json:"driverName" yaml:"driver-name"`
	DataSourceName string `json:"dataSourceName" yaml:"data-source-name"`
}

func (s *DbSource) OpenConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open(s.DriverName, s.DataSourceName)

	if err != nil {
		slog.Error("open db conn", "err", err.Error())
	}

	return db, err
}

func (s *DbSource) UseConnection(f func(*sqlx.DB) error) error {
	db, err := sqlx.Open(s.DriverName, s.DataSourceName)

	if err != nil {
		slog.Error("open db conn", "err", err.Error())
		return err
	}

	defer db.Close()

	return f(db)
}

func (s *DbSource) UseTransaction(db *sqlx.DB, opts *sql.TxOptions, f func(*sqlx.Tx) error) error {
	tx, err := db.BeginTxx(context.Background(), opts)

	if err != nil {
		slog.Error("open db tx", "err", err.Error())
		return err
	}

	err = f(tx)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()

	return nil
}
