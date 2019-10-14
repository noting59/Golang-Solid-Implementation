package interfaces

import "database/sql"

type IDbHandler interface {
	Execute(statement string)(sql.Result, error)
	Query(statement string) (IRow, error)
	Connection() *sql.DB
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
