// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/fridge/ent/compartment"
	"github.com/masseelch/elk/internal/fridge/ent/fridge"
	"github.com/masseelch/elk/internal/fridge/ent/item"
	"github.com/masseelch/elk/internal/fridge/ent/predicate"
)

// CompartmentQuery is the builder for querying Compartment entities.
type CompartmentQuery struct {
	config
	ctx          *QueryContext
	order        []compartment.OrderOption
	inters       []Interceptor
	predicates   []predicate.Compartment
	withFridge   *FridgeQuery
	withContents *ItemQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompartmentQuery builder.
func (cq *CompartmentQuery) Where(ps ...predicate.Compartment) *CompartmentQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CompartmentQuery) Limit(limit int) *CompartmentQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CompartmentQuery) Offset(offset int) *CompartmentQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CompartmentQuery) Unique(unique bool) *CompartmentQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CompartmentQuery) Order(o ...compartment.OrderOption) *CompartmentQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryFridge chains the current query on the "fridge" edge.
func (cq *CompartmentQuery) QueryFridge() *FridgeQuery {
	query := (&FridgeClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(compartment.Table, compartment.FieldID, selector),
			sqlgraph.To(fridge.Table, fridge.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, compartment.FridgeTable, compartment.FridgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContents chains the current query on the "contents" edge.
func (cq *CompartmentQuery) QueryContents() *ItemQuery {
	query := (&ItemClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(compartment.Table, compartment.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, compartment.ContentsTable, compartment.ContentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Compartment entity from the query.
// Returns a *NotFoundError when no Compartment was found.
func (cq *CompartmentQuery) First(ctx context.Context) (*Compartment, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{compartment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CompartmentQuery) FirstX(ctx context.Context) *Compartment {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Compartment ID from the query.
// Returns a *NotFoundError when no Compartment ID was found.
func (cq *CompartmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{compartment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CompartmentQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Compartment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Compartment entity is found.
// Returns a *NotFoundError when no Compartment entities are found.
func (cq *CompartmentQuery) Only(ctx context.Context) (*Compartment, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{compartment.Label}
	default:
		return nil, &NotSingularError{compartment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CompartmentQuery) OnlyX(ctx context.Context) *Compartment {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Compartment ID in the query.
// Returns a *NotSingularError when more than one Compartment ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CompartmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{compartment.Label}
	default:
		err = &NotSingularError{compartment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CompartmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Compartments.
func (cq *CompartmentQuery) All(ctx context.Context) ([]*Compartment, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Compartment, *CompartmentQuery]()
	return withInterceptors[[]*Compartment](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CompartmentQuery) AllX(ctx context.Context) []*Compartment {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Compartment IDs.
func (cq *CompartmentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(compartment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CompartmentQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CompartmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CompartmentQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CompartmentQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CompartmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CompartmentQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompartmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CompartmentQuery) Clone() *CompartmentQuery {
	if cq == nil {
		return nil
	}
	return &CompartmentQuery{
		config:       cq.config,
		ctx:          cq.ctx.Clone(),
		order:        append([]compartment.OrderOption{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Compartment{}, cq.predicates...),
		withFridge:   cq.withFridge.Clone(),
		withContents: cq.withContents.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithFridge tells the query-builder to eager-load the nodes that are connected to
// the "fridge" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompartmentQuery) WithFridge(opts ...func(*FridgeQuery)) *CompartmentQuery {
	query := (&FridgeClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withFridge = query
	return cq
}

// WithContents tells the query-builder to eager-load the nodes that are connected to
// the "contents" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompartmentQuery) WithContents(opts ...func(*ItemQuery)) *CompartmentQuery {
	query := (&ItemClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withContents = query
	return cq
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
//	client.Compartment.Query().
//		GroupBy(compartment.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CompartmentQuery) GroupBy(field string, fields ...string) *CompartmentGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompartmentGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = compartment.Label
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
//	client.Compartment.Query().
//		Select(compartment.FieldName).
//		Scan(ctx, &v)
func (cq *CompartmentQuery) Select(fields ...string) *CompartmentSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CompartmentSelect{CompartmentQuery: cq}
	sbuild.label = compartment.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompartmentSelect configured with the given aggregations.
func (cq *CompartmentQuery) Aggregate(fns ...AggregateFunc) *CompartmentSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CompartmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !compartment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CompartmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Compartment, error) {
	var (
		nodes       = []*Compartment{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withFridge != nil,
			cq.withContents != nil,
		}
	)
	if cq.withFridge != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, compartment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Compartment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Compartment{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withFridge; query != nil {
		if err := cq.loadFridge(ctx, query, nodes, nil,
			func(n *Compartment, e *Fridge) { n.Edges.Fridge = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withContents; query != nil {
		if err := cq.loadContents(ctx, query, nodes,
			func(n *Compartment) { n.Edges.Contents = []*Item{} },
			func(n *Compartment, e *Item) { n.Edges.Contents = append(n.Edges.Contents, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CompartmentQuery) loadFridge(ctx context.Context, query *FridgeQuery, nodes []*Compartment, init func(*Compartment), assign func(*Compartment, *Fridge)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Compartment)
	for i := range nodes {
		if nodes[i].fridge_compartments == nil {
			continue
		}
		fk := *nodes[i].fridge_compartments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(fridge.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "fridge_compartments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CompartmentQuery) loadContents(ctx context.Context, query *ItemQuery, nodes []*Compartment, init func(*Compartment), assign func(*Compartment, *Item)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Compartment)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Item(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(compartment.ContentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.compartment_contents
		if fk == nil {
			return fmt.Errorf(`foreign-key "compartment_contents" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "compartment_contents" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CompartmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CompartmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(compartment.Table, compartment.Columns, sqlgraph.NewFieldSpec(compartment.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, compartment.FieldID)
		for i := range fields {
			if fields[i] != compartment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CompartmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(compartment.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = compartment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompartmentGroupBy is the group-by builder for Compartment entities.
type CompartmentGroupBy struct {
	selector
	build *CompartmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CompartmentGroupBy) Aggregate(fns ...AggregateFunc) *CompartmentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CompartmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompartmentQuery, *CompartmentGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CompartmentGroupBy) sqlScan(ctx context.Context, root *CompartmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompartmentSelect is the builder for selecting fields of Compartment entities.
type CompartmentSelect struct {
	*CompartmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CompartmentSelect) Aggregate(fns ...AggregateFunc) *CompartmentSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CompartmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompartmentQuery, *CompartmentSelect](ctx, cs.CompartmentQuery, cs, cs.inters, v)
}

func (cs *CompartmentSelect) sqlScan(ctx context.Context, root *CompartmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
