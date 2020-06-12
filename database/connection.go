package database

import (
	"database/sql"
	"github.com/qq15725/veno/database/query/grammars"
)

type Connection struct {
	sqlDB   *sql.DB
	Grammar *grammars.Grammar
}

func (conn *Connection) insert(query string, bindings []interface{}) (sql.Result, error) {
	stmt, err := conn.sqlDB.Prepare(query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(bindings...)
	if err != nil {
		return nil, err
	}
	stmt.Close()
	return res, nil
}

func (conn *Connection) SelectOne(query string, bindings []interface{}) map[string]interface{} {
	return conn.Select(query, bindings...)[0]
}

func (conn *Connection) Select(query string, bindings ...interface{}) []map[string]interface{} {
	Rows, _ := conn.sqlDB.Query(query, bindings...)
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
	Rows.Close()
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
