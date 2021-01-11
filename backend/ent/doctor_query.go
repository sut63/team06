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
	"github.com/team06/app/ent/appointmentresults"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/medicalprocedure"
	"github.com/team06/app/ent/predicate"
)

// DoctorQuery is the builder for querying Doctor entities.
type DoctorQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Doctor
	// eager-loading edges.
	withDoctorToDiagnosis          *DiagnosisQuery
	withDoctorToAppointmentResults *AppointmentResultsQuery
	withDoctorToMedicalProcedure   *MedicalProcedureQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DoctorQuery builder.
func (dq *DoctorQuery) Where(ps ...predicate.Doctor) *DoctorQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DoctorQuery) Limit(limit int) *DoctorQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DoctorQuery) Offset(offset int) *DoctorQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DoctorQuery) Order(o ...OrderFunc) *DoctorQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDoctorToDiagnosis chains the current query on the "DoctorToDiagnosis" edge.
func (dq *DoctorQuery) QueryDoctorToDiagnosis() *DiagnosisQuery {
	query := &DiagnosisQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, selector),
			sqlgraph.To(diagnosis.Table, diagnosis.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.DoctorToDiagnosisTable, doctor.DoctorToDiagnosisColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctorToAppointmentResults chains the current query on the "DoctorToAppointmentResults" edge.
func (dq *DoctorQuery) QueryDoctorToAppointmentResults() *AppointmentResultsQuery {
	query := &AppointmentResultsQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, selector),
			sqlgraph.To(appointmentresults.Table, appointmentresults.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.DoctorToAppointmentResultsTable, doctor.DoctorToAppointmentResultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctorToMedicalProcedure chains the current query on the "DoctorToMedicalProcedure" edge.
func (dq *DoctorQuery) QueryDoctorToMedicalProcedure() *MedicalProcedureQuery {
	query := &MedicalProcedureQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, selector),
			sqlgraph.To(medicalprocedure.Table, medicalprocedure.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.DoctorToMedicalProcedureTable, doctor.DoctorToMedicalProcedureColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Doctor entity from the query.
// Returns a *NotFoundError when no Doctor was found.
func (dq *DoctorQuery) First(ctx context.Context) (*Doctor, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{doctor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DoctorQuery) FirstX(ctx context.Context) *Doctor {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Doctor ID from the query.
// Returns a *NotFoundError when no Doctor ID was found.
func (dq *DoctorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{doctor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DoctorQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Doctor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Doctor entity is not found.
// Returns a *NotFoundError when no Doctor entities are found.
func (dq *DoctorQuery) Only(ctx context.Context) (*Doctor, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{doctor.Label}
	default:
		return nil, &NotSingularError{doctor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DoctorQuery) OnlyX(ctx context.Context) *Doctor {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Doctor ID in the query.
// Returns a *NotSingularError when exactly one Doctor ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dq *DoctorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = &NotSingularError{doctor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DoctorQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Doctors.
func (dq *DoctorQuery) All(ctx context.Context) ([]*Doctor, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DoctorQuery) AllX(ctx context.Context) []*Doctor {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Doctor IDs.
func (dq *DoctorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(doctor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DoctorQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DoctorQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DoctorQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DoctorQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DoctorQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DoctorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DoctorQuery) Clone() *DoctorQuery {
	if dq == nil {
		return nil
	}
	return &DoctorQuery{
		config:                         dq.config,
		limit:                          dq.limit,
		offset:                         dq.offset,
		order:                          append([]OrderFunc{}, dq.order...),
		predicates:                     append([]predicate.Doctor{}, dq.predicates...),
		withDoctorToDiagnosis:          dq.withDoctorToDiagnosis.Clone(),
		withDoctorToAppointmentResults: dq.withDoctorToAppointmentResults.Clone(),
		withDoctorToMedicalProcedure:   dq.withDoctorToMedicalProcedure.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDoctorToDiagnosis tells the query-builder to eager-load the nodes that are connected to
// the "DoctorToDiagnosis" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DoctorQuery) WithDoctorToDiagnosis(opts ...func(*DiagnosisQuery)) *DoctorQuery {
	query := &DiagnosisQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDoctorToDiagnosis = query
	return dq
}

// WithDoctorToAppointmentResults tells the query-builder to eager-load the nodes that are connected to
// the "DoctorToAppointmentResults" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DoctorQuery) WithDoctorToAppointmentResults(opts ...func(*AppointmentResultsQuery)) *DoctorQuery {
	query := &AppointmentResultsQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDoctorToAppointmentResults = query
	return dq
}

// WithDoctorToMedicalProcedure tells the query-builder to eager-load the nodes that are connected to
// the "DoctorToMedicalProcedure" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DoctorQuery) WithDoctorToMedicalProcedure(opts ...func(*MedicalProcedureQuery)) *DoctorQuery {
	query := &MedicalProcedureQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDoctorToMedicalProcedure = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DoctorName string `json:"doctorName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Doctor.Query().
//		GroupBy(doctor.FieldDoctorName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) GroupBy(field string, fields ...string) *DoctorGroupBy {
	group := &DoctorGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DoctorName string `json:"doctorName,omitempty"`
//	}
//
//	client.Doctor.Query().
//		Select(doctor.FieldDoctorName).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) Select(field string, fields ...string) *DoctorSelect {
	dq.fields = append([]string{field}, fields...)
	return &DoctorSelect{DoctorQuery: dq}
}

func (dq *DoctorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !doctor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DoctorQuery) sqlAll(ctx context.Context) ([]*Doctor, error) {
	var (
		nodes       = []*Doctor{}
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withDoctorToDiagnosis != nil,
			dq.withDoctorToAppointmentResults != nil,
			dq.withDoctorToMedicalProcedure != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Doctor{config: dq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withDoctorToDiagnosis; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DoctorToDiagnosis = []*Diagnosis{}
		}
		query.withFKs = true
		query.Where(predicate.Diagnosis(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.DoctorToDiagnosisColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_doctor_to_diagnosis
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_doctor_to_diagnosis" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_doctor_to_diagnosis" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DoctorToDiagnosis = append(node.Edges.DoctorToDiagnosis, n)
		}
	}

	if query := dq.withDoctorToAppointmentResults; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DoctorToAppointmentResults = []*AppointmentResults{}
		}
		query.withFKs = true
		query.Where(predicate.AppointmentResults(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.DoctorToAppointmentResultsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_doctor_to_appointment_results
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_doctor_to_appointment_results" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_doctor_to_appointment_results" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DoctorToAppointmentResults = append(node.Edges.DoctorToAppointmentResults, n)
		}
	}

	if query := dq.withDoctorToMedicalProcedure; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DoctorToMedicalProcedure = []*MedicalProcedure{}
		}
		query.withFKs = true
		query.Where(predicate.MedicalProcedure(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.DoctorToMedicalProcedureColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_doctor_to_medical_procedure
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_doctor_to_medical_procedure" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_doctor_to_medical_procedure" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DoctorToMedicalProcedure = append(node.Edges.DoctorToMedicalProcedure, n)
		}
	}

	return nodes, nil
}

func (dq *DoctorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DoctorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DoctorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, doctor.FieldID)
		for i := range fields {
			if fields[i] != doctor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, doctor.ValidColumn)
			}
		}
	}
	return _spec
}

func (dq *DoctorQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(doctor.Table)
	selector := builder.Select(t1.Columns(doctor.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(doctor.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector, doctor.ValidColumn)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoctorGroupBy is the group-by builder for Doctor entities.
type DoctorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DoctorGroupBy) Aggregate(fns ...AggregateFunc) *DoctorGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DoctorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DoctorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DoctorGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DoctorGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DoctorGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DoctorGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DoctorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DoctorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DoctorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DoctorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DoctorGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DoctorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !doctor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DoctorGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector, doctor.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DoctorSelect is the builder for selecting fields of Doctor entities.
type DoctorSelect struct {
	*DoctorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DoctorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DoctorQuery.sqlQuery()
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DoctorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DoctorSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DoctorSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DoctorSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DoctorSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DoctorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DoctorSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DoctorSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DoctorSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DoctorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DoctorSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
