// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
)

// TypeproductCreate is the builder for creating a Typeproduct entity.
type TypeproductCreate struct {
	config
	mutation *TypeproductMutation
	hooks    []Hook
}

// SetTypeproductname sets the Typeproductname field.
func (tc *TypeproductCreate) SetTypeproductname(s string) *TypeproductCreate {
	tc.mutation.SetTypeproductname(s)
	return tc
}

// AddProductIDs adds the product edge to Product by ids.
func (tc *TypeproductCreate) AddProductIDs(ids ...int) *TypeproductCreate {
	tc.mutation.AddProductIDs(ids...)
	return tc
}

// AddProduct adds the product edges to Product.
func (tc *TypeproductCreate) AddProduct(p ...*Product) *TypeproductCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddProductIDs(ids...)
}

// Mutation returns the TypeproductMutation object of the builder.
func (tc *TypeproductCreate) Mutation() *TypeproductMutation {
	return tc.mutation
}

// Save creates the Typeproduct in the database.
func (tc *TypeproductCreate) Save(ctx context.Context) (*Typeproduct, error) {
	if _, ok := tc.mutation.Typeproductname(); !ok {
		return nil, &ValidationError{Name: "Typeproductname", err: errors.New("ent: missing required field \"Typeproductname\"")}
	}
	var (
		err  error
		node *Typeproduct
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeproductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TypeproductCreate) SaveX(ctx context.Context) *Typeproduct {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TypeproductCreate) sqlSave(ctx context.Context) (*Typeproduct, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TypeproductCreate) createSpec() (*Typeproduct, *sqlgraph.CreateSpec) {
	var (
		t     = &Typeproduct{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: typeproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeproduct.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Typeproductname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typeproduct.FieldTypeproductname,
		})
		t.Typeproductname = value
	}
	if nodes := tc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typeproduct.ProductTable,
			Columns: []string{typeproduct.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
