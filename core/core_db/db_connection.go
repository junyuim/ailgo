package core_db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type IDbConnection interface {
	sqlx.Queryer
	sqlx.Execer

	DriverName() string

	Get(dest any, query string, args ...any) error

	Select(dest any, query string, args ...any) error

	NamedQuery(query string, arg any) (*sqlx.Rows, error)

	NamedExec(query string, arg any) (sql.Result, error)
}

//
//
//type IDbConnection interface {
//	sqlx.Queryer
//	sqlx.Execer
//
//	DriverName() string
//
//	Get(dest interface{}, query string, args ...interface{}) error
//
//	Select(dest interface{}, query string, args ...interface{}) error
//
//	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
//
//	NamedExec(query string, arg interface{}) (sql.Result, error)
//}
//
//func OpenConnection(dbConfig *DbConfig) (*sqlx.DB, error) {
//	db, err := sqlx.Open(dbConfig.DriverName, dbConfig.DataSourceName)
//
//	if err != nil {
//		fmt.Println("[aix]: open db error", err)
//	}
//
//	return db, err
//}
//
//func UseConnection(dbConfig *DbConfig, f func(*sqlx.DB) error) error {
//	db, err := sqlx.Open(dbConfig.DriverName, dbConfig.DataSourceName)
//
//	if err != nil {
//		fmt.Println("[aix]: open db error", err)
//		return err
//	}
//
//	defer db.Close()
//
//	return f(db)
//}
//
//func UseTransaction(dbConfig *DbConfig, f func(*sqlx.Tx) error) error {
//	db, err := sqlx.Open(dbConfig.DriverName, dbConfig.DataSourceName)
//
//	if err != nil {
//		fmt.Println("[aix]: open db error", err)
//		return err
//	}
//
//	defer db.Close()
//
//	tx, err := db.Beginx()
//
//	if err != nil {
//		fmt.Println("[aix]: open db tx error", err)
//		return err
//	}
//
//	err = f(tx)
//
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//
//	return nil
//}
//
//func UseTransactionWith(db *sqlx.DB, f func(*sqlx.Tx) error) error {
//	tx, err := db.Beginx()
//
//	if err != nil {
//		fmt.Println("[aix]: open db tx error", err)
//		return err
//	}
//
//	err = f(tx)
//
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//
//	return nil
//}
