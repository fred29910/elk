// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/pets/ent/playgroup"
	"github.com/masseelch/elk/internal/pets/ent/predicate"
)

// PlayGroupDelete is the builder for deleting a PlayGroup entity.
type PlayGroupDelete struct {
	config
	hooks    []Hook
	mutation *PlayGroupMutation
}

// Where appends a list predicates to the PlayGroupDelete builder.
func (pgd *PlayGroupDelete) Where(ps ...predicate.PlayGroup) *PlayGroupDelete {
	pgd.mutation.Where(ps...)
	return pgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pgd *PlayGroupDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pgd.sqlExec, pgd.mutation, pgd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pgd *PlayGroupDelete) ExecX(ctx context.Context) int {
	n, err := pgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pgd *PlayGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(playgroup.Table, sqlgraph.NewFieldSpec(playgroup.FieldID, field.TypeInt))
	if ps := pgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pgd.mutation.done = true
	return affected, err
}

// PlayGroupDeleteOne is the builder for deleting a single PlayGroup entity.
type PlayGroupDeleteOne struct {
	pgd *PlayGroupDelete
}

// Where appends a list predicates to the PlayGroupDelete builder.
func (pgdo *PlayGroupDeleteOne) Where(ps ...predicate.PlayGroup) *PlayGroupDeleteOne {
	pgdo.pgd.mutation.Where(ps...)
	return pgdo
}

// Exec executes the deletion query.
func (pgdo *PlayGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := pgdo.pgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{playgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pgdo *PlayGroupDeleteOne) ExecX(ctx context.Context) {
	if err := pgdo.Exec(ctx); err != nil {
		panic(err)
	}
}
