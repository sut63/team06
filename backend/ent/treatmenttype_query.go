// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/treatmenttype"
)

// TreatmentTypeQuery is the builder for querying TreatmentType entities.
type TreatmentTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.TreatmentType
	// eager-loading edges.
	withTreatmentTypeToDiagnosis *DiagnosisQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TreatmentTypeQuery builder.
func (ttq *TreatmentTypeQuery) Where(ps ...predicate.TreatmentType) *TreatmentTypeQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TreatmentTypeQuery) Limit(limit int) *TreatmentTypeQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TreatmentTypeQuery) Offset(offset int) *TreatmentTypeQuery {
	ttq.offset = &offset
	return ttq
}

// Order adds an order step to the query.
func (ttq *TreatmentTypeQuery) Order(o ...OrderFunc) *TreatmentTypeQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTreatmentTypeToDiagnosis chains the current query on the "TreatmentTypeToDiagnosis" edge.
func (ttq *TreatmentTypeQuery) QueryTreatmentTypeToDiagnosis() *DiagnosisQuery {
	query := &DiagnosisQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(treatmenttype.Table, treatmenttype.FieldID, selector),
			sqlgraph.To(diagnosis.Table, diagnosis.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, treatmenttype.TreatmentTypeToDiagnosisTable, treatmenttype.TreatmentTypeToDiagnosisColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TreatmentType entity from the query.
// Returns a *NotFoundError when no TreatmentType was found.
func (ttq *TreatmentTypeQuery) First(ctx context.Context) (*TreatmentType, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{treatmenttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) FirstX(ctx context.Context) *TreatmentType {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TreatmentType ID from the query.
// Returns a *NotFoundError when no TreatmentType ID was found.
func (ttq *TreatmentTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{treatmenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TreatmentType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TreatmentType entity is not found.
// Returns a *NotFoundError when no TreatmentType entities are found.
func (ttq *TreatmentTypeQuery) Only(ctx context.Context) (*TreatmentType, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{treatmenttype.Label}
	default:
		return nil, &NotSingularError{treatmenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) OnlyX(ctx context.Context) *TreatmentType {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TreatmentType ID in the query.
// Returns a *NotSingularError when exactly one TreatmentType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TreatmentTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = &NotSingularError{treatmenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TreatmentTypes.
func (ttq *TreatmentTypeQuery) All(ctx context.Context) ([]*TreatmentType, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) AllX(ctx context.Context) []*TreatmentType {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TreatmentType IDs.
func (ttq *TreatmentTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ttq.Select(treatmenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TreatmentTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TreatmentTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TreatmentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TreatmentTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TreatmentTypeQuery) Clone() *TreatmentTypeQuery {
	if ttq == nil {
		return nil
	}
	return &TreatmentTypeQuery{
		config:                       ttq.config,
		limit:                        ttq.limit,
		offset:                       ttq.offset,
		order:                        append([]OrderFunc{}, ttq.order...),
		predicates:                   append([]predicate.TreatmentType{}, ttq.predicates...),
		withTreatmentTypeToDiagnosis: ttq.withTreatmentTypeToDiagnosis.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithTreatmentTypeToDiagnosis tells the query-builder to eager-load the nodes that are connected to
// the "TreatmentTypeToDiagnosis" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TreatmentTypeQuery) WithTreatmentTypeToDiagnosis(opts ...func(*DiagnosisQuery)) *TreatmentTypeQuery {
	query := &DiagnosisQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTreatmentTypeToDiagnosis = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type string `json:"Type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TreatmentType.Query().
//		GroupBy(treatmenttype.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TreatmentTypeQuery) GroupBy(field string, fields ...string) *TreatmentTypeGroupBy {
	group := &TreatmentTypeGroupBy{config: ttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type string `json:"Type,omitempty"`
//	}
//
//	client.TreatmentType.Query().
//		Select(treatmenttype.FieldType).
//		Scan(ctx, &v)
//
func (ttq *TreatmentTypeQuery) Select(field string, fields ...string) *TreatmentTypeSelect {
	ttq.fields = append([]string{field}, fields...)
	return &TreatmentTypeSelect{TreatmentTypeQuery: ttq}
}

func (ttq *TreatmentTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !treatmenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TreatmentTypeQuery) sqlAll(ctx context.Context) ([]*TreatmentType, error) {
	var (
		nodes       = []*TreatmentType{}
		_spec       = ttq.querySpec()
		loadedTypes = [1]bool{
			ttq.withTreatmentTypeToDiagnosis != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TreatmentType{config: ttq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ttq.withTreatmentTypeToDiagnosis; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*TreatmentType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TreatmentTypeToDiagnosis = []*Diagnosis{}
		}
		query.withFKs = true
		query.Where(predicate.Diagnosis(func(s *sql.Selector) {
			s.Where(sql.InValues(treatmenttype.TreatmentTypeToDiagnosisColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.treatment_type_treatment_type_to_diagnosis
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "treatment_type_treatment_type_to_diagnosis" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "treatment_type_treatment_type_to_diagnosis" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.TreatmentTypeToDiagnosis = append(node.Edges.TreatmentTypeToDiagnosis, n)
		}
	}

	return nodes, nil
}

func (ttq *TreatmentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TreatmentTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ttq *TreatmentTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   treatmenttype.Table,
			Columns: treatmenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatmenttype.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, treatmenttype.FieldID)
		for i := range fields {
			if fields[i] != treatmenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, treatmenttype.ValidColumn)
			}
		}
	}
	return _spec
}

func (ttq *TreatmentTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(treatmenttype.Table)
	selector := builder.Select(t1.Columns(treatmenttype.Columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(treatmenttype.Columns...)...)
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector, treatmenttype.ValidColumn)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TreatmentTypeGroupBy is the group-by builder for TreatmentType entities.
type TreatmentTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TreatmentTypeGroupBy) Aggregate(fns ...AggregateFunc) *TreatmentTypeGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TreatmentTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TreatmentTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ttgb *TreatmentTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ttgb *TreatmentTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !treatmenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TreatmentTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql
	columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
	columns = append(columns, ttgb.fields...)
	for _, fn := range ttgb.fns {
		columns = append(columns, fn(selector, treatmenttype.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ttgb.fields...)
}

// TreatmentTypeSelect is the builder for selecting fields of TreatmentType entities.
type TreatmentTypeSelect struct {
	*TreatmentTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TreatmentTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TreatmentTypeQuery.sqlQuery()
	return tts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tts *TreatmentTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tts *TreatmentTypeSelect) StringsX(ctx context.Context) []string {
	v, err := tts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tts *TreatmentTypeSelect) StringX(ctx context.Context) string {
	v, err := tts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tts *TreatmentTypeSelect) IntsX(ctx context.Context) []int {
	v, err := tts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tts *TreatmentTypeSelect) IntX(ctx context.Context) int {
	v, err := tts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tts *TreatmentTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tts *TreatmentTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := tts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TreatmentTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tts *TreatmentTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := tts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tts *TreatmentTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{treatmenttype.Label}
	default:
		err = fmt.Errorf("ent: TreatmentTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tts *TreatmentTypeSelect) BoolX(ctx context.Context) bool {
	v, err := tts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tts *TreatmentTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sqlQuery().Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tts *TreatmentTypeSelect) sqlQuery() sql.Querier {
	selector := tts.sql
	selector.Select(selector.Columns(tts.fields...)...)
	return selector
}
