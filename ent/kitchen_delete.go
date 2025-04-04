// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"project/ent/kitchen"
	"project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KitchenDelete is the builder for deleting a Kitchen entity.
type KitchenDelete struct {
	config
	hooks    []Hook
	mutation *KitchenMutation
}

// Where appends a list predicates to the KitchenDelete builder.
func (kd *KitchenDelete) Where(ps ...predicate.Kitchen) *KitchenDelete {
	kd.mutation.Where(ps...)
	return kd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kd *KitchenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, kd.sqlExec, kd.mutation, kd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (kd *KitchenDelete) ExecX(ctx context.Context) int {
	n, err := kd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kd *KitchenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(kitchen.Table, sqlgraph.NewFieldSpec(kitchen.FieldID, field.TypeInt))
	if ps := kd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, kd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	kd.mutation.done = true
	return affected, err
}

// KitchenDeleteOne is the builder for deleting a single Kitchen entity.
type KitchenDeleteOne struct {
	kd *KitchenDelete
}

// Where appends a list predicates to the KitchenDelete builder.
func (kdo *KitchenDeleteOne) Where(ps ...predicate.Kitchen) *KitchenDeleteOne {
	kdo.kd.mutation.Where(ps...)
	return kdo
}

// Exec executes the deletion query.
func (kdo *KitchenDeleteOne) Exec(ctx context.Context) error {
	n, err := kdo.kd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{kitchen.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kdo *KitchenDeleteOne) ExecX(ctx context.Context) {
	if err := kdo.Exec(ctx); err != nil {
		panic(err)
	}
}
