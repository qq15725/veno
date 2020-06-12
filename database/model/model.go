package model

import (
	"github.com/qq15725/veno/database"
	"github.com/qq15725/veno/database/query"
)

type Model struct {
	connection string `model:"mysql"`
	primaryKey uint   `model:"id"`
}

func newQueryBuilder(connectionName string) *query.Builder {
	connection, ok := connections[connectionName]
	if !ok {
		connector, _ := cfg.GetConnection(connectionName)
		connection, _ = database.Open(connector.Driver(), connector.DataSourceName())
		connections[connectionName] = connection
	}
	return query.NewBuilder(connection)
}

func newModelBuilder(model interface{}) *Builder {
	return NewBuilder().SetModel(model)
}

func NewQuery(model interface{}) *Builder {
	modelBuilder := newModelBuilder(model)
	modelBuilder.SetQuery(newQueryBuilder(modelBuilder.GetConnection()))
	return modelBuilder
}

func Query(model interface{}) *Builder {
	return NewQuery(model)
}
