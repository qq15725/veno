package parts

type Where struct {
	Type     string
	Column   string
	Operator string
	Value    interface{}
	Boolean  string
}

func (where *Where) Compile() string {
	return where.Boolean + " " + where.Column + " " + where.Operator + " ?"
}

func NewBasicAndWhere(column string, operator string, value interface{}) *Where {
	return NewWhere("basic", column, operator, value, "and")
}

func NewBasicWhere(column string, operator string, value interface{}, boolean string) *Where {
	return NewWhere("basic", column, operator, value, boolean)
}

func NewWhere(whereType string, column string, operator string, value interface{}, boolean string) *Where {
	return &Where{
		Type:     whereType,
		Column:   column,
		Operator: operator,
		Value:    value,
		Boolean:  boolean,
	}
}
