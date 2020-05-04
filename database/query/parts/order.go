package parts

type Order struct {
	Column    string
	Direction string
}

func (order *Order) Compile() string {
	return order.Column + " " + order.Direction
}

func NewOrder(column string, direction string) *Order {
	return &Order{
		Column:    column,
		Direction: direction,
	}
}
