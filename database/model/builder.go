package model

import (
	"github.com/qq15725/go/database/query"
	"reflect"
)

type Builder struct {
	model     interface{}
	modelType reflect.Type
	query     *query.Builder
}

func (mb *Builder) SetModel(model interface{}) *Builder {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	mb.model = model
	mb.modelType = modelType
	return mb
}

func (mb *Builder) SetQuery(query *query.Builder) *Builder {
	mb.query = query
	mb.query.From(mb.GetTable())
	return mb
}

func (mb *Builder) GetConnection() string {
	connection, _ := mb.modelType.FieldByName("connection")
	return connection.Tag.Get("model")
}

func (mb *Builder) GetTable() string {
	table, ok := mb.modelType.FieldByName("table")
	if ok {
		return table.Tag.Get("model")
	}
	// TODO 蛇形命名
	return mb.modelType.Name()
}

func (mb *Builder) GetKeyName() string {
	primaryKey, ok := mb.modelType.FieldByName("primaryKey")
	if ok {
		return primaryKey.Tag.Get("model")
	}
	// TODO 蛇形命名
	return mb.modelType.Name() + "_id"
}

func (mb *Builder) Select(columns []string) *Builder {
	mb.query.Select(columns)
	return mb
}

func (mb *Builder) Where(args ...interface{}) *Builder {
	mb.query.Where(args...)
	return mb
}

func (mb *Builder) GroupBy(groups ...string) *Builder {
	mb.query.GroupBy(groups...)
	return mb
}

func (mb *Builder) OrderBy(column string, direction string) *Builder {
	mb.query.OrderBy(column, direction)
	return mb
}

func (mb *Builder) Limit(value int) *Builder {
	mb.query.Limit(value)
	return mb
}

func (mb *Builder) Offset(value int) *Builder {
	mb.query.Offset(value)
	return mb
}

func (mb *Builder) ToSql() string {
	return mb.query.ToSql()
}

func (mb *Builder) Get() []map[string]interface{} {
	return mb.query.Get()
}

func (mb *Builder) First() map[string]interface{} {
	return mb.query.Get()[0]
}

func (mb *Builder) WhereKey(id interface{}) *Builder {
	mb.query.Where(mb.GetKeyName(), "=", id)
	return mb
}

func (mb *Builder) Find(id interface{}) map[string]interface{} {
	return mb.WhereKey(id).First()
}

func NewBuilder() *Builder {
	return &Builder{}
}
