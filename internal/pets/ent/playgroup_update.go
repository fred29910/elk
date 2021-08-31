// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/pets/ent/pet"
	"github.com/masseelch/elk/internal/pets/ent/playgroup"
	"github.com/masseelch/elk/internal/pets/ent/predicate"
)

// PlayGroupUpdate is the builder for updating PlayGroup entities.
type PlayGroupUpdate struct {
	config
	hooks    []Hook
	mutation *PlayGroupMutation
}

// Where appends a list predicates to the PlayGroupUpdate builder.
func (pgu *PlayGroupUpdate) Where(ps ...predicate.PlayGroup) *PlayGroupUpdate {
	pgu.mutation.Where(ps...)
	return pgu
}

// SetTitle sets the "title" field.
func (pgu *PlayGroupUpdate) SetTitle(s string) *PlayGroupUpdate {
	pgu.mutation.SetTitle(s)
	return pgu
}

// SetDescription sets the "description" field.
func (pgu *PlayGroupUpdate) SetDescription(s string) *PlayGroupUpdate {
	pgu.mutation.SetDescription(s)
	return pgu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pgu *PlayGroupUpdate) SetNillableDescription(s *string) *PlayGroupUpdate {
	if s != nil {
		pgu.SetDescription(*s)
	}
	return pgu
}

// ClearDescription clears the value of the "description" field.
func (pgu *PlayGroupUpdate) ClearDescription() *PlayGroupUpdate {
	pgu.mutation.ClearDescription()
	return pgu
}

// SetWeekday sets the "weekday" field.
func (pgu *PlayGroupUpdate) SetWeekday(pl playgroup.Weekday) *PlayGroupUpdate {
	pgu.mutation.SetWeekday(pl)
	return pgu
}

// AddParticipantIDs adds the "participants" edge to the Pet entity by IDs.
func (pgu *PlayGroupUpdate) AddParticipantIDs(ids ...int) *PlayGroupUpdate {
	pgu.mutation.AddParticipantIDs(ids...)
	return pgu
}

// AddParticipants adds the "participants" edges to the Pet entity.
func (pgu *PlayGroupUpdate) AddParticipants(p ...*Pet) *PlayGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.AddParticipantIDs(ids...)
}

// Mutation returns the PlayGroupMutation object of the builder.
func (pgu *PlayGroupUpdate) Mutation() *PlayGroupMutation {
	return pgu.mutation
}

// ClearParticipants clears all "participants" edges to the Pet entity.
func (pgu *PlayGroupUpdate) ClearParticipants() *PlayGroupUpdate {
	pgu.mutation.ClearParticipants()
	return pgu
}

// RemoveParticipantIDs removes the "participants" edge to Pet entities by IDs.
func (pgu *PlayGroupUpdate) RemoveParticipantIDs(ids ...int) *PlayGroupUpdate {
	pgu.mutation.RemoveParticipantIDs(ids...)
	return pgu
}

// RemoveParticipants removes "participants" edges to Pet entities.
func (pgu *PlayGroupUpdate) RemoveParticipants(p ...*Pet) *PlayGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.RemoveParticipantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pgu *PlayGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pgu.hooks) == 0 {
		if err = pgu.check(); err != nil {
			return 0, err
		}
		affected, err = pgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pgu.check(); err != nil {
				return 0, err
			}
			pgu.mutation = mutation
			affected, err = pgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pgu.hooks) - 1; i >= 0; i-- {
			if pgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pgu *PlayGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := pgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pgu *PlayGroupUpdate) Exec(ctx context.Context) error {
	_, err := pgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgu *PlayGroupUpdate) ExecX(ctx context.Context) {
	if err := pgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pgu *PlayGroupUpdate) check() error {
	if v, ok := pgu.mutation.Weekday(); ok {
		if err := playgroup.WeekdayValidator(v); err != nil {
			return &ValidationError{Name: "weekday", err: fmt.Errorf("ent: validator failed for field \"weekday\": %w", err)}
		}
	}
	return nil
}

func (pgu *PlayGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playgroup.Table,
			Columns: playgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playgroup.FieldID,
			},
		},
	}
	if ps := pgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pgu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playgroup.FieldTitle,
		})
	}
	if value, ok := pgu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playgroup.FieldDescription,
		})
	}
	if pgu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playgroup.FieldDescription,
		})
	}
	if value, ok := pgu.mutation.Weekday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: playgroup.FieldWeekday,
		})
	}
	if pgu.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.RemovedParticipantsIDs(); len(nodes) > 0 && !pgu.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PlayGroupUpdateOne is the builder for updating a single PlayGroup entity.
type PlayGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayGroupMutation
}

// SetTitle sets the "title" field.
func (pguo *PlayGroupUpdateOne) SetTitle(s string) *PlayGroupUpdateOne {
	pguo.mutation.SetTitle(s)
	return pguo
}

// SetDescription sets the "description" field.
func (pguo *PlayGroupUpdateOne) SetDescription(s string) *PlayGroupUpdateOne {
	pguo.mutation.SetDescription(s)
	return pguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pguo *PlayGroupUpdateOne) SetNillableDescription(s *string) *PlayGroupUpdateOne {
	if s != nil {
		pguo.SetDescription(*s)
	}
	return pguo
}

// ClearDescription clears the value of the "description" field.
func (pguo *PlayGroupUpdateOne) ClearDescription() *PlayGroupUpdateOne {
	pguo.mutation.ClearDescription()
	return pguo
}

// SetWeekday sets the "weekday" field.
func (pguo *PlayGroupUpdateOne) SetWeekday(pl playgroup.Weekday) *PlayGroupUpdateOne {
	pguo.mutation.SetWeekday(pl)
	return pguo
}

// AddParticipantIDs adds the "participants" edge to the Pet entity by IDs.
func (pguo *PlayGroupUpdateOne) AddParticipantIDs(ids ...int) *PlayGroupUpdateOne {
	pguo.mutation.AddParticipantIDs(ids...)
	return pguo
}

// AddParticipants adds the "participants" edges to the Pet entity.
func (pguo *PlayGroupUpdateOne) AddParticipants(p ...*Pet) *PlayGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.AddParticipantIDs(ids...)
}

// Mutation returns the PlayGroupMutation object of the builder.
func (pguo *PlayGroupUpdateOne) Mutation() *PlayGroupMutation {
	return pguo.mutation
}

// ClearParticipants clears all "participants" edges to the Pet entity.
func (pguo *PlayGroupUpdateOne) ClearParticipants() *PlayGroupUpdateOne {
	pguo.mutation.ClearParticipants()
	return pguo
}

// RemoveParticipantIDs removes the "participants" edge to Pet entities by IDs.
func (pguo *PlayGroupUpdateOne) RemoveParticipantIDs(ids ...int) *PlayGroupUpdateOne {
	pguo.mutation.RemoveParticipantIDs(ids...)
	return pguo
}

// RemoveParticipants removes "participants" edges to Pet entities.
func (pguo *PlayGroupUpdateOne) RemoveParticipants(p ...*Pet) *PlayGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.RemoveParticipantIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pguo *PlayGroupUpdateOne) Select(field string, fields ...string) *PlayGroupUpdateOne {
	pguo.fields = append([]string{field}, fields...)
	return pguo
}

// Save executes the query and returns the updated PlayGroup entity.
func (pguo *PlayGroupUpdateOne) Save(ctx context.Context) (*PlayGroup, error) {
	var (
		err  error
		node *PlayGroup
	)
	if len(pguo.hooks) == 0 {
		if err = pguo.check(); err != nil {
			return nil, err
		}
		node, err = pguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pguo.check(); err != nil {
				return nil, err
			}
			pguo.mutation = mutation
			node, err = pguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pguo.hooks) - 1; i >= 0; i-- {
			if pguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pguo *PlayGroupUpdateOne) SaveX(ctx context.Context) *PlayGroup {
	node, err := pguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pguo *PlayGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := pguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pguo *PlayGroupUpdateOne) ExecX(ctx context.Context) {
	if err := pguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pguo *PlayGroupUpdateOne) check() error {
	if v, ok := pguo.mutation.Weekday(); ok {
		if err := playgroup.WeekdayValidator(v); err != nil {
			return &ValidationError{Name: "weekday", err: fmt.Errorf("ent: validator failed for field \"weekday\": %w", err)}
		}
	}
	return nil
}

func (pguo *PlayGroupUpdateOne) sqlSave(ctx context.Context) (_node *PlayGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playgroup.Table,
			Columns: playgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playgroup.FieldID,
			},
		},
	}
	id, ok := pguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PlayGroup.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playgroup.FieldID)
		for _, f := range fields {
			if !playgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != playgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pguo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playgroup.FieldTitle,
		})
	}
	if value, ok := pguo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playgroup.FieldDescription,
		})
	}
	if pguo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playgroup.FieldDescription,
		})
	}
	if value, ok := pguo.mutation.Weekday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: playgroup.FieldWeekday,
		})
	}
	if pguo.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.RemovedParticipantsIDs(); len(nodes) > 0 && !pguo.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   playgroup.ParticipantsTable,
			Columns: playgroup.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlayGroup{config: pguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}