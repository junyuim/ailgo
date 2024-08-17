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
