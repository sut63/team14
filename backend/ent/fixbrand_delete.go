// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// FixbrandDelete is the builder for deleting a Fixbrand entity.
type FixbrandDelete struct {
	config
	hooks      []Hook
	mutation   *FixbrandMutation
	predicates []predicate.Fixbrand
}

// Where adds a new predicate to the delete builder.
func (fd *FixbrandDelete) Where(ps ...predicate.Fixbrand) *FixbrandDelete {
	fd.predicates = append(fd.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FixbrandDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixbrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FixbrandDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FixbrandDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fixbrand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixbrand.FieldID,
			},
		},
	}
	if ps := fd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FixbrandDeleteOne is the builder for deleting a single Fixbrand entity.
type FixbrandDeleteOne struct {
	fd *FixbrandDelete
}

// Exec executes the deletion query.
func (fdo *FixbrandDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fixbrand.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FixbrandDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
