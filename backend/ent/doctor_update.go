// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/appointmentresults"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/medicalprocedure"
	"github.com/team06/app/ent/predicate"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// Where adds a new predicate for the DoctorUpdate builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetDoctorName sets the "doctorName" field.
func (du *DoctorUpdate) SetDoctorName(s string) *DoctorUpdate {
	du.mutation.SetDoctorName(s)
	return du
}

// SetDoctorUsername sets the "doctorUsername" field.
func (du *DoctorUpdate) SetDoctorUsername(s string) *DoctorUpdate {
	du.mutation.SetDoctorUsername(s)
	return du
}

// SetDoctorPassword sets the "doctorPassword" field.
func (du *DoctorUpdate) SetDoctorPassword(s string) *DoctorUpdate {
	du.mutation.SetDoctorPassword(s)
	return du
}

// AddDoctorToDiagnosiIDs adds the "DoctorToDiagnosis" edge to the Diagnosis entity by IDs.
func (du *DoctorUpdate) AddDoctorToDiagnosiIDs(ids ...int) *DoctorUpdate {
	du.mutation.AddDoctorToDiagnosiIDs(ids...)
	return du
}

// AddDoctorToDiagnosis adds the "DoctorToDiagnosis" edges to the Diagnosis entity.
func (du *DoctorUpdate) AddDoctorToDiagnosis(d ...*Diagnosis) *DoctorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDoctorToDiagnosiIDs(ids...)
}

// AddDoctorToAppointmentResultIDs adds the "DoctorToAppointmentResults" edge to the AppointmentResults entity by IDs.
func (du *DoctorUpdate) AddDoctorToAppointmentResultIDs(ids ...int) *DoctorUpdate {
	du.mutation.AddDoctorToAppointmentResultIDs(ids...)
	return du
}

// AddDoctorToAppointmentResults adds the "DoctorToAppointmentResults" edges to the AppointmentResults entity.
func (du *DoctorUpdate) AddDoctorToAppointmentResults(a ...*AppointmentResults) *DoctorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.AddDoctorToAppointmentResultIDs(ids...)
}

// AddDoctorToMedicalProcedureIDs adds the "DoctorToMedicalProcedure" edge to the MedicalProcedure entity by IDs.
func (du *DoctorUpdate) AddDoctorToMedicalProcedureIDs(ids ...int) *DoctorUpdate {
	du.mutation.AddDoctorToMedicalProcedureIDs(ids...)
	return du
}

// AddDoctorToMedicalProcedure adds the "DoctorToMedicalProcedure" edges to the MedicalProcedure entity.
func (du *DoctorUpdate) AddDoctorToMedicalProcedure(m ...*MedicalProcedure) *DoctorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return du.AddDoctorToMedicalProcedureIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// ClearDoctorToDiagnosis clears all "DoctorToDiagnosis" edges to the Diagnosis entity.
func (du *DoctorUpdate) ClearDoctorToDiagnosis() *DoctorUpdate {
	du.mutation.ClearDoctorToDiagnosis()
	return du
}

// RemoveDoctorToDiagnosiIDs removes the "DoctorToDiagnosis" edge to Diagnosis entities by IDs.
func (du *DoctorUpdate) RemoveDoctorToDiagnosiIDs(ids ...int) *DoctorUpdate {
	du.mutation.RemoveDoctorToDiagnosiIDs(ids...)
	return du
}

// RemoveDoctorToDiagnosis removes "DoctorToDiagnosis" edges to Diagnosis entities.
func (du *DoctorUpdate) RemoveDoctorToDiagnosis(d ...*Diagnosis) *DoctorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDoctorToDiagnosiIDs(ids...)
}

// ClearDoctorToAppointmentResults clears all "DoctorToAppointmentResults" edges to the AppointmentResults entity.
func (du *DoctorUpdate) ClearDoctorToAppointmentResults() *DoctorUpdate {
	du.mutation.ClearDoctorToAppointmentResults()
	return du
}

// RemoveDoctorToAppointmentResultIDs removes the "DoctorToAppointmentResults" edge to AppointmentResults entities by IDs.
func (du *DoctorUpdate) RemoveDoctorToAppointmentResultIDs(ids ...int) *DoctorUpdate {
	du.mutation.RemoveDoctorToAppointmentResultIDs(ids...)
	return du
}

// RemoveDoctorToAppointmentResults removes "DoctorToAppointmentResults" edges to AppointmentResults entities.
func (du *DoctorUpdate) RemoveDoctorToAppointmentResults(a ...*AppointmentResults) *DoctorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.RemoveDoctorToAppointmentResultIDs(ids...)
}

// ClearDoctorToMedicalProcedure clears all "DoctorToMedicalProcedure" edges to the MedicalProcedure entity.
func (du *DoctorUpdate) ClearDoctorToMedicalProcedure() *DoctorUpdate {
	du.mutation.ClearDoctorToMedicalProcedure()
	return du
}

// RemoveDoctorToMedicalProcedureIDs removes the "DoctorToMedicalProcedure" edge to MedicalProcedure entities by IDs.
func (du *DoctorUpdate) RemoveDoctorToMedicalProcedureIDs(ids ...int) *DoctorUpdate {
	du.mutation.RemoveDoctorToMedicalProcedureIDs(ids...)
	return du
}

// RemoveDoctorToMedicalProcedure removes "DoctorToMedicalProcedure" edges to MedicalProcedure entities.
func (du *DoctorUpdate) RemoveDoctorToMedicalProcedure(m ...*MedicalProcedure) *DoctorUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return du.RemoveDoctorToMedicalProcedureIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DoctorUpdate) check() error {
	if v, ok := du.mutation.DoctorName(); ok {
		if err := doctor.DoctorNameValidator(v); err != nil {
			return &ValidationError{Name: "doctorName", err: fmt.Errorf("ent: validator failed for field \"doctorName\": %w", err)}
		}
	}
	if v, ok := du.mutation.DoctorUsername(); ok {
		if err := doctor.DoctorUsernameValidator(v); err != nil {
			return &ValidationError{Name: "doctorUsername", err: fmt.Errorf("ent: validator failed for field \"doctorUsername\": %w", err)}
		}
	}
	if v, ok := du.mutation.DoctorPassword(); ok {
		if err := doctor.DoctorPasswordValidator(v); err != nil {
			return &ValidationError{Name: "doctorPassword", err: fmt.Errorf("ent: validator failed for field \"doctorPassword\": %w", err)}
		}
	}
	return nil
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DoctorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorName,
		})
	}
	if value, ok := du.mutation.DoctorUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorUsername,
		})
	}
	if value, ok := du.mutation.DoctorPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorPassword,
		})
	}
	if du.mutation.DoctorToDiagnosisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDoctorToDiagnosisIDs(); len(nodes) > 0 && !du.mutation.DoctorToDiagnosisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DoctorToDiagnosisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DoctorToAppointmentResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDoctorToAppointmentResultsIDs(); len(nodes) > 0 && !du.mutation.DoctorToAppointmentResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DoctorToAppointmentResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DoctorToMedicalProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDoctorToMedicalProcedureIDs(); len(nodes) > 0 && !du.mutation.DoctorToMedicalProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DoctorToMedicalProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// SetDoctorName sets the "doctorName" field.
func (duo *DoctorUpdateOne) SetDoctorName(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorName(s)
	return duo
}

// SetDoctorUsername sets the "doctorUsername" field.
func (duo *DoctorUpdateOne) SetDoctorUsername(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorUsername(s)
	return duo
}

// SetDoctorPassword sets the "doctorPassword" field.
func (duo *DoctorUpdateOne) SetDoctorPassword(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorPassword(s)
	return duo
}

// AddDoctorToDiagnosiIDs adds the "DoctorToDiagnosis" edge to the Diagnosis entity by IDs.
func (duo *DoctorUpdateOne) AddDoctorToDiagnosiIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.AddDoctorToDiagnosiIDs(ids...)
	return duo
}

// AddDoctorToDiagnosis adds the "DoctorToDiagnosis" edges to the Diagnosis entity.
func (duo *DoctorUpdateOne) AddDoctorToDiagnosis(d ...*Diagnosis) *DoctorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDoctorToDiagnosiIDs(ids...)
}

// AddDoctorToAppointmentResultIDs adds the "DoctorToAppointmentResults" edge to the AppointmentResults entity by IDs.
func (duo *DoctorUpdateOne) AddDoctorToAppointmentResultIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.AddDoctorToAppointmentResultIDs(ids...)
	return duo
}

// AddDoctorToAppointmentResults adds the "DoctorToAppointmentResults" edges to the AppointmentResults entity.
func (duo *DoctorUpdateOne) AddDoctorToAppointmentResults(a ...*AppointmentResults) *DoctorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.AddDoctorToAppointmentResultIDs(ids...)
}

// AddDoctorToMedicalProcedureIDs adds the "DoctorToMedicalProcedure" edge to the MedicalProcedure entity by IDs.
func (duo *DoctorUpdateOne) AddDoctorToMedicalProcedureIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.AddDoctorToMedicalProcedureIDs(ids...)
	return duo
}

// AddDoctorToMedicalProcedure adds the "DoctorToMedicalProcedure" edges to the MedicalProcedure entity.
func (duo *DoctorUpdateOne) AddDoctorToMedicalProcedure(m ...*MedicalProcedure) *DoctorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duo.AddDoctorToMedicalProcedureIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// ClearDoctorToDiagnosis clears all "DoctorToDiagnosis" edges to the Diagnosis entity.
func (duo *DoctorUpdateOne) ClearDoctorToDiagnosis() *DoctorUpdateOne {
	duo.mutation.ClearDoctorToDiagnosis()
	return duo
}

// RemoveDoctorToDiagnosiIDs removes the "DoctorToDiagnosis" edge to Diagnosis entities by IDs.
func (duo *DoctorUpdateOne) RemoveDoctorToDiagnosiIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.RemoveDoctorToDiagnosiIDs(ids...)
	return duo
}

// RemoveDoctorToDiagnosis removes "DoctorToDiagnosis" edges to Diagnosis entities.
func (duo *DoctorUpdateOne) RemoveDoctorToDiagnosis(d ...*Diagnosis) *DoctorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDoctorToDiagnosiIDs(ids...)
}

// ClearDoctorToAppointmentResults clears all "DoctorToAppointmentResults" edges to the AppointmentResults entity.
func (duo *DoctorUpdateOne) ClearDoctorToAppointmentResults() *DoctorUpdateOne {
	duo.mutation.ClearDoctorToAppointmentResults()
	return duo
}

// RemoveDoctorToAppointmentResultIDs removes the "DoctorToAppointmentResults" edge to AppointmentResults entities by IDs.
func (duo *DoctorUpdateOne) RemoveDoctorToAppointmentResultIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.RemoveDoctorToAppointmentResultIDs(ids...)
	return duo
}

// RemoveDoctorToAppointmentResults removes "DoctorToAppointmentResults" edges to AppointmentResults entities.
func (duo *DoctorUpdateOne) RemoveDoctorToAppointmentResults(a ...*AppointmentResults) *DoctorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.RemoveDoctorToAppointmentResultIDs(ids...)
}

// ClearDoctorToMedicalProcedure clears all "DoctorToMedicalProcedure" edges to the MedicalProcedure entity.
func (duo *DoctorUpdateOne) ClearDoctorToMedicalProcedure() *DoctorUpdateOne {
	duo.mutation.ClearDoctorToMedicalProcedure()
	return duo
}

// RemoveDoctorToMedicalProcedureIDs removes the "DoctorToMedicalProcedure" edge to MedicalProcedure entities by IDs.
func (duo *DoctorUpdateOne) RemoveDoctorToMedicalProcedureIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.RemoveDoctorToMedicalProcedureIDs(ids...)
	return duo
}

// RemoveDoctorToMedicalProcedure removes "DoctorToMedicalProcedure" edges to MedicalProcedure entities.
func (duo *DoctorUpdateOne) RemoveDoctorToMedicalProcedure(m ...*MedicalProcedure) *DoctorUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duo.RemoveDoctorToMedicalProcedureIDs(ids...)
}

// Save executes the query and returns the updated Doctor entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	var (
		err  error
		node *Doctor
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DoctorUpdateOne) check() error {
	if v, ok := duo.mutation.DoctorName(); ok {
		if err := doctor.DoctorNameValidator(v); err != nil {
			return &ValidationError{Name: "doctorName", err: fmt.Errorf("ent: validator failed for field \"doctorName\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DoctorUsername(); ok {
		if err := doctor.DoctorUsernameValidator(v); err != nil {
			return &ValidationError{Name: "doctorUsername", err: fmt.Errorf("ent: validator failed for field \"doctorUsername\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DoctorPassword(); ok {
		if err := doctor.DoctorPasswordValidator(v); err != nil {
			return &ValidationError{Name: "doctorPassword", err: fmt.Errorf("ent: validator failed for field \"doctorPassword\": %w", err)}
		}
	}
	return nil
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (_node *Doctor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doctor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DoctorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorName,
		})
	}
	if value, ok := duo.mutation.DoctorUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorUsername,
		})
	}
	if value, ok := duo.mutation.DoctorPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorPassword,
		})
	}
	if duo.mutation.DoctorToDiagnosisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDoctorToDiagnosisIDs(); len(nodes) > 0 && !duo.mutation.DoctorToDiagnosisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DoctorToDiagnosisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToDiagnosisTable,
			Columns: []string{doctor.DoctorToDiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DoctorToAppointmentResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDoctorToAppointmentResultsIDs(); len(nodes) > 0 && !duo.mutation.DoctorToAppointmentResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DoctorToAppointmentResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToAppointmentResultsTable,
			Columns: []string{doctor.DoctorToAppointmentResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointmentresults.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DoctorToMedicalProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDoctorToMedicalProcedureIDs(); len(nodes) > 0 && !duo.mutation.DoctorToMedicalProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DoctorToMedicalProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorToMedicalProcedureTable,
			Columns: []string{doctor.DoctorToMedicalProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalprocedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Doctor{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
