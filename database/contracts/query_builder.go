package contracts

import "github.com/qq15725/go/database/query/parts"

type QueryBuilder interface {
	GetBuilderFrom() string
	GetBuilderColumns() []string
	GetBuilderGroups() []string
	GetBuilderLimit() int
	GetBuilderOffset() int
	GetBuilderWheres() []*parts.Where
	GetBuilderOrders() []*parts.Order
}
