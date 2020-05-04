package database

import (
	"database/sql"
	"github.com/qq15725/go/database/query/grammars"
)

func Open(driver string, source string) (conn *Connection, err error) {
	sqlDB, err := sql.Open(driver, source)

	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
	}

	conn = NewConnection(sqlDB, &grammars.Grammar{})

	return
}
