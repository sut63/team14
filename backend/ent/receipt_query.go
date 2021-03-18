// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// ReceiptQuery is the builder for querying Receipt entities.
type ReceiptQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Receipt
	// eager-loading edges.
	withPaymenttype *PaymentTypeQuery
	withAdminrepair *AdminrepairQuery
	withPersonal    *PersonalQuery
	withCustomer    *CustomerQuery
	withProduct     *ProductQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *ReceiptQuery) Where(ps ...predicate.Receipt) *ReceiptQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *ReceiptQuery) Limit(limit int) *ReceiptQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *ReceiptQuery) Offset(offset int) *ReceiptQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *ReceiptQuery) Order(o ...OrderFunc) *ReceiptQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryPaymenttype chains the current query on the paymenttype edge.
func (rq *ReceiptQuery) QueryPaymenttype() *PaymentTypeQuery {
	query := &PaymentTypeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, rq.sqlQuery()),
			sqlgraph.To(paymenttype.Table, paymenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receipt.PaymenttypeTable, receipt.PaymenttypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAdminrepair chains the current query on the adminrepair edge.
func (rq *ReceiptQuery) QueryAdminrepair() *AdminrepairQuery {
	query := &AdminrepairQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, rq.sqlQuery()),
			sqlgraph.To(adminrepair.Table, adminrepair.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receipt.AdminrepairTable, receipt.AdminrepairColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPersonal chains the current query on the personal edge.
func (rq *ReceiptQuery) QueryPersonal() *PersonalQuery {
	query := &PersonalQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, rq.sqlQuery()),
			sqlgraph.To(personal.Table, personal.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receipt.PersonalTable, receipt.PersonalColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the customer edge.
func (rq *ReceiptQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, rq.sqlQuery()),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receipt.CustomerTable, receipt.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the product edge.
func (rq *ReceiptQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, rq.sqlQuery()),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receipt.ProductTable, receipt.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Receipt entity in the query. Returns *NotFoundError when no receipt was found.
func (rq *ReceiptQuery) First(ctx context.Context) (*Receipt, error) {
	rs, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, &NotFoundError{receipt.Label}
	}
	return rs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReceiptQuery) FirstX(ctx context.Context) *Receipt {
	r, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return r
}

// FirstID returns the first Receipt id in the query. Returns *NotFoundError when no id was found.
func (rq *ReceiptQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{receipt.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rq *ReceiptQuery) FirstXID(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Receipt entity in the query, returns an error if not exactly one entity was returned.
func (rq *ReceiptQuery) Only(ctx context.Context) (*Receipt, error) {
	rs, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rs) {
	case 1:
		return rs[0], nil
	case 0:
		return nil, &NotFoundError{receipt.Label}
	default:
		return nil, &NotSingularError{receipt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReceiptQuery) OnlyX(ctx context.Context) *Receipt {
	r, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// OnlyID returns the only Receipt id in the query, returns an error if not exactly one id was returned.
func (rq *ReceiptQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = &NotSingularError{receipt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReceiptQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Receipts.
func (rq *ReceiptQuery) All(ctx context.Context) ([]*Receipt, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReceiptQuery) AllX(ctx context.Context) []*Receipt {
	rs, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// IDs executes the query and returns a list of Receipt ids.
func (rq *ReceiptQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(receipt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReceiptQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReceiptQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReceiptQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReceiptQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReceiptQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReceiptQuery) Clone() *ReceiptQuery {
	return &ReceiptQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Receipt{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithPaymenttype tells the query-builder to eager-loads the nodes that are connected to
// the "paymenttype" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithPaymenttype(opts ...func(*PaymentTypeQuery)) *ReceiptQuery {
	query := &PaymentTypeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPaymenttype = query
	return rq
}

//  WithAdminrepair tells the query-builder to eager-loads the nodes that are connected to
// the "adminrepair" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithAdminrepair(opts ...func(*AdminrepairQuery)) *ReceiptQuery {
	query := &AdminrepairQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withAdminrepair = query
	return rq
}

//  WithPersonal tells the query-builder to eager-loads the nodes that are connected to
// the "personal" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithPersonal(opts ...func(*PersonalQuery)) *ReceiptQuery {
	query := &PersonalQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPersonal = query
	return rq
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithCustomer(opts ...func(*CustomerQuery)) *ReceiptQuery {
	query := &CustomerQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withCustomer = query
	return rq
}

//  WithProduct tells the query-builder to eager-loads the nodes that are connected to
// the "product" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithProduct(opts ...func(*ProductQuery)) *ReceiptQuery {
	query := &ProductQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withProduct = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DateTime time.Time `json:"date_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Receipt.Query().
//		GroupBy(receipt.FieldDateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *ReceiptQuery) GroupBy(field string, fields ...string) *ReceiptGroupBy {
	group := &ReceiptGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		DateTime time.Time `json:"date_time,omitempty"`
//	}
//
//	client.Receipt.Query().
//		Select(receipt.FieldDateTime).
//		Scan(ctx, &v)
//
func (rq *ReceiptQuery) Select(field string, fields ...string) *ReceiptSelect {
	selector := &ReceiptSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *ReceiptQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReceiptQuery) sqlAll(ctx context.Context) ([]*Receipt, error) {
	var (
		nodes       = []*Receipt{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [5]bool{
			rq.withPaymenttype != nil,
			rq.withAdminrepair != nil,
			rq.withPersonal != nil,
			rq.withCustomer != nil,
			rq.withProduct != nil,
		}
	)
	if rq.withPaymenttype != nil || rq.withAdminrepair != nil || rq.withPersonal != nil || rq.withCustomer != nil || rq.withProduct != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, receipt.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Receipt{config: rq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withPaymenttype; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Receipt)
		for i := range nodes {
			if fk := nodes[i].paymenttype_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(paymenttype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "paymenttype_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Paymenttype = n
			}
		}
	}

	if query := rq.withAdminrepair; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Receipt)
		for i := range nodes {
			if fk := nodes[i].adminrepair_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(adminrepair.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adminrepair_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Adminrepair = n
			}
		}
	}

	if query := rq.withPersonal; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Receipt)
		for i := range nodes {
			if fk := nodes[i].personal_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(personal.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "personal_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Personal = n
			}
		}
	}

	if query := rq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Receipt)
		for i := range nodes {
			if fk := nodes[i].customer_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	if query := rq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Receipt)
		for i := range nodes {
			if fk := nodes[i].product_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	return nodes, nil
}

func (rq *ReceiptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReceiptQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *ReceiptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   receipt.Table,
			Columns: receipt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: receipt.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReceiptQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(receipt.Table)
	selector := builder.Select(t1.Columns(receipt.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(receipt.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReceiptGroupBy is the builder for group-by Receipt entities.
type ReceiptGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReceiptGroupBy) Aggregate(fns ...AggregateFunc) *ReceiptGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *ReceiptGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *ReceiptGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReceiptGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *ReceiptGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *ReceiptGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReceiptGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *ReceiptGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *ReceiptGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReceiptGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *ReceiptGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *ReceiptGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReceiptGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *ReceiptGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *ReceiptGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *ReceiptGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *ReceiptGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgb.sqlQuery().Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *ReceiptGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// ReceiptSelect is the builder for select fields of Receipt entities.
type ReceiptSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *ReceiptSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *ReceiptSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReceiptSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *ReceiptSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *ReceiptSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReceiptSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *ReceiptSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *ReceiptSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReceiptSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *ReceiptSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *ReceiptSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReceiptSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *ReceiptSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *ReceiptSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = fmt.Errorf("ent: ReceiptSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *ReceiptSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *ReceiptSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *ReceiptSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
