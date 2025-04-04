// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"project/ent/kitchen"
	"project/ent/predicate"
	"project/ent/shift"
	"project/ent/staff"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StaffQuery is the builder for querying Staff entities.
type StaffQuery struct {
	config
	ctx         *QueryContext
	order       []staff.OrderOption
	inters      []Interceptor
	predicates  []predicate.Staff
	withKitchen *KitchenQuery
	withShifts  *ShiftQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StaffQuery builder.
func (sq *StaffQuery) Where(ps ...predicate.Staff) *StaffQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *StaffQuery) Limit(limit int) *StaffQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *StaffQuery) Offset(offset int) *StaffQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StaffQuery) Unique(unique bool) *StaffQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *StaffQuery) Order(o ...staff.OrderOption) *StaffQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryKitchen chains the current query on the "kitchen" edge.
func (sq *StaffQuery) QueryKitchen() *KitchenQuery {
	query := (&KitchenClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staff.Table, staff.FieldID, selector),
			sqlgraph.To(kitchen.Table, kitchen.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, staff.KitchenTable, staff.KitchenColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShifts chains the current query on the "shifts" edge.
func (sq *StaffQuery) QueryShifts() *ShiftQuery {
	query := (&ShiftClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staff.Table, staff.FieldID, selector),
			sqlgraph.To(shift.Table, shift.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, staff.ShiftsTable, staff.ShiftsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Staff entity from the query.
// Returns a *NotFoundError when no Staff was found.
func (sq *StaffQuery) First(ctx context.Context) (*Staff, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{staff.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StaffQuery) FirstX(ctx context.Context) *Staff {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Staff ID from the query.
// Returns a *NotFoundError when no Staff ID was found.
func (sq *StaffQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{staff.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StaffQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Staff entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Staff entity is found.
// Returns a *NotFoundError when no Staff entities are found.
func (sq *StaffQuery) Only(ctx context.Context) (*Staff, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{staff.Label}
	default:
		return nil, &NotSingularError{staff.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StaffQuery) OnlyX(ctx context.Context) *Staff {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Staff ID in the query.
// Returns a *NotSingularError when more than one Staff ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *StaffQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{staff.Label}
	default:
		err = &NotSingularError{staff.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StaffQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Staffs.
func (sq *StaffQuery) All(ctx context.Context) ([]*Staff, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Staff, *StaffQuery]()
	return withInterceptors[[]*Staff](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *StaffQuery) AllX(ctx context.Context) []*Staff {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Staff IDs.
func (sq *StaffQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(staff.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StaffQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StaffQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*StaffQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StaffQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StaffQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StaffQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StaffQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StaffQuery) Clone() *StaffQuery {
	if sq == nil {
		return nil
	}
	return &StaffQuery{
		config:      sq.config,
		ctx:         sq.ctx.Clone(),
		order:       append([]staff.OrderOption{}, sq.order...),
		inters:      append([]Interceptor{}, sq.inters...),
		predicates:  append([]predicate.Staff{}, sq.predicates...),
		withKitchen: sq.withKitchen.Clone(),
		withShifts:  sq.withShifts.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithKitchen tells the query-builder to eager-load the nodes that are connected to
// the "kitchen" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StaffQuery) WithKitchen(opts ...func(*KitchenQuery)) *StaffQuery {
	query := (&KitchenClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withKitchen = query
	return sq
}

// WithShifts tells the query-builder to eager-load the nodes that are connected to
// the "shifts" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StaffQuery) WithShifts(opts ...func(*ShiftQuery)) *StaffQuery {
	query := (&ShiftClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withShifts = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Staff.Query().
//		GroupBy(staff.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *StaffQuery) GroupBy(field string, fields ...string) *StaffGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StaffGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = staff.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Staff.Query().
//		Select(staff.FieldName).
//		Scan(ctx, &v)
func (sq *StaffQuery) Select(fields ...string) *StaffSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &StaffSelect{StaffQuery: sq}
	sbuild.label = staff.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StaffSelect configured with the given aggregations.
func (sq *StaffQuery) Aggregate(fns ...AggregateFunc) *StaffSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *StaffQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !staff.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StaffQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Staff, error) {
	var (
		nodes       = []*Staff{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withKitchen != nil,
			sq.withShifts != nil,
		}
	)
	if sq.withKitchen != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, staff.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Staff).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Staff{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withKitchen; query != nil {
		if err := sq.loadKitchen(ctx, query, nodes, nil,
			func(n *Staff, e *Kitchen) { n.Edges.Kitchen = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withShifts; query != nil {
		if err := sq.loadShifts(ctx, query, nodes,
			func(n *Staff) { n.Edges.Shifts = []*Shift{} },
			func(n *Staff, e *Shift) { n.Edges.Shifts = append(n.Edges.Shifts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *StaffQuery) loadKitchen(ctx context.Context, query *KitchenQuery, nodes []*Staff, init func(*Staff), assign func(*Staff, *Kitchen)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Staff)
	for i := range nodes {
		if nodes[i].kitchen_staff == nil {
			continue
		}
		fk := *nodes[i].kitchen_staff
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(kitchen.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "kitchen_staff" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *StaffQuery) loadShifts(ctx context.Context, query *ShiftQuery, nodes []*Staff, init func(*Staff), assign func(*Staff, *Shift)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Staff)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Shift(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(staff.ShiftsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.staff_shifts
		if fk == nil {
			return fmt.Errorf(`foreign-key "staff_shifts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "staff_shifts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *StaffQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StaffQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(staff.Table, staff.Columns, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, staff.FieldID)
		for i := range fields {
			if fields[i] != staff.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StaffQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(staff.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = staff.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StaffGroupBy is the group-by builder for Staff entities.
type StaffGroupBy struct {
	selector
	build *StaffQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StaffGroupBy) Aggregate(fns ...AggregateFunc) *StaffGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *StaffGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StaffQuery, *StaffGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *StaffGroupBy) sqlScan(ctx context.Context, root *StaffQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StaffSelect is the builder for selecting fields of Staff entities.
type StaffSelect struct {
	*StaffQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *StaffSelect) Aggregate(fns ...AggregateFunc) *StaffSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StaffSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StaffQuery, *StaffSelect](ctx, ss.StaffQuery, ss, ss.inters, v)
}

func (ss *StaffSelect) sqlScan(ctx context.Context, root *StaffQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
