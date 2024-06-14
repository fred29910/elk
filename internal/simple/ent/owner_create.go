// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/masseelch/elk/internal/simple/ent/owner"
	"github.com/masseelch/elk/internal/simple/ent/pet"
)

// OwnerCreate is the builder for creating a Owner entity.
type OwnerCreate struct {
	config
	mutation *OwnerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (oc *OwnerCreate) SetName(s string) *OwnerCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetAge sets the "age" field.
func (oc *OwnerCreate) SetAge(i int) *OwnerCreate {
	oc.mutation.SetAge(i)
	return oc
}

// SetID sets the "id" field.
func (oc *OwnerCreate) SetID(u uuid.UUID) *OwnerCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OwnerCreate) SetNillableID(u *uuid.UUID) *OwnerCreate {
	if u != nil {
		oc.SetID(*u)
	}
	return oc
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (oc *OwnerCreate) AddPetIDs(ids ...string) *OwnerCreate {
	oc.mutation.AddPetIDs(ids...)
	return oc
}

// AddPets adds the "pets" edges to the Pet entity.
func (oc *OwnerCreate) AddPets(p ...*Pet) *OwnerCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPetIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (oc *OwnerCreate) Mutation() *OwnerMutation {
	return oc.mutation
}

// Save creates the Owner in the database.
func (oc *OwnerCreate) Save(ctx context.Context) (*Owner, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OwnerCreate) SaveX(ctx context.Context) *Owner {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OwnerCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OwnerCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OwnerCreate) defaults() {
	if _, ok := oc.mutation.ID(); !ok {
		v := owner.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OwnerCreate) check() error {
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Owner.name"`)}
	}
	if _, ok := oc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Owner.age"`)}
	}
	return nil
}

func (oc *OwnerCreate) sqlSave(ctx context.Context) (*Owner, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OwnerCreate) createSpec() (*Owner, *sqlgraph.CreateSpec) {
	var (
		_node = &Owner{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(owner.Table, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(owner.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oc.mutation.Age(); ok {
		_spec.SetField(owner.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if nodes := oc.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OwnerCreateBulk is the builder for creating many Owner entities in bulk.
type OwnerCreateBulk struct {
	config
	err      error
	builders []*OwnerCreate
}

// Save creates the Owner entities in the database.
func (ocb *OwnerCreateBulk) Save(ctx context.Context) ([]*Owner, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Owner, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OwnerMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OwnerCreateBulk) SaveX(ctx context.Context) []*Owner {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OwnerCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OwnerCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
