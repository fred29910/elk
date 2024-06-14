// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/simple/ent/collar"
	"github.com/masseelch/elk/internal/simple/ent/pet"
)

// CollarCreate is the builder for creating a Collar entity.
type CollarCreate struct {
	config
	mutation *CollarMutation
	hooks    []Hook
}

// SetColor sets the "color" field.
func (cc *CollarCreate) SetColor(c collar.Color) *CollarCreate {
	cc.mutation.SetColor(c)
	return cc
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (cc *CollarCreate) SetPetID(id string) *CollarCreate {
	cc.mutation.SetPetID(id)
	return cc
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (cc *CollarCreate) SetNillablePetID(id *string) *CollarCreate {
	if id != nil {
		cc = cc.SetPetID(*id)
	}
	return cc
}

// SetPet sets the "pet" edge to the Pet entity.
func (cc *CollarCreate) SetPet(p *Pet) *CollarCreate {
	return cc.SetPetID(p.ID)
}

// Mutation returns the CollarMutation object of the builder.
func (cc *CollarCreate) Mutation() *CollarMutation {
	return cc.mutation
}

// Save creates the Collar in the database.
func (cc *CollarCreate) Save(ctx context.Context) (*Collar, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollarCreate) SaveX(ctx context.Context) *Collar {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollarCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollarCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollarCreate) check() error {
	if _, ok := cc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "Collar.color"`)}
	}
	if v, ok := cc.mutation.Color(); ok {
		if err := collar.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Collar.color": %w`, err)}
		}
	}
	return nil
}

func (cc *CollarCreate) sqlSave(ctx context.Context) (*Collar, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollarCreate) createSpec() (*Collar, *sqlgraph.CreateSpec) {
	var (
		_node = &Collar{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(collar.Table, sqlgraph.NewFieldSpec(collar.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Color(); ok {
		_spec.SetField(collar.FieldColor, field.TypeEnum, value)
		_node.Color = value
	}
	if nodes := cc.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   collar.PetTable,
			Columns: []string{collar.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pet_collar = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CollarCreateBulk is the builder for creating many Collar entities in bulk.
type CollarCreateBulk struct {
	config
	err      error
	builders []*CollarCreate
}

// Save creates the Collar entities in the database.
func (ccb *CollarCreateBulk) Save(ctx context.Context) ([]*Collar, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Collar, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollarMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollarCreateBulk) SaveX(ctx context.Context) []*Collar {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollarCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollarCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
