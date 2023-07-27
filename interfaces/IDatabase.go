package interfaces

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type IDatabase interface {
	Execute(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}
