// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/brand"
	"github.com/tanapon395/playlist-video/ent/product"
)

// BrandCreate is the builder for creating a Brand entity.
type BrandCreate struct {
	config
	mutation *BrandMutation
	hooks    []Hook
}

// SetBrandname sets the Brandname field.
func (bc *BrandCreate) SetBrandname(s string) *BrandCreate {
	bc.mutation.SetBrandname(s)
	return bc
}

// AddProductIDs adds the product edge to Product by ids.
func (bc *BrandCreate) AddProductIDs(ids ...int) *BrandCreate {
	bc.mutation.AddProductIDs(ids...)
	return bc
}

// AddProduct adds the product edges to Product.
func (bc *BrandCreate) AddProduct(p ...*Product) *BrandCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bc *BrandCreate) Mutation() *BrandMutation {
	return bc.mutation
}

// Save creates the Brand in the database.
func (bc *BrandCreate) Save(ctx context.Context) (*Brand, error) {
	if _, ok := bc.mutation.Brandname(); !ok {
		return nil, &ValidationError{Name: "Brandname", err: errors.New("ent: missing required field \"Brandname\"")}
	}
	var (
		err  error
		node *Brand
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrandCreate) SaveX(ctx context.Context) *Brand {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BrandCreate) sqlSave(ctx context.Context) (*Brand, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BrandCreate) createSpec() (*Brand, *sqlgraph.CreateSpec) {
	var (
		b     = &Brand{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: brand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: brand.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Brandname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldBrandname,
		})
		b.Brandname = value
	}
	if nodes := bc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
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
	return b, _spec
}
