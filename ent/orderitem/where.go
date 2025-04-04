// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldID, id))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQuantity, v))
}

// SpecialInstructions applies equality check predicate on the "special_instructions" field. It's identical to SpecialInstructionsEQ.
func SpecialInstructions(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldSpecialInstructions, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldQuantity, v))
}

// SpecialInstructionsEQ applies the EQ predicate on the "special_instructions" field.
func SpecialInstructionsEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldSpecialInstructions, v))
}

// SpecialInstructionsNEQ applies the NEQ predicate on the "special_instructions" field.
func SpecialInstructionsNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldSpecialInstructions, v))
}

// SpecialInstructionsIn applies the In predicate on the "special_instructions" field.
func SpecialInstructionsIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldSpecialInstructions, vs...))
}

// SpecialInstructionsNotIn applies the NotIn predicate on the "special_instructions" field.
func SpecialInstructionsNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldSpecialInstructions, vs...))
}

// SpecialInstructionsGT applies the GT predicate on the "special_instructions" field.
func SpecialInstructionsGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldSpecialInstructions, v))
}

// SpecialInstructionsGTE applies the GTE predicate on the "special_instructions" field.
func SpecialInstructionsGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldSpecialInstructions, v))
}

// SpecialInstructionsLT applies the LT predicate on the "special_instructions" field.
func SpecialInstructionsLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldSpecialInstructions, v))
}

// SpecialInstructionsLTE applies the LTE predicate on the "special_instructions" field.
func SpecialInstructionsLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldSpecialInstructions, v))
}

// SpecialInstructionsContains applies the Contains predicate on the "special_instructions" field.
func SpecialInstructionsContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldSpecialInstructions, v))
}

// SpecialInstructionsHasPrefix applies the HasPrefix predicate on the "special_instructions" field.
func SpecialInstructionsHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldSpecialInstructions, v))
}

// SpecialInstructionsHasSuffix applies the HasSuffix predicate on the "special_instructions" field.
func SpecialInstructionsHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldSpecialInstructions, v))
}

// SpecialInstructionsIsNil applies the IsNil predicate on the "special_instructions" field.
func SpecialInstructionsIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldSpecialInstructions))
}

// SpecialInstructionsNotNil applies the NotNil predicate on the "special_instructions" field.
func SpecialInstructionsNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldSpecialInstructions))
}

// SpecialInstructionsEqualFold applies the EqualFold predicate on the "special_instructions" field.
func SpecialInstructionsEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldSpecialInstructions, v))
}

// SpecialInstructionsContainsFold applies the ContainsFold predicate on the "special_instructions" field.
func SpecialInstructionsContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldSpecialInstructions, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenuItem applies the HasEdge predicate on the "menu_item" edge.
func HasMenuItem() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MenuItemTable, MenuItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuItemWith applies the HasEdge predicate on the "menu_item" edge with a given conditions (other predicates).
func HasMenuItemWith(preds ...predicate.MenuItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newMenuItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.NotPredicates(p))
}
