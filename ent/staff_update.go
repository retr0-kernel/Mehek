// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project/ent/kitchen"
	"project/ent/predicate"
	"project/ent/shift"
	"project/ent/staff"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StaffUpdate is the builder for updating Staff entities.
type StaffUpdate struct {
	config
	hooks    []Hook
	mutation *StaffMutation
}

// Where appends a list predicates to the StaffUpdate builder.
func (su *StaffUpdate) Where(ps ...predicate.Staff) *StaffUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StaffUpdate) SetName(s string) *StaffUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *StaffUpdate) SetNillableName(s *string) *StaffUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetRole sets the "role" field.
func (su *StaffUpdate) SetRole(s string) *StaffUpdate {
	su.mutation.SetRole(s)
	return su
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (su *StaffUpdate) SetNillableRole(s *string) *StaffUpdate {
	if s != nil {
		su.SetRole(*s)
	}
	return su
}

// SetSkills sets the "skills" field.
func (su *StaffUpdate) SetSkills(s string) *StaffUpdate {
	su.mutation.SetSkills(s)
	return su
}

// SetNillableSkills sets the "skills" field if the given value is not nil.
func (su *StaffUpdate) SetNillableSkills(s *string) *StaffUpdate {
	if s != nil {
		su.SetSkills(*s)
	}
	return su
}

// SetAvailability sets the "availability" field.
func (su *StaffUpdate) SetAvailability(m map[string][]string) *StaffUpdate {
	su.mutation.SetAvailability(m)
	return su
}

// SetKitchenID sets the "kitchen" edge to the Kitchen entity by ID.
func (su *StaffUpdate) SetKitchenID(id int) *StaffUpdate {
	su.mutation.SetKitchenID(id)
	return su
}

// SetKitchen sets the "kitchen" edge to the Kitchen entity.
func (su *StaffUpdate) SetKitchen(k *Kitchen) *StaffUpdate {
	return su.SetKitchenID(k.ID)
}

// AddShiftIDs adds the "shifts" edge to the Shift entity by IDs.
func (su *StaffUpdate) AddShiftIDs(ids ...int) *StaffUpdate {
	su.mutation.AddShiftIDs(ids...)
	return su
}

// AddShifts adds the "shifts" edges to the Shift entity.
func (su *StaffUpdate) AddShifts(s ...*Shift) *StaffUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddShiftIDs(ids...)
}

// Mutation returns the StaffMutation object of the builder.
func (su *StaffUpdate) Mutation() *StaffMutation {
	return su.mutation
}

// ClearKitchen clears the "kitchen" edge to the Kitchen entity.
func (su *StaffUpdate) ClearKitchen() *StaffUpdate {
	su.mutation.ClearKitchen()
	return su
}

// ClearShifts clears all "shifts" edges to the Shift entity.
func (su *StaffUpdate) ClearShifts() *StaffUpdate {
	su.mutation.ClearShifts()
	return su
}

// RemoveShiftIDs removes the "shifts" edge to Shift entities by IDs.
func (su *StaffUpdate) RemoveShiftIDs(ids ...int) *StaffUpdate {
	su.mutation.RemoveShiftIDs(ids...)
	return su
}

// RemoveShifts removes "shifts" edges to Shift entities.
func (su *StaffUpdate) RemoveShifts(s ...*Shift) *StaffUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveShiftIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StaffUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StaffUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StaffUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StaffUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StaffUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := staff.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Staff.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Role(); ok {
		if err := staff.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Staff.role": %w`, err)}
		}
	}
	if su.mutation.KitchenCleared() && len(su.mutation.KitchenIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Staff.kitchen"`)
	}
	return nil
}

func (su *StaffUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(staff.Table, staff.Columns, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(staff.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Role(); ok {
		_spec.SetField(staff.FieldRole, field.TypeString, value)
	}
	if value, ok := su.mutation.Skills(); ok {
		_spec.SetField(staff.FieldSkills, field.TypeString, value)
	}
	if value, ok := su.mutation.Availability(); ok {
		_spec.SetField(staff.FieldAvailability, field.TypeJSON, value)
	}
	if su.mutation.KitchenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   staff.KitchenTable,
			Columns: []string{staff.KitchenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kitchen.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.KitchenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   staff.KitchenTable,
			Columns: []string{staff.KitchenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kitchen.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedShiftsIDs(); len(nodes) > 0 && !su.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{staff.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StaffUpdateOne is the builder for updating a single Staff entity.
type StaffUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StaffMutation
}

// SetName sets the "name" field.
func (suo *StaffUpdateOne) SetName(s string) *StaffUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *StaffUpdateOne) SetNillableName(s *string) *StaffUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetRole sets the "role" field.
func (suo *StaffUpdateOne) SetRole(s string) *StaffUpdateOne {
	suo.mutation.SetRole(s)
	return suo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (suo *StaffUpdateOne) SetNillableRole(s *string) *StaffUpdateOne {
	if s != nil {
		suo.SetRole(*s)
	}
	return suo
}

// SetSkills sets the "skills" field.
func (suo *StaffUpdateOne) SetSkills(s string) *StaffUpdateOne {
	suo.mutation.SetSkills(s)
	return suo
}

// SetNillableSkills sets the "skills" field if the given value is not nil.
func (suo *StaffUpdateOne) SetNillableSkills(s *string) *StaffUpdateOne {
	if s != nil {
		suo.SetSkills(*s)
	}
	return suo
}

// SetAvailability sets the "availability" field.
func (suo *StaffUpdateOne) SetAvailability(m map[string][]string) *StaffUpdateOne {
	suo.mutation.SetAvailability(m)
	return suo
}

// SetKitchenID sets the "kitchen" edge to the Kitchen entity by ID.
func (suo *StaffUpdateOne) SetKitchenID(id int) *StaffUpdateOne {
	suo.mutation.SetKitchenID(id)
	return suo
}

// SetKitchen sets the "kitchen" edge to the Kitchen entity.
func (suo *StaffUpdateOne) SetKitchen(k *Kitchen) *StaffUpdateOne {
	return suo.SetKitchenID(k.ID)
}

// AddShiftIDs adds the "shifts" edge to the Shift entity by IDs.
func (suo *StaffUpdateOne) AddShiftIDs(ids ...int) *StaffUpdateOne {
	suo.mutation.AddShiftIDs(ids...)
	return suo
}

// AddShifts adds the "shifts" edges to the Shift entity.
func (suo *StaffUpdateOne) AddShifts(s ...*Shift) *StaffUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddShiftIDs(ids...)
}

// Mutation returns the StaffMutation object of the builder.
func (suo *StaffUpdateOne) Mutation() *StaffMutation {
	return suo.mutation
}

// ClearKitchen clears the "kitchen" edge to the Kitchen entity.
func (suo *StaffUpdateOne) ClearKitchen() *StaffUpdateOne {
	suo.mutation.ClearKitchen()
	return suo
}

// ClearShifts clears all "shifts" edges to the Shift entity.
func (suo *StaffUpdateOne) ClearShifts() *StaffUpdateOne {
	suo.mutation.ClearShifts()
	return suo
}

// RemoveShiftIDs removes the "shifts" edge to Shift entities by IDs.
func (suo *StaffUpdateOne) RemoveShiftIDs(ids ...int) *StaffUpdateOne {
	suo.mutation.RemoveShiftIDs(ids...)
	return suo
}

// RemoveShifts removes "shifts" edges to Shift entities.
func (suo *StaffUpdateOne) RemoveShifts(s ...*Shift) *StaffUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveShiftIDs(ids...)
}

// Where appends a list predicates to the StaffUpdate builder.
func (suo *StaffUpdateOne) Where(ps ...predicate.Staff) *StaffUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StaffUpdateOne) Select(field string, fields ...string) *StaffUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Staff entity.
func (suo *StaffUpdateOne) Save(ctx context.Context) (*Staff, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StaffUpdateOne) SaveX(ctx context.Context) *Staff {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StaffUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StaffUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StaffUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := staff.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Staff.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Role(); ok {
		if err := staff.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Staff.role": %w`, err)}
		}
	}
	if suo.mutation.KitchenCleared() && len(suo.mutation.KitchenIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Staff.kitchen"`)
	}
	return nil
}

func (suo *StaffUpdateOne) sqlSave(ctx context.Context) (_node *Staff, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(staff.Table, staff.Columns, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Staff.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, staff.FieldID)
		for _, f := range fields {
			if !staff.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != staff.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(staff.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Role(); ok {
		_spec.SetField(staff.FieldRole, field.TypeString, value)
	}
	if value, ok := suo.mutation.Skills(); ok {
		_spec.SetField(staff.FieldSkills, field.TypeString, value)
	}
	if value, ok := suo.mutation.Availability(); ok {
		_spec.SetField(staff.FieldAvailability, field.TypeJSON, value)
	}
	if suo.mutation.KitchenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   staff.KitchenTable,
			Columns: []string{staff.KitchenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kitchen.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.KitchenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   staff.KitchenTable,
			Columns: []string{staff.KitchenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kitchen.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedShiftsIDs(); len(nodes) > 0 && !suo.mutation.ShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.ShiftsTable,
			Columns: []string{staff.ShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Staff{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{staff.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
