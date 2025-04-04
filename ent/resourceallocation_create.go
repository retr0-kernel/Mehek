// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project/ent/kitchenresource"
	"project/ent/order"
	"project/ent/resourceallocation"
	"project/ent/shift"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceAllocationCreate is the builder for creating a ResourceAllocation entity.
type ResourceAllocationCreate struct {
	config
	mutation *ResourceAllocationMutation
	hooks    []Hook
}

// SetStartTime sets the "start_time" field.
func (rac *ResourceAllocationCreate) SetStartTime(t time.Time) *ResourceAllocationCreate {
	rac.mutation.SetStartTime(t)
	return rac
}

// SetEndTime sets the "end_time" field.
func (rac *ResourceAllocationCreate) SetEndTime(t time.Time) *ResourceAllocationCreate {
	rac.mutation.SetEndTime(t)
	return rac
}

// SetStatus sets the "status" field.
func (rac *ResourceAllocationCreate) SetStatus(s string) *ResourceAllocationCreate {
	rac.mutation.SetStatus(s)
	return rac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rac *ResourceAllocationCreate) SetNillableStatus(s *string) *ResourceAllocationCreate {
	if s != nil {
		rac.SetStatus(*s)
	}
	return rac
}

// SetResourceID sets the "resource" edge to the KitchenResource entity by ID.
func (rac *ResourceAllocationCreate) SetResourceID(id int) *ResourceAllocationCreate {
	rac.mutation.SetResourceID(id)
	return rac
}

// SetResource sets the "resource" edge to the KitchenResource entity.
func (rac *ResourceAllocationCreate) SetResource(k *KitchenResource) *ResourceAllocationCreate {
	return rac.SetResourceID(k.ID)
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (rac *ResourceAllocationCreate) SetOrderID(id int) *ResourceAllocationCreate {
	rac.mutation.SetOrderID(id)
	return rac
}

// SetOrder sets the "order" edge to the Order entity.
func (rac *ResourceAllocationCreate) SetOrder(o *Order) *ResourceAllocationCreate {
	return rac.SetOrderID(o.ID)
}

// SetShiftID sets the "shift" edge to the Shift entity by ID.
func (rac *ResourceAllocationCreate) SetShiftID(id int) *ResourceAllocationCreate {
	rac.mutation.SetShiftID(id)
	return rac
}

// SetShift sets the "shift" edge to the Shift entity.
func (rac *ResourceAllocationCreate) SetShift(s *Shift) *ResourceAllocationCreate {
	return rac.SetShiftID(s.ID)
}

// Mutation returns the ResourceAllocationMutation object of the builder.
func (rac *ResourceAllocationCreate) Mutation() *ResourceAllocationMutation {
	return rac.mutation
}

// Save creates the ResourceAllocation in the database.
func (rac *ResourceAllocationCreate) Save(ctx context.Context) (*ResourceAllocation, error) {
	rac.defaults()
	return withHooks(ctx, rac.sqlSave, rac.mutation, rac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rac *ResourceAllocationCreate) SaveX(ctx context.Context) *ResourceAllocation {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *ResourceAllocationCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *ResourceAllocationCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rac *ResourceAllocationCreate) defaults() {
	if _, ok := rac.mutation.Status(); !ok {
		v := resourceallocation.DefaultStatus
		rac.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *ResourceAllocationCreate) check() error {
	if _, ok := rac.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "ResourceAllocation.start_time"`)}
	}
	if _, ok := rac.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "ResourceAllocation.end_time"`)}
	}
	if _, ok := rac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ResourceAllocation.status"`)}
	}
	if len(rac.mutation.ResourceIDs()) == 0 {
		return &ValidationError{Name: "resource", err: errors.New(`ent: missing required edge "ResourceAllocation.resource"`)}
	}
	if len(rac.mutation.OrderIDs()) == 0 {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "ResourceAllocation.order"`)}
	}
	if len(rac.mutation.ShiftIDs()) == 0 {
		return &ValidationError{Name: "shift", err: errors.New(`ent: missing required edge "ResourceAllocation.shift"`)}
	}
	return nil
}

func (rac *ResourceAllocationCreate) sqlSave(ctx context.Context) (*ResourceAllocation, error) {
	if err := rac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rac.mutation.id = &_node.ID
	rac.mutation.done = true
	return _node, nil
}

func (rac *ResourceAllocationCreate) createSpec() (*ResourceAllocation, *sqlgraph.CreateSpec) {
	var (
		_node = &ResourceAllocation{config: rac.config}
		_spec = sqlgraph.NewCreateSpec(resourceallocation.Table, sqlgraph.NewFieldSpec(resourceallocation.FieldID, field.TypeInt))
	)
	if value, ok := rac.mutation.StartTime(); ok {
		_spec.SetField(resourceallocation.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := rac.mutation.EndTime(); ok {
		_spec.SetField(resourceallocation.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := rac.mutation.Status(); ok {
		_spec.SetField(resourceallocation.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if nodes := rac.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resourceallocation.ResourceTable,
			Columns: []string{resourceallocation.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kitchenresource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.kitchen_resource_allocations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rac.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourceallocation.OrderTable,
			Columns: []string{resourceallocation.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.resource_allocation_order = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rac.mutation.ShiftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resourceallocation.ShiftTable,
			Columns: []string{resourceallocation.ShiftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.shift_allocations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResourceAllocationCreateBulk is the builder for creating many ResourceAllocation entities in bulk.
type ResourceAllocationCreateBulk struct {
	config
	err      error
	builders []*ResourceAllocationCreate
}

// Save creates the ResourceAllocation entities in the database.
func (racb *ResourceAllocationCreateBulk) Save(ctx context.Context) ([]*ResourceAllocation, error) {
	if racb.err != nil {
		return nil, racb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*ResourceAllocation, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceAllocationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *ResourceAllocationCreateBulk) SaveX(ctx context.Context) []*ResourceAllocation {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *ResourceAllocationCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *ResourceAllocationCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}
