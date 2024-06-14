// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/openapi/ent/owner"
	"github.com/masseelch/elk/internal/openapi/ent/pet"
	"github.com/masseelch/elk/internal/openapi/ent/predicate"
)

// OwnerUpdate is the builder for updating Owner entities.
type OwnerUpdate struct {
	config
	hooks    []Hook
	mutation *OwnerMutation
}

// Where appends a list predicates to the OwnerUpdate builder.
func (ou *OwnerUpdate) Where(ps ...predicate.Owner) *OwnerUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OwnerUpdate) SetName(s string) *OwnerUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OwnerUpdate) SetNillableName(s *string) *OwnerUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetAge sets the "age" field.
func (ou *OwnerUpdate) SetAge(i int) *OwnerUpdate {
	ou.mutation.ResetAge()
	ou.mutation.SetAge(i)
	return ou
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (ou *OwnerUpdate) SetNillableAge(i *int) *OwnerUpdate {
	if i != nil {
		ou.SetAge(*i)
	}
	return ou
}

// AddAge adds i to the "age" field.
func (ou *OwnerUpdate) AddAge(i int) *OwnerUpdate {
	ou.mutation.AddAge(i)
	return ou
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (ou *OwnerUpdate) AddPetIDs(ids ...int) *OwnerUpdate {
	ou.mutation.AddPetIDs(ids...)
	return ou
}

// AddPets adds the "pets" edges to the Pet entity.
func (ou *OwnerUpdate) AddPets(p ...*Pet) *OwnerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddPetIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ou *OwnerUpdate) Mutation() *OwnerMutation {
	return ou.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (ou *OwnerUpdate) ClearPets() *OwnerUpdate {
	ou.mutation.ClearPets()
	return ou
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (ou *OwnerUpdate) RemovePetIDs(ids ...int) *OwnerUpdate {
	ou.mutation.RemovePetIDs(ids...)
	return ou
}

// RemovePets removes "pets" edges to Pet entities.
func (ou *OwnerUpdate) RemovePets(p ...*Pet) *OwnerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemovePetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OwnerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OwnerUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OwnerUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OwnerUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OwnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(owner.Table, owner.Columns, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(owner.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Age(); ok {
		_spec.SetField(owner.FieldAge, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedAge(); ok {
		_spec.AddField(owner.FieldAge, field.TypeInt, value)
	}
	if ou.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedPetsIDs(); len(nodes) > 0 && !ou.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OwnerUpdateOne is the builder for updating a single Owner entity.
type OwnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OwnerMutation
}

// SetName sets the "name" field.
func (ouo *OwnerUpdateOne) SetName(s string) *OwnerUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OwnerUpdateOne) SetNillableName(s *string) *OwnerUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetAge sets the "age" field.
func (ouo *OwnerUpdateOne) SetAge(i int) *OwnerUpdateOne {
	ouo.mutation.ResetAge()
	ouo.mutation.SetAge(i)
	return ouo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (ouo *OwnerUpdateOne) SetNillableAge(i *int) *OwnerUpdateOne {
	if i != nil {
		ouo.SetAge(*i)
	}
	return ouo
}

// AddAge adds i to the "age" field.
func (ouo *OwnerUpdateOne) AddAge(i int) *OwnerUpdateOne {
	ouo.mutation.AddAge(i)
	return ouo
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (ouo *OwnerUpdateOne) AddPetIDs(ids ...int) *OwnerUpdateOne {
	ouo.mutation.AddPetIDs(ids...)
	return ouo
}

// AddPets adds the "pets" edges to the Pet entity.
func (ouo *OwnerUpdateOne) AddPets(p ...*Pet) *OwnerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddPetIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ouo *OwnerUpdateOne) Mutation() *OwnerMutation {
	return ouo.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (ouo *OwnerUpdateOne) ClearPets() *OwnerUpdateOne {
	ouo.mutation.ClearPets()
	return ouo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (ouo *OwnerUpdateOne) RemovePetIDs(ids ...int) *OwnerUpdateOne {
	ouo.mutation.RemovePetIDs(ids...)
	return ouo
}

// RemovePets removes "pets" edges to Pet entities.
func (ouo *OwnerUpdateOne) RemovePets(p ...*Pet) *OwnerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemovePetIDs(ids...)
}

// Where appends a list predicates to the OwnerUpdate builder.
func (ouo *OwnerUpdateOne) Where(ps ...predicate.Owner) *OwnerUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OwnerUpdateOne) Select(field string, fields ...string) *OwnerUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Owner entity.
func (ouo *OwnerUpdateOne) Save(ctx context.Context) (*Owner, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OwnerUpdateOne) SaveX(ctx context.Context) *Owner {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OwnerUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OwnerUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OwnerUpdateOne) sqlSave(ctx context.Context) (_node *Owner, err error) {
	_spec := sqlgraph.NewUpdateSpec(owner.Table, owner.Columns, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Owner.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, owner.FieldID)
		for _, f := range fields {
			if !owner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != owner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(owner.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Age(); ok {
		_spec.SetField(owner.FieldAge, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedAge(); ok {
		_spec.AddField(owner.FieldAge, field.TypeInt, value)
	}
	if ouo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !ouo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Owner{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
