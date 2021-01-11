// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/appointmentresults"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/medicalprocedure"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/prefix"
	"github.com/team06/app/ent/righttotreatment"
	"github.com/team06/app/ent/triageresult"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetPersonalID sets the "personalID" field.
func (pc *PatientCreate) SetPersonalID(i int) *PatientCreate {
	pc.mutation.SetPersonalID(i)
	return pc
}

// SetPatientName sets the "patientName" field.
func (pc *PatientCreate) SetPatientName(s string) *PatientCreate {
	pc.mutation.SetPatientName(s)
	return pc
}

// SetAge sets the "age" field.
func (pc *PatientCreate) SetAge(i int) *PatientCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetHospitalNumber sets the "hospitalNumber" field.
func (pc *PatientCreate) SetHospitalNumber(s string) *PatientCreate {
	pc.mutation.SetHospitalNumber(s)
	return pc
}

// SetDrugAllergy sets the "drugAllergy" field.
func (pc *PatientCreate) SetDrugAllergy(s string) *PatientCreate {
	pc.mutation.SetDrugAllergy(s)
	return pc
}

// SetAddedTime sets the "added_time" field.
func (pc *PatientCreate) SetAddedTime(t time.Time) *PatientCreate {
	pc.mutation.SetAddedTime(t)
	return pc
}

// SetPrefixID sets the "Prefix" edge to the Prefix entity by ID.
func (pc *PatientCreate) SetPrefixID(id int) *PatientCreate {
	pc.mutation.SetPrefixID(id)
	return pc
}

// SetNillablePrefixID sets the "Prefix" edge to the Prefix entity by ID if the given value is not nil.
func (pc *PatientCreate) SetNillablePrefixID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetPrefixID(*id)
	}
	return pc
}

// SetPrefix sets the "Prefix" edge to the Prefix entity.
func (pc *PatientCreate) SetPrefix(p *Prefix) *PatientCreate {
	return pc.SetPrefixID(p.ID)
}

// SetGenderID sets the "Gender" edge to the Gender entity by ID.
func (pc *PatientCreate) SetGenderID(id int) *PatientCreate {
	pc.mutation.SetGenderID(id)
	return pc
}

// SetNillableGenderID sets the "Gender" edge to the Gender entity by ID if the given value is not nil.
func (pc *PatientCreate) SetNillableGenderID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetGenderID(*id)
	}
	return pc
}

// SetGender sets the "Gender" edge to the Gender entity.
func (pc *PatientCreate) SetGender(g *Gender) *PatientCreate {
	return pc.SetGenderID(g.ID)
}

// SetBloodtypeID sets the "Bloodtype" edge to the BloodType entity by ID.
func (pc *PatientCreate) SetBloodtypeID(id int) *PatientCreate {
	pc.mutation.SetBloodtypeID(id)
	return pc
}

// SetNillableBloodtypeID sets the "Bloodtype" edge to the BloodType entity by ID if the given value is not nil.
func (pc *PatientCreate) SetNillableBloodtypeID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetBloodtypeID(*id)
	}
	return pc
}

// SetBloodtype sets the "Bloodtype" edge to the BloodType entity.
func (pc *PatientCreate) SetBloodtype(b *BloodType) *PatientCreate {
	return pc.SetBloodtypeID(b.ID)
}

// AddPatientToTriageResultIDs adds the "patientToTriageResult" edge to the TriageResult entity by IDs.
func (pc *PatientCreate) AddPatientToTriageResultIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientToTriageResultIDs(ids...)
	return pc
}

// AddPatientToTriageResult adds the "patientToTriageResult" edges to the TriageResult entity.
func (pc *PatientCreate) AddPatientToTriageResult(t ...*TriageResult) *PatientCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddPatientToTriageResultIDs(ids...)
}

// AddPatientToAppointmentResultIDs adds the "PatientToAppointmentResults" edge to the AppointmentResults entity by IDs.
func (pc *PatientCreate) AddPatientToAppointmentResultIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientToAppointmentResultIDs(ids...)
	return pc
}

// AddPatientToAppointmentResults adds the "PatientToAppointmentResults" edges to the AppointmentResults entity.
func (pc *PatientCreate) AddPatientToAppointmentResults(a ...*AppointmentResults) *PatientCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddPatientToAppointmentResultIDs(ids...)
}

// AddPatientToMedicalProcedureIDs adds the "PatientToMedicalProcedure" edge to the MedicalProcedure entity by IDs.
func (pc *PatientCreate) AddPatientToMedicalProcedureIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientToMedicalProcedureIDs(ids...)
	return pc
}

// AddPatientToMedicalProcedure adds the "PatientToMedicalProcedure" edges to the MedicalProcedure entity.
func (pc *PatientCreate) AddPatientToMedicalProcedure(m ...*MedicalProcedure) *PatientCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddPatientToMedicalProcedureIDs(ids...)
}

// AddPatientToRightToTreatmentIDs adds the "PatientToRightToTreatment" edge to the RightToTreatment entity by IDs.
func (pc *PatientCreate) AddPatientToRightToTreatmentIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientToRightToTreatmentIDs(ids...)
	return pc
}

// AddPatientToRightToTreatment adds the "PatientToRightToTreatment" edges to the RightToTreatment entity.
func (pc *PatientCreate) AddPatientToRightToTreatment(r ...*RightToTreatment) *PatientCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddPatientToRightToTreatmentIDs(ids...)
}

// AddPatientToDiagnosiIDs adds the "PatientToDiagnosis" edge to the Diagnosis entity by IDs.
func (pc *PatientCreate) AddPatientToDiagnosiIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientToDiagnosiIDs(ids...)
	return pc
}

// AddPatientToDiagnosis adds the "PatientToDiagnosis" edges to the Diagnosis entity.
func (pc *PatientCreate) AddPatientToDiagnosis(d ...*Diagnosis) *PatientCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddPatientToDiagnosiIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PatientCreate) check() error {
	if _, ok := pc.mutation.PersonalID(); !ok {
		return &ValidationError{Name: "personalID", err: errors.New("ent: missing required field \"personalID\"")}
	}
	if v, ok := pc.mutation.PersonalID(); ok {
		if err := patient.PersonalIDValidator(v); err != nil {
			return &ValidationError{Name: "personalID", err: fmt.Errorf("ent: validator failed for field \"personalID\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientName(); !ok {
		return &ValidationError{Name: "patientName", err: errors.New("ent: missing required field \"patientName\"")}
	}
	if v, ok := pc.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return &ValidationError{Name: "patientName", err: fmt.Errorf("ent: validator failed for field \"patientName\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if v, ok := pc.mutation.Age(); ok {
		if err := patient.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if _, ok := pc.mutation.HospitalNumber(); !ok {
		return &ValidationError{Name: "hospitalNumber", err: errors.New("ent: missing required field \"hospitalNumber\"")}
	}
	if v, ok := pc.mutation.HospitalNumber(); ok {
		if err := patient.HospitalNumberValidator(v); err != nil {
			return &ValidationError{Name: "hospitalNumber", err: fmt.Errorf("ent: validator failed for field \"hospitalNumber\": %w", err)}
		}
	}
	if _, ok := pc.mutation.DrugAllergy(); !ok {
		return &ValidationError{Name: "drugAllergy", err: errors.New("ent: missing required field \"drugAllergy\"")}
	}
	if v, ok := pc.mutation.DrugAllergy(); ok {
		if err := patient.DrugAllergyValidator(v); err != nil {
			return &ValidationError{Name: "drugAllergy", err: fmt.Errorf("ent: validator failed for field \"drugAllergy\": %w", err)}
		}
	}
	if _, ok := pc.mutation.AddedTime(); !ok {
		return &ValidationError{Name: "added_time", err: errors.New("ent: missing required field \"added_time\"")}
	}
	return nil
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		_node = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PersonalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPersonalID,
		})
		_node.PersonalID = value
	}
	if value, ok := pc.mutation.PatientName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
		_node.PatientName = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := pc.mutation.HospitalNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldHospitalNumber,
		})
		_node.HospitalNumber = value
	}
	if value, ok := pc.mutation.DrugAllergy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldDrugAllergy,
		})
		_node.DrugAllergy = value
	}
	if value, ok := pc.mutation.AddedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patient.FieldAddedTime,
		})
		_node.AddedTime = value
	}
	if nodes := pc.mutation.PrefixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.PrefixTable,
			Columns: []string{patient.PrefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientToTriageResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientToTriageResultTable,
			Columns: []string{patient.PatientToTriageResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: triageresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientToAppointmentResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientToAppointmentResultsTable,
			Columns: []string{patient.PatientToAppointmentResultsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientToMedicalProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientToMedicalProcedureTable,
			Columns: []string{patient.PatientToMedicalProcedureColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientToRightToTreatmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientToRightToTreatmentTable,
			Columns: []string{patient.PatientToRightToTreatmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientToDiagnosisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientToDiagnosisTable,
			Columns: []string{patient.PatientToDiagnosisColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PatientCreateBulk is the builder for creating many Patient entities in bulk.
type PatientCreateBulk struct {
	config
	builders []*PatientCreate
}

// Save creates the Patient entities in the database.
func (pcb *PatientCreateBulk) Save(ctx context.Context) ([]*Patient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Patient, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PatientMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PatientCreateBulk) SaveX(ctx context.Context) []*Patient {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
