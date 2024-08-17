package core_db

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/junyuim/ailgo/core/core_utils"
)

type DbSource struct {
	DriverName     string `json:"driverName" yaml:"driver-name"`
	DataSourceName string `json:"dataSourceName" yaml:"data-source-name"`
}

func (s *DbSource) OpenConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open(s.DriverName, s.DataSourceName)

	if err != nil {
		core_utils.LogError("open connection error:%s", err.Error())
	}

	return db, err
}

func (s *DbSource) UseConnection(f func(*sqlx.DB) error) error {
	db, err := sqlx.Open(s.DriverName, s.DataSourceName)

	if err != nil {
		core_utils.LogError("open connection error:%s", err.Error())
		return err
	}

	defer db.Close()

	return f(db)
}

func (s *DbSource) UseTransaction(db *sqlx.DB, opts *sql.TxOptions, f func(*sqlx.Tx) error) error {
	tx, err := db.BeginTxx(context.Background(), opts)

	if err != nil {
		core_utils.LogError("open db tx error:%s", err.Error())
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
