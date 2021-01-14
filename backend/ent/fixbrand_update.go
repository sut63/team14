// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// FixbrandUpdate is the builder for updating Fixbrand entities.
type FixbrandUpdate struct {
	config
	hooks      []Hook
	mutation   *FixbrandMutation
	predicates []predicate.Fixbrand
}

// Where adds a new predicate for the builder.
func (fu *FixbrandUpdate) Where(ps ...predicate.Fixbrand) *FixbrandUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetFixbrandname sets the fixbrandname field.
func (fu *FixbrandUpdate) SetFixbrandname(s string) *FixbrandUpdate {
	fu.mutation.SetFixbrandname(s)
	return fu
}

// AddFixIDs adds the fix edge to Fix by ids.
func (fu *FixbrandUpdate) AddFixIDs(ids ...int) *FixbrandUpdate {
	fu.mutation.AddFixIDs(ids...)
	return fu
}

// AddFix adds the fix edges to Fix.
func (fu *FixbrandUpdate) AddFix(f ...*Fix) *FixbrandUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddFixIDs(ids...)
}

// Mutation returns the FixbrandMutation object of the builder.
func (fu *FixbrandUpdate) Mutation() *FixbrandMutation {
	return fu.mutation
}

// RemoveFixIDs removes the fix edge to Fix by ids.
func (fu *FixbrandUpdate) RemoveFixIDs(ids ...int) *FixbrandUpdate {
	fu.mutation.RemoveFixIDs(ids...)
	return fu
}

// RemoveFix removes fix edges to Fix.
func (fu *FixbrandUpdate) RemoveFix(f ...*Fix) *FixbrandUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveFixIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FixbrandUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixbrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FixbrandUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FixbrandUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FixbrandUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FixbrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fixbrand.Table,
			Columns: fixbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixbrand.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Fixbrandname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixbrand.FieldFixbrandname,
		})
	}
	if nodes := fu.mutation.RemovedFixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixbrand.FixTable,
			Columns: []string{fixbrand.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixbrand.FixTable,
			Columns: []string{fixbrand.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fixbrand.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FixbrandUpdateOne is the builder for updating a single Fixbrand entity.
type FixbrandUpdateOne struct {
	config
	hooks    []Hook
	mutation *FixbrandMutation
}

// SetFixbrandname sets the fixbrandname field.
func (fuo *FixbrandUpdateOne) SetFixbrandname(s string) *FixbrandUpdateOne {
	fuo.mutation.SetFixbrandname(s)
	return fuo
}

// AddFixIDs adds the fix edge to Fix by ids.
func (fuo *FixbrandUpdateOne) AddFixIDs(ids ...int) *FixbrandUpdateOne {
	fuo.mutation.AddFixIDs(ids...)
	return fuo
}

// AddFix adds the fix edges to Fix.
func (fuo *FixbrandUpdateOne) AddFix(f ...*Fix) *FixbrandUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddFixIDs(ids...)
}

// Mutation returns the FixbrandMutation object of the builder.
func (fuo *FixbrandUpdateOne) Mutation() *FixbrandMutation {
	return fuo.mutation
}

// RemoveFixIDs removes the fix edge to Fix by ids.
func (fuo *FixbrandUpdateOne) RemoveFixIDs(ids ...int) *FixbrandUpdateOne {
	fuo.mutation.RemoveFixIDs(ids...)
	return fuo
}

// RemoveFix removes fix edges to Fix.
func (fuo *FixbrandUpdateOne) RemoveFix(f ...*Fix) *FixbrandUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveFixIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fuo *FixbrandUpdateOne) Save(ctx context.Context) (*Fixbrand, error) {

	var (
		err  error
		node *Fixbrand
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixbrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FixbrandUpdateOne) SaveX(ctx context.Context) *Fixbrand {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FixbrandUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FixbrandUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FixbrandUpdateOne) sqlSave(ctx context.Context) (f *Fixbrand, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fixbrand.Table,
			Columns: fixbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixbrand.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Fixbrand.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Fixbrandname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixbrand.FieldFixbrandname,
		})
	}
	if nodes := fuo.mutation.RemovedFixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixbrand.FixTable,
			Columns: []string{fixbrand.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixbrand.FixTable,
			Columns: []string{fixbrand.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	f = &Fixbrand{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fixbrand.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}