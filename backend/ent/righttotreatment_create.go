// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/hospital"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/righttotreatment"
	"github.com/team06/app/ent/righttotreatmenttype"
)

// RightToTreatmentCreate is the builder for creating a RightToTreatment entity.
type RightToTreatmentCreate struct {
	config
	mutation *RightToTreatmentMutation
	hooks    []Hook
}

// SetStartTime sets the "StartTime" field.
func (rttc *RightToTreatmentCreate) SetStartTime(t time.Time) *RightToTreatmentCreate {
	rttc.mutation.SetStartTime(t)
	return rttc
}

// SetEndTime sets the "EndTime" field.
func (rttc *RightToTreatmentCreate) SetEndTime(t time.Time) *RightToTreatmentCreate {
	rttc.mutation.SetEndTime(t)
	return rttc
}

// SetHospitalID sets the "Hospital" edge to the Hospital entity by ID.
func (rttc *RightToTreatmentCreate) SetHospitalID(id int) *RightToTreatmentCreate {
	rttc.mutation.SetHospitalID(id)
	return rttc
}

// SetNillableHospitalID sets the "Hospital" edge to the Hospital entity by ID if the given value is not nil.
func (rttc *RightToTreatmentCreate) SetNillableHospitalID(id *int) *RightToTreatmentCreate {
	if id != nil {
		rttc = rttc.SetHospitalID(*id)
	}
	return rttc
}

// SetHospital sets the "Hospital" edge to the Hospital entity.
func (rttc *RightToTreatmentCreate) SetHospital(h *Hospital) *RightToTreatmentCreate {
	return rttc.SetHospitalID(h.ID)
}

// SetRightToTreatmentTypeID sets the "RightToTreatmentType" edge to the RightToTreatmentType entity by ID.
func (rttc *RightToTreatmentCreate) SetRightToTreatmentTypeID(id int) *RightToTreatmentCreate {
	rttc.mutation.SetRightToTreatmentTypeID(id)
	return rttc
}

// SetNillableRightToTreatmentTypeID sets the "RightToTreatmentType" edge to the RightToTreatmentType entity by ID if the given value is not nil.
func (rttc *RightToTreatmentCreate) SetNillableRightToTreatmentTypeID(id *int) *RightToTreatmentCreate {
	if id != nil {
		rttc = rttc.SetRightToTreatmentTypeID(*id)
	}
	return rttc
}

// SetRightToTreatmentType sets the "RightToTreatmentType" edge to the RightToTreatmentType entity.
func (rttc *RightToTreatmentCreate) SetRightToTreatmentType(r *RightToTreatmentType) *RightToTreatmentCreate {
	return rttc.SetRightToTreatmentTypeID(r.ID)
}

// SetPatientID sets the "Patient" edge to the Patient entity by ID.
func (rttc *RightToTreatmentCreate) SetPatientID(id int) *RightToTreatmentCreate {
	rttc.mutation.SetPatientID(id)
	return rttc
}

// SetNillablePatientID sets the "Patient" edge to the Patient entity by ID if the given value is not nil.
func (rttc *RightToTreatmentCreate) SetNillablePatientID(id *int) *RightToTreatmentCreate {
	if id != nil {
		rttc = rttc.SetPatientID(*id)
	}
	return rttc
}

// SetPatient sets the "Patient" edge to the Patient entity.
func (rttc *RightToTreatmentCreate) SetPatient(p *Patient) *RightToTreatmentCreate {
	return rttc.SetPatientID(p.ID)
}

// Mutation returns the RightToTreatmentMutation object of the builder.
func (rttc *RightToTreatmentCreate) Mutation() *RightToTreatmentMutation {
	return rttc.mutation
}

// Save creates the RightToTreatment in the database.
func (rttc *RightToTreatmentCreate) Save(ctx context.Context) (*RightToTreatment, error) {
	var (
		err  error
		node *RightToTreatment
	)
	if len(rttc.hooks) == 0 {
		if err = rttc.check(); err != nil {
			return nil, err
		}
		node, err = rttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RightToTreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rttc.check(); err != nil {
				return nil, err
			}
			rttc.mutation = mutation
			node, err = rttc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rttc.hooks) - 1; i >= 0; i-- {
			mut = rttc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rttc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rttc *RightToTreatmentCreate) SaveX(ctx context.Context) *RightToTreatment {
	v, err := rttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (rttc *RightToTreatmentCreate) check() error {
	if _, ok := rttc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "StartTime", err: errors.New("ent: missing required field \"StartTime\"")}
	}
	if _, ok := rttc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "EndTime", err: errors.New("ent: missing required field \"EndTime\"")}
	}
	return nil
}

func (rttc *RightToTreatmentCreate) sqlSave(ctx context.Context) (*RightToTreatment, error) {
	_node, _spec := rttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rttc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rttc *RightToTreatmentCreate) createSpec() (*RightToTreatment, *sqlgraph.CreateSpec) {
	var (
		_node = &RightToTreatment{config: rttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: righttotreatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: righttotreatment.FieldID,
			},
		}
	)
	if value, ok := rttc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: righttotreatment.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := rttc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: righttotreatment.FieldEndTime,
		})
		_node.EndTime = value
	}
	if nodes := rttc.mutation.HospitalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   righttotreatment.HospitalTable,
			Columns: []string{righttotreatment.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rttc.mutation.RightToTreatmentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   righttotreatment.RightToTreatmentTypeTable,
			Columns: []string{righttotreatment.RightToTreatmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatmenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rttc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   righttotreatment.PatientTable,
			Columns: []string{righttotreatment.PatientColumn},
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
	return _node, _spec
}

// RightToTreatmentCreateBulk is the builder for creating many RightToTreatment entities in bulk.
type RightToTreatmentCreateBulk struct {
	config
	builders []*RightToTreatmentCreate
}

// Save creates the RightToTreatment entities in the database.
func (rttcb *RightToTreatmentCreateBulk) Save(ctx context.Context) ([]*RightToTreatment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rttcb.builders))
	nodes := make([]*RightToTreatment, len(rttcb.builders))
	mutators := make([]Mutator, len(rttcb.builders))
	for i := range rttcb.builders {
		func(i int, root context.Context) {
			builder := rttcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RightToTreatmentMutation)
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
					_, err = mutators[i+1].Mutate(root, rttcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rttcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rttcb *RightToTreatmentCreateBulk) SaveX(ctx context.Context) []*RightToTreatment {
	v, err := rttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}