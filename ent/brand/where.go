// Code generated by ent, DO NOT EDIT.

package brand

import (
	"project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldName, v))
}

// CuisineType applies equality check predicate on the "cuisine_type" field. It's identical to CuisineTypeEQ.
func CuisineType(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldCuisineType, v))
}

// LogoURL applies equality check predicate on the "logo_url" field. It's identical to LogoURLEQ.
func LogoURL(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldLogoURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContainsFold(FieldName, v))
}

// CuisineTypeEQ applies the EQ predicate on the "cuisine_type" field.
func CuisineTypeEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldCuisineType, v))
}

// CuisineTypeNEQ applies the NEQ predicate on the "cuisine_type" field.
func CuisineTypeNEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldCuisineType, v))
}

// CuisineTypeIn applies the In predicate on the "cuisine_type" field.
func CuisineTypeIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldCuisineType, vs...))
}

// CuisineTypeNotIn applies the NotIn predicate on the "cuisine_type" field.
func CuisineTypeNotIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldCuisineType, vs...))
}

// CuisineTypeGT applies the GT predicate on the "cuisine_type" field.
func CuisineTypeGT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldCuisineType, v))
}

// CuisineTypeGTE applies the GTE predicate on the "cuisine_type" field.
func CuisineTypeGTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldCuisineType, v))
}

// CuisineTypeLT applies the LT predicate on the "cuisine_type" field.
func CuisineTypeLT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldCuisineType, v))
}

// CuisineTypeLTE applies the LTE predicate on the "cuisine_type" field.
func CuisineTypeLTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldCuisineType, v))
}

// CuisineTypeContains applies the Contains predicate on the "cuisine_type" field.
func CuisineTypeContains(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContains(FieldCuisineType, v))
}

// CuisineTypeHasPrefix applies the HasPrefix predicate on the "cuisine_type" field.
func CuisineTypeHasPrefix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasPrefix(FieldCuisineType, v))
}

// CuisineTypeHasSuffix applies the HasSuffix predicate on the "cuisine_type" field.
func CuisineTypeHasSuffix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasSuffix(FieldCuisineType, v))
}

// CuisineTypeEqualFold applies the EqualFold predicate on the "cuisine_type" field.
func CuisineTypeEqualFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEqualFold(FieldCuisineType, v))
}

// CuisineTypeContainsFold applies the ContainsFold predicate on the "cuisine_type" field.
func CuisineTypeContainsFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContainsFold(FieldCuisineType, v))
}

// LogoURLEQ applies the EQ predicate on the "logo_url" field.
func LogoURLEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldLogoURL, v))
}

// LogoURLNEQ applies the NEQ predicate on the "logo_url" field.
func LogoURLNEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldLogoURL, v))
}

// LogoURLIn applies the In predicate on the "logo_url" field.
func LogoURLIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldLogoURL, vs...))
}

// LogoURLNotIn applies the NotIn predicate on the "logo_url" field.
func LogoURLNotIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldLogoURL, vs...))
}

// LogoURLGT applies the GT predicate on the "logo_url" field.
func LogoURLGT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldLogoURL, v))
}

// LogoURLGTE applies the GTE predicate on the "logo_url" field.
func LogoURLGTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldLogoURL, v))
}

// LogoURLLT applies the LT predicate on the "logo_url" field.
func LogoURLLT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldLogoURL, v))
}

// LogoURLLTE applies the LTE predicate on the "logo_url" field.
func LogoURLLTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldLogoURL, v))
}

// LogoURLContains applies the Contains predicate on the "logo_url" field.
func LogoURLContains(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContains(FieldLogoURL, v))
}

// LogoURLHasPrefix applies the HasPrefix predicate on the "logo_url" field.
func LogoURLHasPrefix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasPrefix(FieldLogoURL, v))
}

// LogoURLHasSuffix applies the HasSuffix predicate on the "logo_url" field.
func LogoURLHasSuffix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasSuffix(FieldLogoURL, v))
}

// LogoURLEqualFold applies the EqualFold predicate on the "logo_url" field.
func LogoURLEqualFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEqualFold(FieldLogoURL, v))
}

// LogoURLContainsFold applies the ContainsFold predicate on the "logo_url" field.
func LogoURLContainsFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContainsFold(FieldLogoURL, v))
}

// HasKitchen applies the HasEdge predicate on the "kitchen" edge.
func HasKitchen() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, KitchenTable, KitchenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasKitchenWith applies the HasEdge predicate on the "kitchen" edge with a given conditions (other predicates).
func HasKitchenWith(preds ...predicate.Kitchen) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newKitchenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenus applies the HasEdge predicate on the "menus" edge.
func HasMenus() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MenusTable, MenusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenusWith applies the HasEdge predicate on the "menus" edge with a given conditions (other predicates).
func HasMenusWith(preds ...predicate.Menu) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newMenusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Brand) predicate.Brand {
	return predicate.Brand(sql.NotPredicates(p))
}
