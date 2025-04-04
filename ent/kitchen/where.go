// Code generated by ent, DO NOT EDIT.

package kitchen

import (
	"project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldName, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldLocation, v))
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldCapacity, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldContainsFold(FieldName, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldContainsFold(FieldLocation, v))
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldEQ(FieldCapacity, v))
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNEQ(FieldCapacity, v))
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldIn(FieldCapacity, vs...))
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldNotIn(FieldCapacity, vs...))
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGT(FieldCapacity, v))
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldGTE(FieldCapacity, v))
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLT(FieldCapacity, v))
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v int) predicate.Kitchen {
	return predicate.Kitchen(sql.FieldLTE(FieldCapacity, v))
}

// HasResources applies the HasEdge predicate on the "resources" edge.
func HasResources() predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourcesWith applies the HasEdge predicate on the "resources" edge with a given conditions (other predicates).
func HasResourcesWith(preds ...predicate.KitchenResource) predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := newResourcesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStaff applies the HasEdge predicate on the "staff" edge.
func HasStaff() predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StaffTable, StaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaffWith applies the HasEdge predicate on the "staff" edge with a given conditions (other predicates).
func HasStaffWith(preds ...predicate.Staff) predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := newStaffStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBrands applies the HasEdge predicate on the "brands" edge.
func HasBrands() predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BrandsTable, BrandsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBrandsWith applies the HasEdge predicate on the "brands" edge with a given conditions (other predicates).
func HasBrandsWith(preds ...predicate.Brand) predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := newBrandsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInventory applies the HasEdge predicate on the "inventory" edge.
func HasInventory() predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InventoryTable, InventoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInventoryWith applies the HasEdge predicate on the "inventory" edge with a given conditions (other predicates).
func HasInventoryWith(preds ...predicate.InventoryItem) predicate.Kitchen {
	return predicate.Kitchen(func(s *sql.Selector) {
		step := newInventoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Kitchen) predicate.Kitchen {
	return predicate.Kitchen(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Kitchen) predicate.Kitchen {
	return predicate.Kitchen(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Kitchen) predicate.Kitchen {
	return predicate.Kitchen(sql.NotPredicates(p))
}
