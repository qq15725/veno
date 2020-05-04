package query

import (
	"github.com/qq15725/go/database"
	"github.com/qq15725/go/database/query/grammars"
	"github.com/qq15725/go/database/query/parts"
)

type Builder struct {
	Connection *database.Connection
	Grammar    *grammars.Grammar

	builderFrom    string
	builderColumns []string
	builderWheres  []*parts.Where
	builderGroups  []string
	builderOrders  []*parts.Order
	builderLimit   int
	builderOffset  int

	operators []string

	bindings map[string][]interface{}
}

func (qb *Builder) GetBuilderFrom() string {
	return qb.builderFrom
}

func (qb *Builder) GetBuilderColumns() []string {
	return qb.builderColumns
}

func (qb *Builder) GetBuilderWheres() []*parts.Where {
	return qb.builderWheres
}

func (qb *Builder) GetBuilderGroups() []string {
	return qb.builderGroups
}

func (qb *Builder) GetBuilderOrders() []*parts.Order {
	return qb.builderOrders
}

func (qb *Builder) GetBuilderLimit() int {
	return qb.builderLimit
}

func (qb *Builder) GetBuilderOffset() int {
	return qb.builderOffset
}

func (qb *Builder) GetRawBindings() map[string][]interface{} {
	return qb.bindings
}

func (qb *Builder) GetBindings() []interface{} {
	bindings := make([]interface{}, 0)
	for _, subBindings := range qb.bindings {
		for _, subBinding := range subBindings {
			bindings = append(bindings, subBinding)
		}
	}
	return bindings
}

func (qb *Builder) From(table string) *Builder {
	qb.builderFrom = table
	return qb
}

func (qb *Builder) Select(columns []string) *Builder {
	qb.builderColumns = columns
	return qb
}

func (qb *Builder) Limit(value int) *Builder {
	qb.builderLimit = value
	return qb
}

func (qb *Builder) Offset(value int) *Builder {
	qb.builderOffset = value
	return qb
}

func (qb *Builder) GroupBy(groups ...string) *Builder {
	for _, group := range groups {
		qb.builderGroups = append(qb.builderGroups, group)
	}
	return qb
}

func (qb *Builder) OrderBy(column string, direction string) *Builder {
	if direction == "asc" {
		direction = "asc"
	} else {
		direction = "desc"
	}
	qb.builderOrders = append(qb.builderOrders, parts.NewOrder(column, direction))
	return qb
}

func (qb *Builder) Where(args ...interface{}) *Builder {
	length := len(args)

	var column, operator, value, boolean interface{}

	if length < 2 {
		panic("The parameter should be greater than 2")
	} else if length == 2 {
		column = args[0]
		operator = "="
		value = args[1]
		boolean = "and"
	} else if length == 3 {
		column = args[0]
		operator = args[1]
		value = args[2]
		boolean = "and"
	} else if length == 4 {
		column = args[0]
		operator = args[1]
		value = args[2]
		boolean = args[3]
	}

	if wheres, ok := column.(map[string]interface{}); ok {
		for k, v := range wheres {
			qb.Where(k, "=", v, boolean)
		}

		return qb
	}

	if wheres, ok := column.([][]interface{}); ok {
		for _, v := range wheres {
			qb.Where(v...)
		}

		return qb
	}

	if stringOperator, ok := operator.(string); ok {
		if qb.invalidOperator(stringOperator) {
			value = (interface{})(stringOperator)
			operator = "="
		}
	}

	where := parts.NewBasicWhere(column.(string), operator.(string), value.(interface{}), boolean.(string))

	qb.builderWheres = append(qb.builderWheres, where)

	qb.bindings["where"] = append(qb.bindings["where"], value)

	return qb
}

func (qb *Builder) Get() []map[string]interface{} {
	return qb.Connection.Select(
		qb.ToSql(),
		qb.GetBindings()...,
	)
}

func (qb *Builder) ToSql() string {
	return qb.Grammar.CompileSelect(qb)
}

func (qb *Builder) invalidOperator(operator string) bool {
	for _, otr := range qb.operators {
		if operator == otr {
			return false
		}
	}
	return true
}

func NewBuilder(connection *database.Connection) *Builder {
	return &Builder{
		Connection:     connection,
		Grammar:        connection.Grammar,
		builderColumns: make([]string, 0),
		builderWheres:  make([]*parts.Where, 0),
		builderGroups:  make([]string, 0),
		builderOrders:  make([]*parts.Order, 0),
		operators: []string{
			"=", "<", ">", "<=", ">=", "<>", "!=", "<=>",
			"like", "like binary", "not like", "ilike",
			"&", "|", "^", "<<", ">>",
			"rlike", "regexp", "not regexp",
			"~", "~*", "!~", "!~*", "similar to",
			"not similar to", "not ilike", "~~*", "!~~*",
		},
		bindings: map[string][]interface{}{
			"select": make([]interface{}, 0),
			"from":   make([]interface{}, 0),
			"join":   make([]interface{}, 0),
			"where":  make([]interface{}, 0),
			"having": make([]interface{}, 0),
			"order":  make([]interface{}, 0),
			"union":  make([]interface{}, 0),
		},
	}
}
