// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/hospital"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/righttotreatment"
	"github.com/team06/app/ent/righttotreatmenttype"
)

// RightToTreatmentQuery is the builder for querying RightToTreatment entities.
type RightToTreatmentQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.RightToTreatment
	// eager-loading edges.
	withHospital             *HospitalQuery
	withRightToTreatmentType *RightToTreatmentTypeQuery
	withPatient              *PatientQuery
	withFKs                  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RightToTreatmentQuery builder.
func (rttq *RightToTreatmentQuery) Where(ps ...predicate.RightToTreatment) *RightToTreatmentQuery {
	rttq.predicates = append(rttq.predicates, ps...)
	return rttq
}

// Limit adds a limit step to the query.
func (rttq *RightToTreatmentQuery) Limit(limit int) *RightToTreatmentQuery {
	rttq.limit = &limit
	return rttq
}

// Offset adds an offset step to the query.
func (rttq *RightToTreatmentQuery) Offset(offset int) *RightToTreatmentQuery {
	rttq.offset = &offset
	return rttq
}

// Order adds an order step to the query.
func (rttq *RightToTreatmentQuery) Order(o ...OrderFunc) *RightToTreatmentQuery {
	rttq.order = append(rttq.order, o...)
	return rttq
}

// QueryHospital chains the current query on the "Hospital" edge.
func (rttq *RightToTreatmentQuery) QueryHospital() *HospitalQuery {
	query := &HospitalQuery{config: rttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rttq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(righttotreatment.Table, righttotreatment.FieldID, selector),
			sqlgraph.To(hospital.Table, hospital.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, righttotreatment.HospitalTable, righttotreatment.HospitalColumn),
		)
		fromU = sqlgraph.SetNeighbors(rttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRightToTreatmentType chains the current query on the "RightToTreatmentType" edge.
func (rttq *RightToTreatmentQuery) QueryRightToTreatmentType() *RightToTreatmentTypeQuery {
	query := &RightToTreatmentTypeQuery{config: rttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rttq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(righttotreatment.Table, righttotreatment.FieldID, selector),
			sqlgraph.To(righttotreatmenttype.Table, righttotreatmenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, righttotreatment.RightToTreatmentTypeTable, righttotreatment.RightToTreatmentTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the "Patient" edge.
func (rttq *RightToTreatmentQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: rttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rttq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(righttotreatment.Table, righttotreatment.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, righttotreatment.PatientTable, righttotreatment.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(rttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RightToTreatment entity from the query.
// Returns a *NotFoundError when no RightToTreatment was found.
func (rttq *RightToTreatmentQuery) First(ctx context.Context) (*RightToTreatment, error) {
	nodes, err := rttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{righttotreatment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) FirstX(ctx context.Context) *RightToTreatment {
	node, err := rttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RightToTreatment ID from the query.
// Returns a *NotFoundError when no RightToTreatment ID was found.
func (rttq *RightToTreatmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{righttotreatment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) FirstIDX(ctx context.Context) int {
	id, err := rttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RightToTreatment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RightToTreatment entity is not found.
// Returns a *NotFoundError when no RightToTreatment entities are found.
func (rttq *RightToTreatmentQuery) Only(ctx context.Context) (*RightToTreatment, error) {
	nodes, err := rttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{righttotreatment.Label}
	default:
		return nil, &NotSingularError{righttotreatment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) OnlyX(ctx context.Context) *RightToTreatment {
	node, err := rttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RightToTreatment ID in the query.
// Returns a *NotSingularError when exactly one RightToTreatment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rttq *RightToTreatmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = &NotSingularError{righttotreatment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := rttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RightToTreatments.
func (rttq *RightToTreatmentQuery) All(ctx context.Context) ([]*RightToTreatment, error) {
	if err := rttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) AllX(ctx context.Context) []*RightToTreatment {
	nodes, err := rttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RightToTreatment IDs.
func (rttq *RightToTreatmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rttq.Select(righttotreatment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) IDsX(ctx context.Context) []int {
	ids, err := rttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rttq *RightToTreatmentQuery) Count(ctx context.Context) (int, error) {
	if err := rttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) CountX(ctx context.Context) int {
	count, err := rttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rttq *RightToTreatmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := rttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rttq *RightToTreatmentQuery) ExistX(ctx context.Context) bool {
	exist, err := rttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RightToTreatmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rttq *RightToTreatmentQuery) Clone() *RightToTreatmentQuery {
	if rttq == nil {
		return nil
	}
	return &RightToTreatmentQuery{
		config:                   rttq.config,
		limit:                    rttq.limit,
		offset:                   rttq.offset,
		order:                    append([]OrderFunc{}, rttq.order...),
		predicates:               append([]predicate.RightToTreatment{}, rttq.predicates...),
		withHospital:             rttq.withHospital.Clone(),
		withRightToTreatmentType: rttq.withRightToTreatmentType.Clone(),
		withPatient:              rttq.withPatient.Clone(),
		// clone intermediate query.
		sql:  rttq.sql.Clone(),
		path: rttq.path,
	}
}

// WithHospital tells the query-builder to eager-load the nodes that are connected to
// the "Hospital" edge. The optional arguments are used to configure the query builder of the edge.
func (rttq *RightToTreatmentQuery) WithHospital(opts ...func(*HospitalQuery)) *RightToTreatmentQuery {
	query := &HospitalQuery{config: rttq.config}
	for _, opt := range opts {
		opt(query)
	}
	rttq.withHospital = query
	return rttq
}

// WithRightToTreatmentType tells the query-builder to eager-load the nodes that are connected to
// the "RightToTreatmentType" edge. The optional arguments are used to configure the query builder of the edge.
func (rttq *RightToTreatmentQuery) WithRightToTreatmentType(opts ...func(*RightToTreatmentTypeQuery)) *RightToTreatmentQuery {
	query := &RightToTreatmentTypeQuery{config: rttq.config}
	for _, opt := range opts {
		opt(query)
	}
	rttq.withRightToTreatmentType = query
	return rttq
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "Patient" edge. The optional arguments are used to configure the query builder of the edge.
func (rttq *RightToTreatmentQuery) WithPatient(opts ...func(*PatientQuery)) *RightToTreatmentQuery {
	query := &PatientQuery{config: rttq.config}
	for _, opt := range opts {
		opt(query)
	}
	rttq.withPatient = query
	return rttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartTime time.Time `json:"StartTime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RightToTreatment.Query().
//		GroupBy(righttotreatment.FieldStartTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rttq *RightToTreatmentQuery) GroupBy(field string, fields ...string) *RightToTreatmentGroupBy {
	group := &RightToTreatmentGroupBy{config: rttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rttq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StartTime time.Time `json:"StartTime,omitempty"`
//	}
//
//	client.RightToTreatment.Query().
//		Select(righttotreatment.FieldStartTime).
//		Scan(ctx, &v)
//
func (rttq *RightToTreatmentQuery) Select(field string, fields ...string) *RightToTreatmentSelect {
	rttq.fields = append([]string{field}, fields...)
	return &RightToTreatmentSelect{RightToTreatmentQuery: rttq}
}

func (rttq *RightToTreatmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rttq.fields {
		if !righttotreatment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rttq.path != nil {
		prev, err := rttq.path(ctx)
		if err != nil {
			return err
		}
		rttq.sql = prev
	}
	return nil
}

func (rttq *RightToTreatmentQuery) sqlAll(ctx context.Context) ([]*RightToTreatment, error) {
	var (
		nodes       = []*RightToTreatment{}
		withFKs     = rttq.withFKs
		_spec       = rttq.querySpec()
		loadedTypes = [3]bool{
			rttq.withHospital != nil,
			rttq.withRightToTreatmentType != nil,
			rttq.withPatient != nil,
		}
	)
	if rttq.withHospital != nil || rttq.withRightToTreatmentType != nil || rttq.withPatient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, righttotreatment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RightToTreatment{config: rttq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rttq.withHospital; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RightToTreatment)
		for i := range nodes {
			if fk := nodes[i].hospital_hospital; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(hospital.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "hospital_hospital" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Hospital = n
			}
		}
	}

	if query := rttq.withRightToTreatmentType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RightToTreatment)
		for i := range nodes {
			if fk := nodes[i].right_to_treatment_type_type; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(righttotreatmenttype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "right_to_treatment_type_type" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.RightToTreatmentType = n
			}
		}
	}

	if query := rttq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RightToTreatment)
		for i := range nodes {
			if fk := nodes[i].patient_patient_to_right_to_treatment; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_patient_to_right_to_treatment" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	return nodes, nil
}

func (rttq *RightToTreatmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rttq.querySpec()
	return sqlgraph.CountNodes(ctx, rttq.driver, _spec)
}

func (rttq *RightToTreatmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rttq *RightToTreatmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   righttotreatment.Table,
			Columns: righttotreatment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: righttotreatment.FieldID,
			},
		},
		From:   rttq.sql,
		Unique: true,
	}
	if fields := rttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, righttotreatment.FieldID)
		for i := range fields {
			if fields[i] != righttotreatment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, righttotreatment.ValidColumn)
			}
		}
	}
	return _spec
}

func (rttq *RightToTreatmentQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rttq.driver.Dialect())
	t1 := builder.Table(righttotreatment.Table)
	selector := builder.Select(t1.Columns(righttotreatment.Columns...)...).From(t1)
	if rttq.sql != nil {
		selector = rttq.sql
		selector.Select(selector.Columns(righttotreatment.Columns...)...)
	}
	for _, p := range rttq.predicates {
		p(selector)
	}
	for _, p := range rttq.order {
		p(selector, righttotreatment.ValidColumn)
	}
	if offset := rttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RightToTreatmentGroupBy is the group-by builder for RightToTreatment entities.
type RightToTreatmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rttgb *RightToTreatmentGroupBy) Aggregate(fns ...AggregateFunc) *RightToTreatmentGroupBy {
	rttgb.fns = append(rttgb.fns, fns...)
	return rttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rttgb *RightToTreatmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rttgb.path(ctx)
	if err != nil {
		return err
	}
	rttgb.sql = query
	return rttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rttgb.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := rttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) StringX(ctx context.Context) string {
	v, err := rttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rttgb.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := rttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) IntX(ctx context.Context) int {
	v, err := rttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rttgb.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rttgb.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rttgb *RightToTreatmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rttgb *RightToTreatmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := rttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rttgb *RightToTreatmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rttgb.fields {
		if !righttotreatment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rttgb *RightToTreatmentGroupBy) sqlQuery() *sql.Selector {
	selector := rttgb.sql
	columns := make([]string, 0, len(rttgb.fields)+len(rttgb.fns))
	columns = append(columns, rttgb.fields...)
	for _, fn := range rttgb.fns {
		columns = append(columns, fn(selector, righttotreatment.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(rttgb.fields...)
}

// RightToTreatmentSelect is the builder for selecting fields of RightToTreatment entities.
type RightToTreatmentSelect struct {
	*RightToTreatmentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rtts *RightToTreatmentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rtts.prepareQuery(ctx); err != nil {
		return err
	}
	rtts.sql = rtts.RightToTreatmentQuery.sqlQuery()
	return rtts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rtts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rtts.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) StringsX(ctx context.Context) []string {
	v, err := rtts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rtts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) StringX(ctx context.Context) string {
	v, err := rtts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rtts.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) IntsX(ctx context.Context) []int {
	v, err := rtts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rtts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) IntX(ctx context.Context) int {
	v, err := rtts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rtts.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rtts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rtts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) Float64X(ctx context.Context) float64 {
	v, err := rtts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rtts.fields) > 1 {
		return nil, errors.New("ent: RightToTreatmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := rtts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rtts *RightToTreatmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rtts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{righttotreatment.Label}
	default:
		err = fmt.Errorf("ent: RightToTreatmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rtts *RightToTreatmentSelect) BoolX(ctx context.Context) bool {
	v, err := rtts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rtts *RightToTreatmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rtts.sqlQuery().Query()
	if err := rtts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rtts *RightToTreatmentSelect) sqlQuery() sql.Querier {
	selector := rtts.sql
	selector.Select(selector.Columns(rtts.fields...)...)
	return selector
}
