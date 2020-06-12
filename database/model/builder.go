package model

import (
	"github.com/qq15725/veno/database/query"
	"github.com/qq15725/veno/support/refl"
	"reflect"
)

type Builder struct {
	model   interface{}
	modelT  reflect.Type
	query   *query.Builder
	tagName string
}

func (mb *Builder) SetModel(model interface{}) *Builder {
	mb.model = model
	mb.modelT = refl.IndirectType(reflect.TypeOf(model))
	return mb
}

func (mb *Builder) SetQuery(query *query.Builder) *Builder {
	mb.query = query
	mb.query.From(mb.GetTable())
	return mb
}

func (mb *Builder) GetConnection() string {
	return refl.FieldTagValue(mb.modelT, "connection", mb.tagName)
}

func (mb *Builder) GetTable() string {
	if v := refl.FieldTagValue(mb.modelT, "table", mb.tagName); v != "" {
		return v
	}
	// TODO 蛇形命名
	return mb.modelT.Name()
}

func (mb *Builder) GetKeyName() string {
	if v := refl.FieldTagValue(mb.modelT, "primaryKey", mb.tagName); v != "" {
		return v
	}
	// TODO 蛇形命名
	return mb.modelT.Name() + "_id"
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

func (mb *Builder) Get() []interface{} {
	models := make([]interface{}, 0)
	rows := mb.query.Get()
	for _, row := range rows {
		model := reflect.New(mb.modelT).Elem()
		for k, v := range row {
			refl.SetFieldValue(model, k, v)
		}
		models = append(models, model.Interface())
	}
	return models
}

func (mb *Builder) First() interface{} {
	return mb.Limit(1).Get()[0]
}

func (mb *Builder) WhereKey(id interface{}) *Builder {
	mb.query.Where(mb.GetKeyName(), "=", id)
	return mb
}

func (mb *Builder) Find(id interface{}) interface{} {
	return mb.WhereKey(id).First()
}

func NewBuilder() *Builder {
	return &Builder{
		tagName: "model",
	}
}
