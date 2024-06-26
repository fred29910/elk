// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/masseelch/elk/internal/pets/ent/pet"
	"github.com/masseelch/elk/internal/pets/ent/predicate"
	"github.com/masseelch/elk/internal/pets/ent/toy"
)

// ToyQuery is the builder for querying Toy entities.
type ToyQuery struct {
	config
	ctx        *QueryContext
	order      []toy.OrderOption
	inters     []Interceptor
	predicates []predicate.Toy
	withOwner  *PetQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ToyQuery builder.
func (tq *ToyQuery) Where(ps ...predicate.Toy) *ToyQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *ToyQuery) Limit(limit int) *ToyQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *ToyQuery) Offset(offset int) *ToyQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *ToyQuery) Unique(unique bool) *ToyQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *ToyQuery) Order(o ...toy.OrderOption) *ToyQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryOwner chains the current query on the "owner" edge.
func (tq *ToyQuery) QueryOwner() *PetQuery {
	query := (&PetClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(toy.Table, toy.FieldID, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, toy.OwnerTable, toy.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Toy entity from the query.
// Returns a *NotFoundError when no Toy was found.
func (tq *ToyQuery) First(ctx context.Context) (*Toy, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{toy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *ToyQuery) FirstX(ctx context.Context) *Toy {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Toy ID from the query.
// Returns a *NotFoundError when no Toy ID was found.
func (tq *ToyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{toy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *ToyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Toy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Toy entity is found.
// Returns a *NotFoundError when no Toy entities are found.
func (tq *ToyQuery) Only(ctx context.Context) (*Toy, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{toy.Label}
	default:
		return nil, &NotSingularError{toy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *ToyQuery) OnlyX(ctx context.Context) *Toy {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Toy ID in the query.
// Returns a *NotSingularError when more than one Toy ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *ToyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{toy.Label}
	default:
		err = &NotSingularError{toy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *ToyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Toys.
func (tq *ToyQuery) All(ctx context.Context) ([]*Toy, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Toy, *ToyQuery]()
	return withInterceptors[[]*Toy](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *ToyQuery) AllX(ctx context.Context) []*Toy {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Toy IDs.
func (tq *ToyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(toy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *ToyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *ToyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*ToyQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *ToyQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *ToyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *ToyQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ToyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *ToyQuery) Clone() *ToyQuery {
	if tq == nil {
		return nil
	}
	return &ToyQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]toy.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Toy{}, tq.predicates...),
		withOwner:  tq.withOwner.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ToyQuery) WithOwner(opts ...func(*PetQuery)) *ToyQuery {
	query := (&PetClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withOwner = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Color toy.Color `json:"color,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Toy.Query().
//		GroupBy(toy.FieldColor).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *ToyQuery) GroupBy(field string, fields ...string) *ToyGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ToyGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = toy.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Color toy.Color `json:"color,omitempty"`
//	}
//
//	client.Toy.Query().
//		Select(toy.FieldColor).
//		Scan(ctx, &v)
func (tq *ToyQuery) Select(fields ...string) *ToySelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &ToySelect{ToyQuery: tq}
	sbuild.label = toy.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ToySelect configured with the given aggregations.
func (tq *ToyQuery) Aggregate(fns ...AggregateFunc) *ToySelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *ToyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !toy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *ToyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Toy, error) {
	var (
		nodes       = []*Toy{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withOwner != nil,
		}
	)
	if tq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, toy.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Toy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Toy{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withOwner; query != nil {
		if err := tq.loadOwner(ctx, query, nodes, nil,
			func(n *Toy, e *Pet) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *ToyQuery) loadOwner(ctx context.Context, query *PetQuery, nodes []*Toy, init func(*Toy), assign func(*Toy, *Pet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Toy)
	for i := range nodes {
		if nodes[i].pet_toys == nil {
			continue
		}
		fk := *nodes[i].pet_toys
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(pet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pet_toys" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *ToyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *ToyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(toy.Table, toy.Columns, sqlgraph.NewFieldSpec(toy.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, toy.FieldID)
		for i := range fields {
			if fields[i] != toy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *ToyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(toy.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = toy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ToyGroupBy is the group-by builder for Toy entities.
type ToyGroupBy struct {
	selector
	build *ToyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *ToyGroupBy) Aggregate(fns ...AggregateFunc) *ToyGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *ToyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ToyQuery, *ToyGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *ToyGroupBy) sqlScan(ctx context.Context, root *ToyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ToySelect is the builder for selecting fields of Toy entities.
type ToySelect struct {
	*ToyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *ToySelect) Aggregate(fns ...AggregateFunc) *ToySelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *ToySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ToyQuery, *ToySelect](ctx, ts.ToyQuery, ts, ts.inters, v)
}

func (ts *ToySelect) sqlScan(ctx context.Context, root *ToyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
