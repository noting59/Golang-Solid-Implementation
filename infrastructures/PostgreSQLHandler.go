package infrastructures

import (
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
	"database/sql"
	"fmt"
)

type PostgreSQLHandler struct {
	Conn *sql.DB
}

func (handler *PostgreSQLHandler) Execute(statement string) (sql.Result, error) {
	return handler.Conn.Exec(statement)
}

func (handler *PostgreSQLHandler) Query(statement string) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(PostgreSQLRow),err
	}
	row := new(PostgreSQLRow)
	row.Rows = rows

	return row, nil
}

func (handler *PostgreSQLHandler) Connection () *sql.DB {
	return handler.Conn
}

type PostgreSQLRow struct {
	Rows *sql.Rows
}

func (r PostgreSQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return  nil
}

func (r PostgreSQLRow) Next() bool {
	return r.Rows.Next()
}
