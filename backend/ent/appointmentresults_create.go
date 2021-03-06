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
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/nurse"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/room"
)

// AppointmentResultsCreate is the builder for creating a AppointmentResults entity.
type AppointmentResultsCreate struct {
	config
	mutation *AppointmentResultsMutation
	hooks    []Hook
}

// SetCauseAppoint sets the "causeAppoint" field.
func (arc *AppointmentResultsCreate) SetCauseAppoint(s string) *AppointmentResultsCreate {
	arc.mutation.SetCauseAppoint(s)
	return arc
}

// SetAdvice sets the "advice" field.
func (arc *AppointmentResultsCreate) SetAdvice(s string) *AppointmentResultsCreate {
	arc.mutation.SetAdvice(s)
	return arc
}

// SetDateAppoint sets the "dateAppoint" field.
func (arc *AppointmentResultsCreate) SetDateAppoint(t time.Time) *AppointmentResultsCreate {
	arc.mutation.SetDateAppoint(t)
	return arc
}

// SetTimeAppoint sets the "timeAppoint" field.
func (arc *AppointmentResultsCreate) SetTimeAppoint(t time.Time) *AppointmentResultsCreate {
	arc.mutation.SetTimeAppoint(t)
	return arc
}

// SetAddtimeSave sets the "addtimeSave" field.
func (arc *AppointmentResultsCreate) SetAddtimeSave(t time.Time) *AppointmentResultsCreate {
	arc.mutation.SetAddtimeSave(t)
	return arc
}

// SetNillableAddtimeSave sets the "addtimeSave" field if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableAddtimeSave(t *time.Time) *AppointmentResultsCreate {
	if t != nil {
		arc.SetAddtimeSave(*t)
	}
	return arc
}

// SetHourBeforeAppoint sets the "hourBeforeAppoint" field.
func (arc *AppointmentResultsCreate) SetHourBeforeAppoint(i int) *AppointmentResultsCreate {
	arc.mutation.SetHourBeforeAppoint(i)
	return arc
}

// SetNillableHourBeforeAppoint sets the "hourBeforeAppoint" field if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableHourBeforeAppoint(i *int) *AppointmentResultsCreate {
	if i != nil {
		arc.SetHourBeforeAppoint(*i)
	}
	return arc
}

// SetMinuteBeforeAppoint sets the "minuteBeforeAppoint" field.
func (arc *AppointmentResultsCreate) SetMinuteBeforeAppoint(i int) *AppointmentResultsCreate {
	arc.mutation.SetMinuteBeforeAppoint(i)
	return arc
}

// SetNillableMinuteBeforeAppoint sets the "minuteBeforeAppoint" field if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableMinuteBeforeAppoint(i *int) *AppointmentResultsCreate {
	if i != nil {
		arc.SetMinuteBeforeAppoint(*i)
	}
	return arc
}

// SetAppointmentResultsToPatientID sets the "appointmentResultsToPatient" edge to the Patient entity by ID.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToPatientID(id int) *AppointmentResultsCreate {
	arc.mutation.SetAppointmentResultsToPatientID(id)
	return arc
}

// SetNillableAppointmentResultsToPatientID sets the "appointmentResultsToPatient" edge to the Patient entity by ID if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableAppointmentResultsToPatientID(id *int) *AppointmentResultsCreate {
	if id != nil {
		arc = arc.SetAppointmentResultsToPatientID(*id)
	}
	return arc
}

// SetAppointmentResultsToPatient sets the "appointmentResultsToPatient" edge to the Patient entity.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToPatient(p *Patient) *AppointmentResultsCreate {
	return arc.SetAppointmentResultsToPatientID(p.ID)
}

// SetAppointmentResultsToNurseID sets the "appointmentResultsToNurse" edge to the Nurse entity by ID.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToNurseID(id int) *AppointmentResultsCreate {
	arc.mutation.SetAppointmentResultsToNurseID(id)
	return arc
}

// SetNillableAppointmentResultsToNurseID sets the "appointmentResultsToNurse" edge to the Nurse entity by ID if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableAppointmentResultsToNurseID(id *int) *AppointmentResultsCreate {
	if id != nil {
		arc = arc.SetAppointmentResultsToNurseID(*id)
	}
	return arc
}

// SetAppointmentResultsToNurse sets the "appointmentResultsToNurse" edge to the Nurse entity.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToNurse(n *Nurse) *AppointmentResultsCreate {
	return arc.SetAppointmentResultsToNurseID(n.ID)
}

// SetAppointmentResultsToDoctorID sets the "appointmentResultsToDoctor" edge to the Doctor entity by ID.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToDoctorID(id int) *AppointmentResultsCreate {
	arc.mutation.SetAppointmentResultsToDoctorID(id)
	return arc
}

// SetNillableAppointmentResultsToDoctorID sets the "appointmentResultsToDoctor" edge to the Doctor entity by ID if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableAppointmentResultsToDoctorID(id *int) *AppointmentResultsCreate {
	if id != nil {
		arc = arc.SetAppointmentResultsToDoctorID(*id)
	}
	return arc
}

// SetAppointmentResultsToDoctor sets the "appointmentResultsToDoctor" edge to the Doctor entity.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToDoctor(d *Doctor) *AppointmentResultsCreate {
	return arc.SetAppointmentResultsToDoctorID(d.ID)
}

// SetAppointmentResultsToRoomID sets the "appointmentResultsToRoom" edge to the Room entity by ID.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToRoomID(id int) *AppointmentResultsCreate {
	arc.mutation.SetAppointmentResultsToRoomID(id)
	return arc
}

// SetNillableAppointmentResultsToRoomID sets the "appointmentResultsToRoom" edge to the Room entity by ID if the given value is not nil.
func (arc *AppointmentResultsCreate) SetNillableAppointmentResultsToRoomID(id *int) *AppointmentResultsCreate {
	if id != nil {
		arc = arc.SetAppointmentResultsToRoomID(*id)
	}
	return arc
}

// SetAppointmentResultsToRoom sets the "appointmentResultsToRoom" edge to the Room entity.
func (arc *AppointmentResultsCreate) SetAppointmentResultsToRoom(r *Room) *AppointmentResultsCreate {
	return arc.SetAppointmentResultsToRoomID(r.ID)
}

// Mutation returns the AppointmentResultsMutation object of the builder.
func (arc *AppointmentResultsCreate) Mutation() *AppointmentResultsMutation {
	return arc.mutation
}

// Save creates the AppointmentResults in the database.
func (arc *AppointmentResultsCreate) Save(ctx context.Context) (*AppointmentResults, error) {
	var (
		err  error
		node *AppointmentResults
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppointmentResultsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			node, err = arc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			mut = arc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AppointmentResultsCreate) SaveX(ctx context.Context) *AppointmentResults {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (arc *AppointmentResultsCreate) defaults() {
	if _, ok := arc.mutation.AddtimeSave(); !ok {
		v := appointmentresults.DefaultAddtimeSave()
		arc.mutation.SetAddtimeSave(v)
	}
	if _, ok := arc.mutation.HourBeforeAppoint(); !ok {
		v := appointmentresults.DefaultHourBeforeAppoint
		arc.mutation.SetHourBeforeAppoint(v)
	}
	if _, ok := arc.mutation.MinuteBeforeAppoint(); !ok {
		v := appointmentresults.DefaultMinuteBeforeAppoint
		arc.mutation.SetMinuteBeforeAppoint(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AppointmentResultsCreate) check() error {
	if _, ok := arc.mutation.CauseAppoint(); !ok {
		return &ValidationError{Name: "causeAppoint", err: errors.New("ent: missing required field \"causeAppoint\"")}
	}
	if v, ok := arc.mutation.CauseAppoint(); ok {
		if err := appointmentresults.CauseAppointValidator(v); err != nil {
			return &ValidationError{Name: "causeAppoint", err: fmt.Errorf("ent: validator failed for field \"causeAppoint\": %w", err)}
		}
	}
	if _, ok := arc.mutation.Advice(); !ok {
		return &ValidationError{Name: "advice", err: errors.New("ent: missing required field \"advice\"")}
	}
	if v, ok := arc.mutation.Advice(); ok {
		if err := appointmentresults.AdviceValidator(v); err != nil {
			return &ValidationError{Name: "advice", err: fmt.Errorf("ent: validator failed for field \"advice\": %w", err)}
		}
	}
	if _, ok := arc.mutation.DateAppoint(); !ok {
		return &ValidationError{Name: "dateAppoint", err: errors.New("ent: missing required field \"dateAppoint\"")}
	}
	if _, ok := arc.mutation.TimeAppoint(); !ok {
		return &ValidationError{Name: "timeAppoint", err: errors.New("ent: missing required field \"timeAppoint\"")}
	}
	if _, ok := arc.mutation.AddtimeSave(); !ok {
		return &ValidationError{Name: "addtimeSave", err: errors.New("ent: missing required field \"addtimeSave\"")}
	}
	if _, ok := arc.mutation.HourBeforeAppoint(); !ok {
		return &ValidationError{Name: "hourBeforeAppoint", err: errors.New("ent: missing required field \"hourBeforeAppoint\"")}
	}
	if v, ok := arc.mutation.HourBeforeAppoint(); ok {
		if err := appointmentresults.HourBeforeAppointValidator(v); err != nil {
			return &ValidationError{Name: "hourBeforeAppoint", err: fmt.Errorf("ent: validator failed for field \"hourBeforeAppoint\": %w", err)}
		}
	}
	if _, ok := arc.mutation.MinuteBeforeAppoint(); !ok {
		return &ValidationError{Name: "minuteBeforeAppoint", err: errors.New("ent: missing required field \"minuteBeforeAppoint\"")}
	}
	if v, ok := arc.mutation.MinuteBeforeAppoint(); ok {
		if err := appointmentresults.MinuteBeforeAppointValidator(v); err != nil {
			return &ValidationError{Name: "minuteBeforeAppoint", err: fmt.Errorf("ent: validator failed for field \"minuteBeforeAppoint\": %w", err)}
		}
	}
	return nil
}

func (arc *AppointmentResultsCreate) sqlSave(ctx context.Context) (*AppointmentResults, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (arc *AppointmentResultsCreate) createSpec() (*AppointmentResults, *sqlgraph.CreateSpec) {
	var (
		_node = &AppointmentResults{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appointmentresults.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appointmentresults.FieldID,
			},
		}
	)
	if value, ok := arc.mutation.CauseAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appointmentresults.FieldCauseAppoint,
		})
		_node.CauseAppoint = value
	}
	if value, ok := arc.mutation.Advice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appointmentresults.FieldAdvice,
		})
		_node.Advice = value
	}
	if value, ok := arc.mutation.DateAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointmentresults.FieldDateAppoint,
		})
		_node.DateAppoint = value
	}
	if value, ok := arc.mutation.TimeAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointmentresults.FieldTimeAppoint,
		})
		_node.TimeAppoint = value
	}
	if value, ok := arc.mutation.AddtimeSave(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appointmentresults.FieldAddtimeSave,
		})
		_node.AddtimeSave = value
	}
	if value, ok := arc.mutation.HourBeforeAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: appointmentresults.FieldHourBeforeAppoint,
		})
		_node.HourBeforeAppoint = value
	}
	if value, ok := arc.mutation.MinuteBeforeAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: appointmentresults.FieldMinuteBeforeAppoint,
		})
		_node.MinuteBeforeAppoint = value
	}
	if nodes := arc.mutation.AppointmentResultsToPatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointmentresults.AppointmentResultsToPatientTable,
			Columns: []string{appointmentresults.AppointmentResultsToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AppointmentResultsToNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointmentresults.AppointmentResultsToNurseTable,
			Columns: []string{appointmentresults.AppointmentResultsToNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AppointmentResultsToDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointmentresults.AppointmentResultsToDoctorTable,
			Columns: []string{appointmentresults.AppointmentResultsToDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AppointmentResultsToRoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appointmentresults.AppointmentResultsToRoomTable,
			Columns: []string{appointmentresults.AppointmentResultsToRoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
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

// AppointmentResultsCreateBulk is the builder for creating many AppointmentResults entities in bulk.
type AppointmentResultsCreateBulk struct {
	config
	builders []*AppointmentResultsCreate
}

// Save creates the AppointmentResults entities in the database.
func (arcb *AppointmentResultsCreateBulk) Save(ctx context.Context) ([]*AppointmentResults, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AppointmentResults, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppointmentResultsMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AppointmentResultsCreateBulk) SaveX(ctx context.Context) []*AppointmentResults {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
