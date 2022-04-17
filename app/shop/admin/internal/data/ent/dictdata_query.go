// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/dictdata"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/predicate"
)

// DictDataQuery is the builder for querying DictData entities.
type DictDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DictData
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictDataQuery builder.
func (ddq *DictDataQuery) Where(ps ...predicate.DictData) *DictDataQuery {
	ddq.predicates = append(ddq.predicates, ps...)
	return ddq
}

// Limit adds a limit step to the query.
func (ddq *DictDataQuery) Limit(limit int) *DictDataQuery {
	ddq.limit = &limit
	return ddq
}

// Offset adds an offset step to the query.
func (ddq *DictDataQuery) Offset(offset int) *DictDataQuery {
	ddq.offset = &offset
	return ddq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ddq *DictDataQuery) Unique(unique bool) *DictDataQuery {
	ddq.unique = &unique
	return ddq
}

// Order adds an order step to the query.
func (ddq *DictDataQuery) Order(o ...OrderFunc) *DictDataQuery {
	ddq.order = append(ddq.order, o...)
	return ddq
}

// First returns the first DictData entity from the query.
// Returns a *NotFoundError when no DictData was found.
func (ddq *DictDataQuery) First(ctx context.Context) (*DictData, error) {
	nodes, err := ddq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dictdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ddq *DictDataQuery) FirstX(ctx context.Context) *DictData {
	node, err := ddq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DictData ID from the query.
// Returns a *NotFoundError when no DictData ID was found.
func (ddq *DictDataQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ddq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dictdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ddq *DictDataQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ddq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DictData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DictData entity is found.
// Returns a *NotFoundError when no DictData entities are found.
func (ddq *DictDataQuery) Only(ctx context.Context) (*DictData, error) {
	nodes, err := ddq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dictdata.Label}
	default:
		return nil, &NotSingularError{dictdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ddq *DictDataQuery) OnlyX(ctx context.Context) *DictData {
	node, err := ddq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DictData ID in the query.
// Returns a *NotSingularError when more than one DictData ID is found.
// Returns a *NotFoundError when no entities are found.
func (ddq *DictDataQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ddq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = &NotSingularError{dictdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ddq *DictDataQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ddq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DictDataSlice.
func (ddq *DictDataQuery) All(ctx context.Context) ([]*DictData, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ddq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ddq *DictDataQuery) AllX(ctx context.Context) []*DictData {
	nodes, err := ddq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DictData IDs.
func (ddq *DictDataQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ddq.Select(dictdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ddq *DictDataQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ddq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ddq *DictDataQuery) Count(ctx context.Context) (int, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ddq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ddq *DictDataQuery) CountX(ctx context.Context) int {
	count, err := ddq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ddq *DictDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ddq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ddq *DictDataQuery) ExistX(ctx context.Context) bool {
	exist, err := ddq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ddq *DictDataQuery) Clone() *DictDataQuery {
	if ddq == nil {
		return nil
	}
	return &DictDataQuery{
		config:     ddq.config,
		limit:      ddq.limit,
		offset:     ddq.offset,
		order:      append([]OrderFunc{}, ddq.order...),
		predicates: append([]predicate.DictData{}, ddq.predicates...),
		// clone intermediate query.
		sql:    ddq.sql.Clone(),
		path:   ddq.path,
		unique: ddq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DictData.Query().
//		GroupBy(dictdata.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ddq *DictDataQuery) GroupBy(field string, fields ...string) *DictDataGroupBy {
	group := &DictDataGroupBy{config: ddq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ddq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ddq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DictData.Query().
//		Select(dictdata.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ddq *DictDataQuery) Select(fields ...string) *DictDataSelect {
	ddq.fields = append(ddq.fields, fields...)
	return &DictDataSelect{DictDataQuery: ddq}
}

func (ddq *DictDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ddq.fields {
		if !dictdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ddq.path != nil {
		prev, err := ddq.path(ctx)
		if err != nil {
			return err
		}
		ddq.sql = prev
	}
	return nil
}

func (ddq *DictDataQuery) sqlAll(ctx context.Context) ([]*DictData, error) {
	var (
		nodes = []*DictData{}
		_spec = ddq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DictData{config: ddq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ddq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ddq *DictDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ddq.querySpec()
	_spec.Node.Columns = ddq.fields
	if len(ddq.fields) > 0 {
		_spec.Unique = ddq.unique != nil && *ddq.unique
	}
	return sqlgraph.CountNodes(ctx, ddq.driver, _spec)
}

func (ddq *DictDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ddq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ddq *DictDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dictdata.Table,
			Columns: dictdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: dictdata.FieldID,
			},
		},
		From:   ddq.sql,
		Unique: true,
	}
	if unique := ddq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ddq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictdata.FieldID)
		for i := range fields {
			if fields[i] != dictdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ddq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ddq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ddq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ddq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ddq *DictDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ddq.driver.Dialect())
	t1 := builder.Table(dictdata.Table)
	columns := ddq.fields
	if len(columns) == 0 {
		columns = dictdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ddq.sql != nil {
		selector = ddq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ddq.unique != nil && *ddq.unique {
		selector.Distinct()
	}
	for _, p := range ddq.predicates {
		p(selector)
	}
	for _, p := range ddq.order {
		p(selector)
	}
	if offset := ddq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ddq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DictDataGroupBy is the group-by builder for DictData entities.
type DictDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ddgb *DictDataGroupBy) Aggregate(fns ...AggregateFunc) *DictDataGroupBy {
	ddgb.fns = append(ddgb.fns, fns...)
	return ddgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ddgb *DictDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ddgb.path(ctx)
	if err != nil {
		return err
	}
	ddgb.sql = query
	return ddgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ddgb *DictDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ddgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ddgb.fields) > 1 {
		return nil, errors.New("ent: DictDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ddgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ddgb *DictDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := ddgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ddgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ddgb *DictDataGroupBy) StringX(ctx context.Context) string {
	v, err := ddgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ddgb.fields) > 1 {
		return nil, errors.New("ent: DictDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ddgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ddgb *DictDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := ddgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ddgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ddgb *DictDataGroupBy) IntX(ctx context.Context) int {
	v, err := ddgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ddgb.fields) > 1 {
		return nil, errors.New("ent: DictDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ddgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ddgb *DictDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ddgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ddgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ddgb *DictDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ddgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ddgb.fields) > 1 {
		return nil, errors.New("ent: DictDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ddgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ddgb *DictDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ddgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddgb *DictDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ddgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ddgb *DictDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := ddgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ddgb *DictDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ddgb.fields {
		if !dictdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ddgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ddgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ddgb *DictDataGroupBy) sqlQuery() *sql.Selector {
	selector := ddgb.sql.Select()
	aggregation := make([]string, 0, len(ddgb.fns))
	for _, fn := range ddgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ddgb.fields)+len(ddgb.fns))
		for _, f := range ddgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ddgb.fields...)...)
}

// DictDataSelect is the builder for selecting fields of DictData entities.
type DictDataSelect struct {
	*DictDataQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dds *DictDataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dds.prepareQuery(ctx); err != nil {
		return err
	}
	dds.sql = dds.DictDataQuery.sqlQuery(ctx)
	return dds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dds *DictDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dds.fields) > 1 {
		return nil, errors.New("ent: DictDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dds *DictDataSelect) StringsX(ctx context.Context) []string {
	v, err := dds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dds *DictDataSelect) StringX(ctx context.Context) string {
	v, err := dds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dds.fields) > 1 {
		return nil, errors.New("ent: DictDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dds *DictDataSelect) IntsX(ctx context.Context) []int {
	v, err := dds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dds *DictDataSelect) IntX(ctx context.Context) int {
	v, err := dds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dds.fields) > 1 {
		return nil, errors.New("ent: DictDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dds *DictDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dds *DictDataSelect) Float64X(ctx context.Context) float64 {
	v, err := dds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dds.fields) > 1 {
		return nil, errors.New("ent: DictDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dds *DictDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := dds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dds *DictDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dictdata.Label}
	default:
		err = fmt.Errorf("ent: DictDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dds *DictDataSelect) BoolX(ctx context.Context) bool {
	v, err := dds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dds *DictDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dds.sql.Query()
	if err := dds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}