// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// PaymentTypeCreate is the builder for creating a PaymentType entity.
type PaymentTypeCreate struct {
	config
	mutation *PaymentTypeMutation
	hooks    []Hook
}

// SetTypename sets the Typename field.
func (ptc *PaymentTypeCreate) SetTypename(s string) *PaymentTypeCreate {
	ptc.mutation.SetTypename(s)
	return ptc
}

// AddReceiptIDs adds the receipt edge to Receipt by ids.
func (ptc *PaymentTypeCreate) AddReceiptIDs(ids ...int) *PaymentTypeCreate {
	ptc.mutation.AddReceiptIDs(ids...)
	return ptc
}

// AddReceipt adds the receipt edges to Receipt.
func (ptc *PaymentTypeCreate) AddReceipt(r ...*Receipt) *PaymentTypeCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ptc.AddReceiptIDs(ids...)
}

// Mutation returns the PaymentTypeMutation object of the builder.
func (ptc *PaymentTypeCreate) Mutation() *PaymentTypeMutation {
	return ptc.mutation
}

// Save creates the PaymentType in the database.
func (ptc *PaymentTypeCreate) Save(ctx context.Context) (*PaymentType, error) {
	if _, ok := ptc.mutation.Typename(); !ok {
		return nil, &ValidationError{Name: "Typename", err: errors.New("ent: missing required field \"Typename\"")}
	}
	var (
		err  error
		node *PaymentType
	)
	if len(ptc.hooks) == 0 {
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptc.mutation = mutation
			node, err = ptc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			mut = ptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PaymentTypeCreate) SaveX(ctx context.Context) *PaymentType {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptc *PaymentTypeCreate) sqlSave(ctx context.Context) (*PaymentType, error) {
	pt, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pt.ID = int(id)
	return pt, nil
}

func (ptc *PaymentTypeCreate) createSpec() (*PaymentType, *sqlgraph.CreateSpec) {
	var (
		pt    = &PaymentType{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paymenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymenttype.FieldID,
			},
		}
	)
	if value, ok := ptc.mutation.Typename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paymenttype.FieldTypename,
		})
		pt.Typename = value
	}
	if nodes := ptc.mutation.ReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymenttype.ReceiptTable,
			Columns: []string{paymenttype.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: receipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pt, _spec
}
