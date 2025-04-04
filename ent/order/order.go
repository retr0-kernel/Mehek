// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldRequiredBy holds the string denoting the required_by field in the database.
	FieldRequiredBy = "required_by"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// EdgeBrand holds the string denoting the brand edge name in mutations.
	EdgeBrand = "brand"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeResourceAllocations holds the string denoting the resource_allocations edge name in mutations.
	EdgeResourceAllocations = "resource_allocations"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// BrandTable is the table that holds the brand relation/edge.
	BrandTable = "orders"
	// BrandInverseTable is the table name for the Brand entity.
	// It exists in this package in order to avoid circular dependency with the "brand" package.
	BrandInverseTable = "brands"
	// BrandColumn is the table column denoting the brand relation/edge.
	BrandColumn = "brand_orders"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "order_items"
	// ItemsInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	ItemsInverseTable = "order_items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "order_items"
	// ResourceAllocationsTable is the table that holds the resource_allocations relation/edge.
	ResourceAllocationsTable = "resource_allocations"
	// ResourceAllocationsInverseTable is the table name for the ResourceAllocation entity.
	// It exists in this package in order to avoid circular dependency with the "resourceallocation" package.
	ResourceAllocationsInverseTable = "resource_allocations"
	// ResourceAllocationsColumn is the table column denoting the resource_allocations relation/edge.
	ResourceAllocationsColumn = "resource_allocation_order"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldRequiredBy,
	FieldStatus,
	FieldTotalPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"brand_orders",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultTotalPrice holds the default value on creation for the "total_price" field.
	DefaultTotalPrice float64
)

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByRequiredBy orders the results by the required_by field.
func ByRequiredBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiredBy, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTotalPrice orders the results by the total_price field.
func ByTotalPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPrice, opts...).ToFunc()
}

// ByBrandField orders the results by brand field.
func ByBrandField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBrandStep(), sql.OrderByField(field, opts...))
	}
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByResourceAllocationsCount orders the results by resource_allocations count.
func ByResourceAllocationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResourceAllocationsStep(), opts...)
	}
}

// ByResourceAllocations orders the results by resource_allocations terms.
func ByResourceAllocations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourceAllocationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBrandStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BrandInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BrandTable, BrandColumn),
	)
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}
func newResourceAllocationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourceAllocationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ResourceAllocationsTable, ResourceAllocationsColumn),
	)
}
