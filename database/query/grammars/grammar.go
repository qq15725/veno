package grammars

import (
	"fmt"
	"github.com/qq15725/veno/database/contracts"
	"regexp"
	"strconv"
	"strings"
)

type Grammar struct {
}

func (gc *Grammar) CompileColumns(queryBuilder contracts.QueryBuilder) string {
	columns := queryBuilder.GetBuilderColumns()

	if len(columns) == 0 {
		return "select *"
	}

	return "select " + strings.Join(columns, ", ")
}

func (gc *Grammar) CompileFrom(queryBuilder contracts.QueryBuilder) string {
	return "from " + queryBuilder.GetBuilderFrom()
}

func (gr *Grammar) CompileWheres(queryBuilder contracts.QueryBuilder) string {
	wheres := queryBuilder.GetBuilderWheres()

	if len(wheres) == 0 {
		return ""
	}

	compiledWheres := make([]string, 0)

	for _, where := range wheres {
		compiledWheres = append(compiledWheres, where.Compile())
	}

	return "where " + regexp.MustCompile("(?i)^and |or ").ReplaceAllString(strings.Join(compiledWheres, " "), "")
}

func (gr *Grammar) CompileGroups(queryBuilder contracts.QueryBuilder) string {
	groups := queryBuilder.GetBuilderGroups()

	if len(groups) == 0 {
		return ""
	}

	return "group by " + strings.Join(groups, ", ")
}

func (gr *Grammar) CompileOrders(queryBuilder contracts.QueryBuilder) string {
	orders := queryBuilder.GetBuilderOrders()

	if len(orders) == 0 {
		return ""
	}

	compiledOrders := make([]string, 0)

	for _, order := range orders {
		compiledOrders = append(compiledOrders, order.Compile())
	}

	return "order by " + strings.Join(compiledOrders, ", ")
}

func (gr *Grammar) CompileLimit(queryBuilder contracts.QueryBuilder) string {
	limit := queryBuilder.GetBuilderLimit()

	if limit == 0 {
		return ""
	}

	return "limit " + strconv.Itoa(limit)
}

func (gr *Grammar) CompileOffset(queryBuilder contracts.QueryBuilder) string {
	offset := queryBuilder.GetBuilderOffset()

	if offset == 0 {
		return ""
	}

	return "offset " + strconv.Itoa(offset)
}

func (gr *Grammar) CompileSelect(queryBuilder contracts.QueryBuilder) string {
	components := []string{
		gr.CompileColumns(queryBuilder),
		gr.CompileFrom(queryBuilder),
		gr.CompileWheres(queryBuilder),
		gr.CompileGroups(queryBuilder),
		gr.CompileOrders(queryBuilder),
		gr.CompileLimit(queryBuilder),
		gr.CompileOffset(queryBuilder),
	}

	filteredComponents := make([]string, 0)

	for i := 0; i < len(components); i++ {
		if component := components[i]; component != "" {
			filteredComponents = append(filteredComponents, component)
		}
	}

	return strings.Join(filteredComponents, " ")
}

func (gr *Grammar) CompileInsert(queryBuilder contracts.QueryBuilder, values []map[string]interface{}) string {
	columns := make([]string, 0)
	parameterizes := make([]string, 0)

	for column := range values[0] {
		columns = append(columns, column)
	}

	for _, parameter := range values {
		parameterizes = append(parameterizes, "("+gr.parameterize(parameter)+")")
	}

	return fmt.Sprintf(
		"insert into %s (%s) values %s",
		queryBuilder.GetBuilderFrom(),
		gr.columnize(columns),
		strings.Join(parameterizes, ", "),
	)
}

func (gr *Grammar) columnize(columns []string) string {
	columnize := make([]string, 0)
	for _, column := range columns {
		columnize = append(columnize, column)
	}
	return strings.Join(columnize, ", ")
}

func (gr *Grammar) parameterize(values map[string]interface{}) string {
	parameterize := make([]string, 0)
	for _, v := range values {
		parameterize = append(parameterize, gr.parameter(v))
	}
	return strings.Join(parameterize, ", ")
}

func (gr *Grammar) parameter(value interface{}) string {
	return "?"
}
