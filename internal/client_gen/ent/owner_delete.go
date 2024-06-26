// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/client_gen/ent/owner"
	"github.com/masseelch/elk/internal/client_gen/ent/predicate"
)

// OwnerDelete is the builder for deleting a Owner entity.
type OwnerDelete struct {
	config
	hooks    []Hook
	mutation *OwnerMutation
}

// Where appends a list predicates to the OwnerDelete builder.
func (od *OwnerDelete) Where(ps ...predicate.Owner) *OwnerDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OwnerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OwnerDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(owner.Table, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeUUID))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OwnerDeleteOne is the builder for deleting a single Owner entity.
type OwnerDeleteOne struct {
	od *OwnerDelete
}

// Where appends a list predicates to the OwnerDelete builder.
func (odo *OwnerDeleteOne) Where(ps ...predicate.Owner) *OwnerDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{owner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OwnerDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
