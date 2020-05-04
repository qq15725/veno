package database

import (
	"database/sql"
	"github.com/qq15725/go/database/query/grammars"
)

type Connection struct {
	sqlDB   *sql.DB
	Grammar *grammars.Grammar
}

func (conn *Connection) Select(query string, args ...interface{}) []map[string]interface{} {
	Rows, _ := conn.sqlDB.Query(query, args...)
	cols, _ := Rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range scans {
		scans[i] = &values[i]
	}
	rows := make([]map[string]interface{}, 0)
	for Rows.Next() {
		Rows.Scan(scans...)
		row := make(map[string]interface{})
		for k, v := range values {
			row[cols[k]] = string(v)
		}
		rows = append(rows, row)
	}
	return rows
}

func (conn *Connection) Close() error {
	return conn.sqlDB.Close()
}

func NewConnection(sqlDB *sql.DB, grammar *grammars.Grammar) *Connection {
	return &Connection{
		sqlDB:   sqlDB,
		Grammar: grammar,
	}
}
