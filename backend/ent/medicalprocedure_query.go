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
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/medicalprocedure"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/proceduretype"
)

// MedicalProcedureQuery is the builder for querying MedicalProcedure entities.
type MedicalProcedureQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.MedicalProcedure
	// eager-loading edges.
	withPatient       *PatientQuery
	withProcedureType *ProcedureTypeQuery
	withDoctor        *DoctorQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MedicalProcedureQuery builder.
func (mpq *MedicalProcedureQuery) Where(ps ...predicate.MedicalProcedure) *MedicalProcedureQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit adds a limit step to the query.
func (mpq *MedicalProcedureQuery) Limit(limit int) *MedicalProcedureQuery {
	mpq.limit = &limit
	return mpq
}

// Offset adds an offset step to the query.
func (mpq *MedicalProcedureQuery) Offset(offset int) *MedicalProcedureQuery {
	mpq.offset = &offset
	return mpq
}

// Order adds an order step to the query.
func (mpq *MedicalProcedureQuery) Order(o ...OrderFunc) *MedicalProcedureQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryPatient chains the current query on the "Patient" edge.
func (mpq *MedicalProcedureQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: mpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalprocedure.Table, medicalprocedure.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, medicalprocedure.PatientTable, medicalprocedure.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProcedureType chains the current query on the "ProcedureType" edge.
func (mpq *MedicalProcedureQuery) QueryProcedureType() *ProcedureTypeQuery {
	query := &ProcedureTypeQuery{config: mpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalprocedure.Table, medicalprocedure.FieldID, selector),
			sqlgraph.To(proceduretype.Table, proceduretype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, medicalprocedure.ProcedureTypeTable, medicalprocedure.ProcedureTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctor chains the current query on the "Doctor" edge.
func (mpq *MedicalProcedureQuery) QueryDoctor() *DoctorQuery {
	query := &DoctorQuery{config: mpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalprocedure.Table, medicalprocedure.FieldID, selector),
			sqlgraph.To(doctor.Table, doctor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, medicalprocedure.DoctorTable, medicalprocedure.DoctorColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MedicalProcedure entity from the query.
// Returns a *NotFoundError when no MedicalProcedure was found.
func (mpq *MedicalProcedureQuery) First(ctx context.Context) (*MedicalProcedure, error) {
	nodes, err := mpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{medicalprocedure.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) FirstX(ctx context.Context) *MedicalProcedure {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MedicalProcedure ID from the query.
// Returns a *NotFoundError when no MedicalProcedure ID was found.
func (mpq *MedicalProcedureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medicalprocedure.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) FirstIDX(ctx context.Context) int {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MedicalProcedure entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MedicalProcedure entity is not found.
// Returns a *NotFoundError when no MedicalProcedure entities are found.
func (mpq *MedicalProcedureQuery) Only(ctx context.Context) (*MedicalProcedure, error) {
	nodes, err := mpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{medicalprocedure.Label}
	default:
		return nil, &NotSingularError{medicalprocedure.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) OnlyX(ctx context.Context) *MedicalProcedure {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MedicalProcedure ID in the query.
// Returns a *NotSingularError when exactly one MedicalProcedure ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MedicalProcedureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = &NotSingularError{medicalprocedure.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) OnlyIDX(ctx context.Context) int {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MedicalProcedures.
func (mpq *MedicalProcedureQuery) All(ctx context.Context) ([]*MedicalProcedure, error) {
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) AllX(ctx context.Context) []*MedicalProcedure {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MedicalProcedure IDs.
func (mpq *MedicalProcedureQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mpq.Select(medicalprocedure.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) IDsX(ctx context.Context) []int {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MedicalProcedureQuery) Count(ctx context.Context) (int, error) {
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MedicalProcedureQuery) Exist(ctx context.Context) (bool, error) {
	if err := mpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *MedicalProcedureQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MedicalProcedureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MedicalProcedureQuery) Clone() *MedicalProcedureQuery {
	if mpq == nil {
		return nil
	}
	return &MedicalProcedureQuery{
		config:            mpq.config,
		limit:             mpq.limit,
		offset:            mpq.offset,
		order:             append([]OrderFunc{}, mpq.order...),
		predicates:        append([]predicate.MedicalProcedure{}, mpq.predicates...),
		withPatient:       mpq.withPatient.Clone(),
		withProcedureType: mpq.withProcedureType.Clone(),
		withDoctor:        mpq.withDoctor.Clone(),
		// clone intermediate query.
		sql:  mpq.sql.Clone(),
		path: mpq.path,
	}
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "Patient" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MedicalProcedureQuery) WithPatient(opts ...func(*PatientQuery)) *MedicalProcedureQuery {
	query := &PatientQuery{config: mpq.config}
	for _, opt := range opts {
		opt(query)
	}
	mpq.withPatient = query
	return mpq
}

// WithProcedureType tells the query-builder to eager-load the nodes that are connected to
// the "ProcedureType" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MedicalProcedureQuery) WithProcedureType(opts ...func(*ProcedureTypeQuery)) *MedicalProcedureQuery {
	query := &ProcedureTypeQuery{config: mpq.config}
	for _, opt := range opts {
		opt(query)
	}
	mpq.withProcedureType = query
	return mpq
}

// WithDoctor tells the query-builder to eager-load the nodes that are connected to
// the "Doctor" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MedicalProcedureQuery) WithDoctor(opts ...func(*DoctorQuery)) *MedicalProcedureQuery {
	query := &DoctorQuery{config: mpq.config}
	for _, opt := range opts {
		opt(query)
	}
	mpq.withDoctor = query
	return mpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProcedureOrder string `json:"procedureOrder,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MedicalProcedure.Query().
//		GroupBy(medicalprocedure.FieldProcedureOrder).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mpq *MedicalProcedureQuery) GroupBy(field string, fields ...string) *MedicalProcedureGroupBy {
	group := &MedicalProcedureGroupBy{config: mpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mpq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProcedureOrder string `json:"procedureOrder,omitempty"`
//	}
//
//	client.MedicalProcedure.Query().
//		Select(medicalprocedure.FieldProcedureOrder).
//		Scan(ctx, &v)
//
func (mpq *MedicalProcedureQuery) Select(field string, fields ...string) *MedicalProcedureSelect {
	mpq.fields = append([]string{field}, fields...)
	return &MedicalProcedureSelect{MedicalProcedureQuery: mpq}
}

func (mpq *MedicalProcedureQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mpq.fields {
		if !medicalprocedure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *MedicalProcedureQuery) sqlAll(ctx context.Context) ([]*MedicalProcedure, error) {
	var (
		nodes       = []*MedicalProcedure{}
		withFKs     = mpq.withFKs
		_spec       = mpq.querySpec()
		loadedTypes = [3]bool{
			mpq.withPatient != nil,
			mpq.withProcedureType != nil,
			mpq.withDoctor != nil,
		}
	)
	if mpq.withPatient != nil || mpq.withProcedureType != nil || mpq.withDoctor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, medicalprocedure.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MedicalProcedure{config: mpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mpq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MedicalProcedure)
		for i := range nodes {
			if fk := nodes[i].patient_patient_to_medical_procedure; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "patient_patient_to_medical_procedure" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := mpq.withProcedureType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MedicalProcedure)
		for i := range nodes {
			if fk := nodes[i].procedure_type_procedure_to_medical_procedure; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(proceduretype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "procedure_type_procedure_to_medical_procedure" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ProcedureType = n
			}
		}
	}

	if query := mpq.withDoctor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MedicalProcedure)
		for i := range nodes {
			if fk := nodes[i].doctor_doctor_to_medical_procedure; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(doctor.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_doctor_to_medical_procedure" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Doctor = n
			}
		}
	}

	return nodes, nil
}

func (mpq *MedicalProcedureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MedicalProcedureQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mpq *MedicalProcedureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalprocedure.Table,
			Columns: medicalprocedure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalprocedure.FieldID,
			},
		},
		From:   mpq.sql,
		Unique: true,
	}
	if fields := mpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, medicalprocedure.FieldID)
		for i := range fields {
			if fields[i] != medicalprocedure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, medicalprocedure.ValidColumn)
			}
		}
	}
	return _spec
}

func (mpq *MedicalProcedureQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(medicalprocedure.Table)
	selector := builder.Select(t1.Columns(medicalprocedure.Columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(medicalprocedure.Columns...)...)
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector, medicalprocedure.ValidColumn)
	}
	if offset := mpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalProcedureGroupBy is the group-by builder for MedicalProcedure entities.
type MedicalProcedureGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MedicalProcedureGroupBy) Aggregate(fns ...AggregateFunc) *MedicalProcedureGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mpgb *MedicalProcedureGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mpgb.path(ctx)
	if err != nil {
		return err
	}
	mpgb.sql = query
	return mpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mpgb.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) StringsX(ctx context.Context) []string {
	v, err := mpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) StringX(ctx context.Context) string {
	v, err := mpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mpgb.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) IntsX(ctx context.Context) []int {
	v, err := mpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) IntX(ctx context.Context) int {
	v, err := mpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mpgb.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mpgb.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mpgb *MedicalProcedureGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mpgb *MedicalProcedureGroupBy) BoolX(ctx context.Context) bool {
	v, err := mpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mpgb *MedicalProcedureGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mpgb.fields {
		if !medicalprocedure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mpgb *MedicalProcedureGroupBy) sqlQuery() *sql.Selector {
	selector := mpgb.sql
	columns := make([]string, 0, len(mpgb.fields)+len(mpgb.fns))
	columns = append(columns, mpgb.fields...)
	for _, fn := range mpgb.fns {
		columns = append(columns, fn(selector, medicalprocedure.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(mpgb.fields...)
}

// MedicalProcedureSelect is the builder for selecting fields of MedicalProcedure entities.
type MedicalProcedureSelect struct {
	*MedicalProcedureQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MedicalProcedureSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	mps.sql = mps.MedicalProcedureQuery.sqlQuery()
	return mps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mps *MedicalProcedureSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mps.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mps *MedicalProcedureSelect) StringsX(ctx context.Context) []string {
	v, err := mps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mps *MedicalProcedureSelect) StringX(ctx context.Context) string {
	v, err := mps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mps.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mps *MedicalProcedureSelect) IntsX(ctx context.Context) []int {
	v, err := mps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mps *MedicalProcedureSelect) IntX(ctx context.Context) int {
	v, err := mps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mps.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mps *MedicalProcedureSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mps *MedicalProcedureSelect) Float64X(ctx context.Context) float64 {
	v, err := mps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mps.fields) > 1 {
		return nil, errors.New("ent: MedicalProcedureSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mps *MedicalProcedureSelect) BoolsX(ctx context.Context) []bool {
	v, err := mps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mps *MedicalProcedureSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalprocedure.Label}
	default:
		err = fmt.Errorf("ent: MedicalProcedureSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mps *MedicalProcedureSelect) BoolX(ctx context.Context) bool {
	v, err := mps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mps *MedicalProcedureSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mps.sqlQuery().Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mps *MedicalProcedureSelect) sqlQuery() sql.Querier {
	selector := mps.sql
	selector.Select(selector.Columns(mps.fields...)...)
	return selector
}
