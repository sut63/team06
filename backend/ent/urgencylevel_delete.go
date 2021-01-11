// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/urgencylevel"
)

// UrgencyLevelDelete is the builder for deleting a UrgencyLevel entity.
type UrgencyLevelDelete struct {
	config
	hooks    []Hook
	mutation *UrgencyLevelMutation
}

// Where adds a new predicate to the UrgencyLevelDelete builder.
func (uld *UrgencyLevelDelete) Where(ps ...predicate.UrgencyLevel) *UrgencyLevelDelete {
	uld.mutation.predicates = append(uld.mutation.predicates, ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UrgencyLevelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uld.hooks) == 0 {
		affected, err = uld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UrgencyLevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uld.mutation = mutation
			affected, err = uld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uld.hooks) - 1; i >= 0; i-- {
			mut = uld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UrgencyLevelDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UrgencyLevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: urgencylevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: urgencylevel.FieldID,
			},
		},
	}
	if ps := uld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
}

// UrgencyLevelDeleteOne is the builder for deleting a single UrgencyLevel entity.
type UrgencyLevelDeleteOne struct {
	uld *UrgencyLevelDelete
}

// Exec executes the deletion query.
func (uldo *UrgencyLevelDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{urgencylevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UrgencyLevelDeleteOne) ExecX(ctx context.Context) {
	uldo.uld.ExecX(ctx)
}
