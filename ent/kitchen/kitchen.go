// Code generated by ent, DO NOT EDIT.

package kitchen

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the kitchen type in the database.
	Label = "kitchen"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldCapacity holds the string denoting the capacity field in the database.
	FieldCapacity = "capacity"
	// FieldOperatingHours holds the string denoting the operating_hours field in the database.
	FieldOperatingHours = "operating_hours"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// EdgeStaff holds the string denoting the staff edge name in mutations.
	EdgeStaff = "staff"
	// EdgeBrands holds the string denoting the brands edge name in mutations.
	EdgeBrands = "brands"
	// EdgeInventory holds the string denoting the inventory edge name in mutations.
	EdgeInventory = "inventory"
	// Table holds the table name of the kitchen in the database.
	Table = "kitchens"
	// ResourcesTable is the table that holds the resources relation/edge.
	ResourcesTable = "kitchen_resources"
	// ResourcesInverseTable is the table name for the KitchenResource entity.
	// It exists in this package in order to avoid circular dependency with the "kitchenresource" package.
	ResourcesInverseTable = "kitchen_resources"
	// ResourcesColumn is the table column denoting the resources relation/edge.
	ResourcesColumn = "kitchen_resources"
	// StaffTable is the table that holds the staff relation/edge.
	StaffTable = "staffs"
	// StaffInverseTable is the table name for the Staff entity.
	// It exists in this package in order to avoid circular dependency with the "staff" package.
	StaffInverseTable = "staffs"
	// StaffColumn is the table column denoting the staff relation/edge.
	StaffColumn = "kitchen_staff"
	// BrandsTable is the table that holds the brands relation/edge.
	BrandsTable = "brands"
	// BrandsInverseTable is the table name for the Brand entity.
	// It exists in this package in order to avoid circular dependency with the "brand" package.
	BrandsInverseTable = "brands"
	// BrandsColumn is the table column denoting the brands relation/edge.
	BrandsColumn = "kitchen_brands"
	// InventoryTable is the table that holds the inventory relation/edge.
	InventoryTable = "inventory_items"
	// InventoryInverseTable is the table name for the InventoryItem entity.
	// It exists in this package in order to avoid circular dependency with the "inventoryitem" package.
	InventoryInverseTable = "inventory_items"
	// InventoryColumn is the table column denoting the inventory relation/edge.
	InventoryColumn = "kitchen_inventory"
)

// Columns holds all SQL columns for kitchen fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLocation,
	FieldCapacity,
	FieldOperatingHours,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	CapacityValidator func(int) error
)

// OrderOption defines the ordering options for the Kitchen queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByCapacity orders the results by the capacity field.
func ByCapacity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapacity, opts...).ToFunc()
}

// ByResourcesCount orders the results by resources count.
func ByResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResourcesStep(), opts...)
	}
}

// ByResources orders the results by resources terms.
func ByResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStaffCount orders the results by staff count.
func ByStaffCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStaffStep(), opts...)
	}
}

// ByStaff orders the results by staff terms.
func ByStaff(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStaffStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBrandsCount orders the results by brands count.
func ByBrandsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBrandsStep(), opts...)
	}
}

// ByBrands orders the results by brands terms.
func ByBrands(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBrandsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInventoryCount orders the results by inventory count.
func ByInventoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInventoryStep(), opts...)
	}
}

// ByInventory orders the results by inventory terms.
func ByInventory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInventoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
	)
}
func newStaffStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StaffInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StaffTable, StaffColumn),
	)
}
func newBrandsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BrandsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BrandsTable, BrandsColumn),
	)
}
func newInventoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InventoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InventoryTable, InventoryColumn),
	)
}
