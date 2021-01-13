// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// PaymentTypeQuery is the builder for querying PaymentType entities.
type PaymentTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.PaymentType
	// eager-loading edges.
	withReceipt *ReceiptQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ptq *PaymentTypeQuery) Where(ps ...predicate.PaymentType) *PaymentTypeQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit adds a limit step to the query.
func (ptq *PaymentTypeQuery) Limit(limit int) *PaymentTypeQuery {
	ptq.limit = &limit
	return ptq
}

// Offset adds an offset step to the query.
func (ptq *PaymentTypeQuery) Offset(offset int) *PaymentTypeQuery {
	ptq.offset = &offset
	return ptq
}

// Order adds an order step to the query.
func (ptq *PaymentTypeQuery) Order(o ...OrderFunc) *PaymentTypeQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryReceipt chains the current query on the receipt edge.
func (ptq *PaymentTypeQuery) QueryReceipt() *ReceiptQuery {
	query := &ReceiptQuery{config: ptq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(paymenttype.Table, paymenttype.FieldID, ptq.sqlQuery()),
			sqlgraph.To(receipt.Table, receipt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, paymenttype.ReceiptTable, paymenttype.ReceiptColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaymentType entity in the query. Returns *NotFoundError when no paymenttype was found.
func (ptq *PaymentTypeQuery) First(ctx context.Context) (*PaymentType, error) {
	pts, err := ptq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pts) == 0 {
		return nil, &NotFoundError{paymenttype.Label}
	}
	return pts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PaymentTypeQuery) FirstX(ctx context.Context) *PaymentType {
	pt, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pt
}

// FirstID returns the first PaymentType id in the query. Returns *NotFoundError when no id was found.
func (ptq *PaymentTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{paymenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ptq *PaymentTypeQuery) FirstXID(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only PaymentType entity in the query, returns an error if not exactly one entity was returned.
func (ptq *PaymentTypeQuery) Only(ctx context.Context) (*PaymentType, error) {
	pts, err := ptq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pts) {
	case 1:
		return pts[0], nil
	case 0:
		return nil, &NotFoundError{paymenttype.Label}
	default:
		return nil, &NotSingularError{paymenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PaymentTypeQuery) OnlyX(ctx context.Context) *PaymentType {
	pt, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pt
}

// OnlyID returns the only PaymentType id in the query, returns an error if not exactly one id was returned.
func (ptq *PaymentTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = &NotSingularError{paymenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PaymentTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaymentTypes.
func (ptq *PaymentTypeQuery) All(ctx context.Context) ([]*PaymentType, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PaymentTypeQuery) AllX(ctx context.Context) []*PaymentType {
	pts, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pts
}

// IDs executes the query and returns a list of PaymentType ids.
func (ptq *PaymentTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ptq.Select(paymenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PaymentTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PaymentTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PaymentTypeQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PaymentTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PaymentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PaymentTypeQuery) Clone() *PaymentTypeQuery {
	return &PaymentTypeQuery{
		config:     ptq.config,
		limit:      ptq.limit,
		offset:     ptq.offset,
		order:      append([]OrderFunc{}, ptq.order...),
		unique:     append([]string{}, ptq.unique...),
		predicates: append([]predicate.PaymentType{}, ptq.predicates...),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

//  WithReceipt tells the query-builder to eager-loads the nodes that are connected to
// the "receipt" edge. The optional arguments used to configure the query builder of the edge.
func (ptq *PaymentTypeQuery) WithReceipt(opts ...func(*ReceiptQuery)) *PaymentTypeQuery {
	query := &ReceiptQuery{config: ptq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptq.withReceipt = query
	return ptq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Typename string `json:"Typename,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PaymentType.Query().
//		GroupBy(paymenttype.FieldTypename).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptq *PaymentTypeQuery) GroupBy(field string, fields ...string) *PaymentTypeGroupBy {
	group := &PaymentTypeGroupBy{config: ptq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Typename string `json:"Typename,omitempty"`
//	}
//
//	client.PaymentType.Query().
//		Select(paymenttype.FieldTypename).
//		Scan(ctx, &v)
//
func (ptq *PaymentTypeQuery) Select(field string, fields ...string) *PaymentTypeSelect {
	selector := &PaymentTypeSelect{config: ptq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptq.sqlQuery(), nil
	}
	return selector
}

func (ptq *PaymentTypeQuery) prepareQuery(ctx context.Context) error {
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PaymentTypeQuery) sqlAll(ctx context.Context) ([]*PaymentType, error) {
	var (
		nodes       = []*PaymentType{}
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withReceipt != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &PaymentType{config: ptq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptq.withReceipt; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*PaymentType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Receipt(func(s *sql.Selector) {
			s.Where(sql.InValues(paymenttype.ReceiptColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.paymenttype_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "paymenttype_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "paymenttype_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Receipt = append(node.Edges.Receipt, n)
		}
	}

	return nodes, nil
}

func (ptq *PaymentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PaymentTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ptq *PaymentTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymenttype.Table,
			Columns: paymenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymenttype.FieldID,
			},
		},
		From:   ptq.sql,
		Unique: true,
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PaymentTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(paymenttype.Table)
	selector := builder.Select(t1.Columns(paymenttype.Columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(paymenttype.Columns...)...)
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaymentTypeGroupBy is the builder for group-by PaymentType entities.
type PaymentTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PaymentTypeGroupBy) Aggregate(fns ...AggregateFunc) *PaymentTypeGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ptgb *PaymentTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptgb.path(ctx)
	if err != nil {
		return err
	}
	ptgb.sql = query
	return ptgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ptgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ptgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ptgb *PaymentTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptgb *PaymentTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptgb *PaymentTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ptgb.sqlQuery().Query()
	if err := ptgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptgb *PaymentTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ptgb.sql
	columns := make([]string, 0, len(ptgb.fields)+len(ptgb.fns))
	columns = append(columns, ptgb.fields...)
	for _, fn := range ptgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ptgb.fields...)
}

// PaymentTypeSelect is the builder for select fields of PaymentType entities.
type PaymentTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (pts *PaymentTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := pts.path(ctx)
	if err != nil {
		return err
	}
	pts.sql = query
	return pts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pts *PaymentTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pts *PaymentTypeSelect) StringsX(ctx context.Context) []string {
	v, err := pts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pts *PaymentTypeSelect) StringX(ctx context.Context) string {
	v, err := pts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pts *PaymentTypeSelect) IntsX(ctx context.Context) []int {
	v, err := pts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pts *PaymentTypeSelect) IntX(ctx context.Context) int {
	v, err := pts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pts *PaymentTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pts *PaymentTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := pts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PaymentTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pts *PaymentTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (pts *PaymentTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttype.Label}
	default:
		err = fmt.Errorf("ent: PaymentTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pts *PaymentTypeSelect) BoolX(ctx context.Context) bool {
	v, err := pts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pts *PaymentTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pts.sqlQuery().Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pts *PaymentTypeSelect) sqlQuery() sql.Querier {
	selector := pts.sql
	selector.Select(selector.Columns(pts.fields...)...)
	return selector
}
