// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/integration/pets/ent/pet"
	"github.com/masseelch/elk/internal/integration/pets/ent/playgroup"
	"github.com/masseelch/elk/internal/integration/pets/ent/predicate"
)

// PlayGroupQuery is the builder for querying PlayGroup entities.
type PlayGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PlayGroup
	// eager-loading edges.
	withParticipants *PetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlayGroupQuery builder.
func (pgq *PlayGroupQuery) Where(ps ...predicate.PlayGroup) *PlayGroupQuery {
	pgq.predicates = append(pgq.predicates, ps...)
	return pgq
}

// Limit adds a limit step to the query.
func (pgq *PlayGroupQuery) Limit(limit int) *PlayGroupQuery {
	pgq.limit = &limit
	return pgq
}

// Offset adds an offset step to the query.
func (pgq *PlayGroupQuery) Offset(offset int) *PlayGroupQuery {
	pgq.offset = &offset
	return pgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pgq *PlayGroupQuery) Unique(unique bool) *PlayGroupQuery {
	pgq.unique = &unique
	return pgq
}

// Order adds an order step to the query.
func (pgq *PlayGroupQuery) Order(o ...OrderFunc) *PlayGroupQuery {
	pgq.order = append(pgq.order, o...)
	return pgq
}

// QueryParticipants chains the current query on the "participants" edge.
func (pgq *PlayGroupQuery) QueryParticipants() *PetQuery {
	query := &PetQuery{config: pgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playgroup.Table, playgroup.FieldID, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, playgroup.ParticipantsTable, playgroup.ParticipantsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlayGroup entity from the query.
// Returns a *NotFoundError when no PlayGroup was found.
func (pgq *PlayGroupQuery) First(ctx context.Context) (*PlayGroup, error) {
	nodes, err := pgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pgq *PlayGroupQuery) FirstX(ctx context.Context) *PlayGroup {
	node, err := pgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlayGroup ID from the query.
// Returns a *NotFoundError when no PlayGroup ID was found.
func (pgq *PlayGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pgq *PlayGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := pgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlayGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PlayGroup entity is not found.
// Returns a *NotFoundError when no PlayGroup entities are found.
func (pgq *PlayGroupQuery) Only(ctx context.Context) (*PlayGroup, error) {
	nodes, err := pgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playgroup.Label}
	default:
		return nil, &NotSingularError{playgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pgq *PlayGroupQuery) OnlyX(ctx context.Context) *PlayGroup {
	node, err := pgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlayGroup ID in the query.
// Returns a *NotSingularError when exactly one PlayGroup ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pgq *PlayGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = &NotSingularError{playgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pgq *PlayGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := pgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlayGroups.
func (pgq *PlayGroupQuery) All(ctx context.Context) ([]*PlayGroup, error) {
	if err := pgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pgq *PlayGroupQuery) AllX(ctx context.Context) []*PlayGroup {
	nodes, err := pgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlayGroup IDs.
func (pgq *PlayGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pgq.Select(playgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pgq *PlayGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := pgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pgq *PlayGroupQuery) Count(ctx context.Context) (int, error) {
	if err := pgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pgq *PlayGroupQuery) CountX(ctx context.Context) int {
	count, err := pgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pgq *PlayGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := pgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pgq *PlayGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := pgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlayGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pgq *PlayGroupQuery) Clone() *PlayGroupQuery {
	if pgq == nil {
		return nil
	}
	return &PlayGroupQuery{
		config:           pgq.config,
		limit:            pgq.limit,
		offset:           pgq.offset,
		order:            append([]OrderFunc{}, pgq.order...),
		predicates:       append([]predicate.PlayGroup{}, pgq.predicates...),
		withParticipants: pgq.withParticipants.Clone(),
		// clone intermediate query.
		sql:  pgq.sql.Clone(),
		path: pgq.path,
	}
}

// WithParticipants tells the query-builder to eager-load the nodes that are connected to
// the "participants" edge. The optional arguments are used to configure the query builder of the edge.
func (pgq *PlayGroupQuery) WithParticipants(opts ...func(*PetQuery)) *PlayGroupQuery {
	query := &PetQuery{config: pgq.config}
	for _, opt := range opts {
		opt(query)
	}
	pgq.withParticipants = query
	return pgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlayGroup.Query().
//		GroupBy(playgroup.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pgq *PlayGroupQuery) GroupBy(field string, fields ...string) *PlayGroupGroupBy {
	group := &PlayGroupGroupBy{config: pgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pgq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.PlayGroup.Query().
//		Select(playgroup.FieldTitle).
//		Scan(ctx, &v)
//
func (pgq *PlayGroupQuery) Select(fields ...string) *PlayGroupSelect {
	pgq.fields = append(pgq.fields, fields...)
	return &PlayGroupSelect{PlayGroupQuery: pgq}
}

func (pgq *PlayGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pgq.fields {
		if !playgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pgq.path != nil {
		prev, err := pgq.path(ctx)
		if err != nil {
			return err
		}
		pgq.sql = prev
	}
	return nil
}

func (pgq *PlayGroupQuery) sqlAll(ctx context.Context) ([]*PlayGroup, error) {
	var (
		nodes       = []*PlayGroup{}
		_spec       = pgq.querySpec()
		loadedTypes = [1]bool{
			pgq.withParticipants != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PlayGroup{config: pgq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pgq.withParticipants; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*PlayGroup, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Participants = []*Pet{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*PlayGroup)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   playgroup.ParticipantsTable,
				Columns: playgroup.ParticipantsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(playgroup.ParticipantsPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, pgq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "participants": %w`, err)
		}
		query.Where(pet.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "participants" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Participants = append(nodes[i].Edges.Participants, n)
			}
		}
	}

	return nodes, nil
}

func (pgq *PlayGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pgq.querySpec()
	return sqlgraph.CountNodes(ctx, pgq.driver, _spec)
}

func (pgq *PlayGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pgq *PlayGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playgroup.Table,
			Columns: playgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playgroup.FieldID,
			},
		},
		From:   pgq.sql,
		Unique: true,
	}
	if unique := pgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playgroup.FieldID)
		for i := range fields {
			if fields[i] != playgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pgq *PlayGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pgq.driver.Dialect())
	t1 := builder.Table(playgroup.Table)
	columns := pgq.fields
	if len(columns) == 0 {
		columns = playgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pgq.sql != nil {
		selector = pgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pgq.predicates {
		p(selector)
	}
	for _, p := range pgq.order {
		p(selector)
	}
	if offset := pgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlayGroupGroupBy is the group-by builder for PlayGroup entities.
type PlayGroupGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pggb *PlayGroupGroupBy) Aggregate(fns ...AggregateFunc) *PlayGroupGroupBy {
	pggb.fns = append(pggb.fns, fns...)
	return pggb
}

// Scan applies the group-by query and scans the result into the given value.
func (pggb *PlayGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pggb.path(ctx)
	if err != nil {
		return err
	}
	pggb.sql = query
	return pggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pggb.fields) > 1 {
		return nil, errors.New("ent: PlayGroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := pggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) StringX(ctx context.Context) string {
	v, err := pggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pggb.fields) > 1 {
		return nil, errors.New("ent: PlayGroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := pggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) IntX(ctx context.Context) int {
	v, err := pggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pggb.fields) > 1 {
		return nil, errors.New("ent: PlayGroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pggb.fields) > 1 {
		return nil, errors.New("ent: PlayGroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pggb *PlayGroupGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pggb *PlayGroupGroupBy) BoolX(ctx context.Context) bool {
	v, err := pggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pggb *PlayGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pggb.fields {
		if !playgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pggb *PlayGroupGroupBy) sqlQuery() *sql.Selector {
	selector := pggb.sql.Select()
	aggregation := make([]string, 0, len(pggb.fns))
	for _, fn := range pggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pggb.fields)+len(pggb.fns))
		for _, f := range pggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pggb.fields...)...)
}

// PlayGroupSelect is the builder for selecting fields of PlayGroup entities.
type PlayGroupSelect struct {
	*PlayGroupQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pgs *PlayGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pgs.prepareQuery(ctx); err != nil {
		return err
	}
	pgs.sql = pgs.PlayGroupQuery.sqlQuery(ctx)
	return pgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgs *PlayGroupSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pgs.fields) > 1 {
		return nil, errors.New("ent: PlayGroupSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgs *PlayGroupSelect) StringsX(ctx context.Context) []string {
	v, err := pgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgs *PlayGroupSelect) StringX(ctx context.Context) string {
	v, err := pgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pgs.fields) > 1 {
		return nil, errors.New("ent: PlayGroupSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgs *PlayGroupSelect) IntsX(ctx context.Context) []int {
	v, err := pgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgs *PlayGroupSelect) IntX(ctx context.Context) int {
	v, err := pgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgs.fields) > 1 {
		return nil, errors.New("ent: PlayGroupSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgs *PlayGroupSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgs *PlayGroupSelect) Float64X(ctx context.Context) float64 {
	v, err := pgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pgs.fields) > 1 {
		return nil, errors.New("ent: PlayGroupSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgs *PlayGroupSelect) BoolsX(ctx context.Context) []bool {
	v, err := pgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pgs *PlayGroupSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playgroup.Label}
	default:
		err = fmt.Errorf("ent: PlayGroupSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgs *PlayGroupSelect) BoolX(ctx context.Context) bool {
	v, err := pgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgs *PlayGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgs.sql.Query()
	if err := pgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}