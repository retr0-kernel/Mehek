// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project/ent/menuitem"
	"project/ent/order"
	"project/ent/orderitem"
	"project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetQuantity sets the "quantity" field.
func (oiu *OrderItemUpdate) SetQuantity(i int) *OrderItemUpdate {
	oiu.mutation.ResetQuantity()
	oiu.mutation.SetQuantity(i)
	return oiu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableQuantity(i *int) *OrderItemUpdate {
	if i != nil {
		oiu.SetQuantity(*i)
	}
	return oiu
}

// AddQuantity adds i to the "quantity" field.
func (oiu *OrderItemUpdate) AddQuantity(i int) *OrderItemUpdate {
	oiu.mutation.AddQuantity(i)
	return oiu
}

// SetSpecialInstructions sets the "special_instructions" field.
func (oiu *OrderItemUpdate) SetSpecialInstructions(s string) *OrderItemUpdate {
	oiu.mutation.SetSpecialInstructions(s)
	return oiu
}

// SetNillableSpecialInstructions sets the "special_instructions" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableSpecialInstructions(s *string) *OrderItemUpdate {
	if s != nil {
		oiu.SetSpecialInstructions(*s)
	}
	return oiu
}

// ClearSpecialInstructions clears the value of the "special_instructions" field.
func (oiu *OrderItemUpdate) ClearSpecialInstructions() *OrderItemUpdate {
	oiu.mutation.ClearSpecialInstructions()
	return oiu
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oiu *OrderItemUpdate) SetOrderID(id int) *OrderItemUpdate {
	oiu.mutation.SetOrderID(id)
	return oiu
}

// SetOrder sets the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) SetOrder(o *Order) *OrderItemUpdate {
	return oiu.SetOrderID(o.ID)
}

// SetMenuItemID sets the "menu_item" edge to the MenuItem entity by ID.
func (oiu *OrderItemUpdate) SetMenuItemID(id int) *OrderItemUpdate {
	oiu.mutation.SetMenuItemID(id)
	return oiu
}

// SetMenuItem sets the "menu_item" edge to the MenuItem entity.
func (oiu *OrderItemUpdate) SetMenuItem(m *MenuItem) *OrderItemUpdate {
	return oiu.SetMenuItemID(m.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) ClearOrder() *OrderItemUpdate {
	oiu.mutation.ClearOrder()
	return oiu
}

// ClearMenuItem clears the "menu_item" edge to the MenuItem entity.
func (oiu *OrderItemUpdate) ClearMenuItem() *OrderItemUpdate {
	oiu.mutation.ClearMenuItem()
	return oiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiu *OrderItemUpdate) check() error {
	if v, ok := oiu.mutation.Quantity(); ok {
		if err := orderitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "OrderItem.quantity": %w`, err)}
		}
	}
	if oiu.mutation.OrderCleared() && len(oiu.mutation.OrderIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.order"`)
	}
	if oiu.mutation.MenuItemCleared() && len(oiu.mutation.MenuItemIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.menu_item"`)
	}
	return nil
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.SpecialInstructions(); ok {
		_spec.SetField(orderitem.FieldSpecialInstructions, field.TypeString, value)
	}
	if oiu.mutation.SpecialInstructionsCleared() {
		_spec.ClearField(orderitem.FieldSpecialInstructions, field.TypeString)
	}
	if oiu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiu.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.MenuItemTable,
			Columns: []string{orderitem.MenuItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.MenuItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.MenuItemTable,
			Columns: []string{orderitem.MenuItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderItemMutation
}

// SetQuantity sets the "quantity" field.
func (oiuo *OrderItemUpdateOne) SetQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.ResetQuantity()
	oiuo.mutation.SetQuantity(i)
	return oiuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableQuantity(i *int) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetQuantity(*i)
	}
	return oiuo
}

// AddQuantity adds i to the "quantity" field.
func (oiuo *OrderItemUpdateOne) AddQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.AddQuantity(i)
	return oiuo
}

// SetSpecialInstructions sets the "special_instructions" field.
func (oiuo *OrderItemUpdateOne) SetSpecialInstructions(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetSpecialInstructions(s)
	return oiuo
}

// SetNillableSpecialInstructions sets the "special_instructions" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableSpecialInstructions(s *string) *OrderItemUpdateOne {
	if s != nil {
		oiuo.SetSpecialInstructions(*s)
	}
	return oiuo
}

// ClearSpecialInstructions clears the value of the "special_instructions" field.
func (oiuo *OrderItemUpdateOne) ClearSpecialInstructions() *OrderItemUpdateOne {
	oiuo.mutation.ClearSpecialInstructions()
	return oiuo
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oiuo *OrderItemUpdateOne) SetOrderID(id int) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderID(id)
	return oiuo
}

// SetOrder sets the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) SetOrder(o *Order) *OrderItemUpdateOne {
	return oiuo.SetOrderID(o.ID)
}

// SetMenuItemID sets the "menu_item" edge to the MenuItem entity by ID.
func (oiuo *OrderItemUpdateOne) SetMenuItemID(id int) *OrderItemUpdateOne {
	oiuo.mutation.SetMenuItemID(id)
	return oiuo
}

// SetMenuItem sets the "menu_item" edge to the MenuItem entity.
func (oiuo *OrderItemUpdateOne) SetMenuItem(m *MenuItem) *OrderItemUpdateOne {
	return oiuo.SetMenuItemID(m.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearOrder() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrder()
	return oiuo
}

// ClearMenuItem clears the "menu_item" edge to the MenuItem entity.
func (oiuo *OrderItemUpdateOne) ClearMenuItem() *OrderItemUpdateOne {
	oiuo.mutation.ClearMenuItem()
	return oiuo
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiuo *OrderItemUpdateOne) check() error {
	if v, ok := oiuo.mutation.Quantity(); ok {
		if err := orderitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "OrderItem.quantity": %w`, err)}
		}
	}
	if oiuo.mutation.OrderCleared() && len(oiuo.mutation.OrderIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.order"`)
	}
	if oiuo.mutation.MenuItemCleared() && len(oiuo.mutation.MenuItemIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.menu_item"`)
	}
	return nil
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	if err := oiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.SpecialInstructions(); ok {
		_spec.SetField(orderitem.FieldSpecialInstructions, field.TypeString, value)
	}
	if oiuo.mutation.SpecialInstructionsCleared() {
		_spec.ClearField(orderitem.FieldSpecialInstructions, field.TypeString)
	}
	if oiuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiuo.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.MenuItemTable,
			Columns: []string{orderitem.MenuItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.MenuItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.MenuItemTable,
			Columns: []string{orderitem.MenuItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}
