// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/patient"
)

// GenderCreate is the builder for creating a Gender entity.
type GenderCreate struct {
	config
	mutation *GenderMutation
	hooks    []Hook
}

// SetGenderValue sets the "genderValue" field.
func (gc *GenderCreate) SetGenderValue(s string) *GenderCreate {
	gc.mutation.SetGenderValue(s)
	return gc
}

// AddPatientIDs adds the "patients" edge to the Patient entity by IDs.
func (gc *GenderCreate) AddPatientIDs(ids ...int) *GenderCreate {
	gc.mutation.AddPatientIDs(ids...)
	return gc
}

// AddPatients adds the "patients" edges to the Patient entity.
func (gc *GenderCreate) AddPatients(p ...*Patient) *GenderCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gc.AddPatientIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gc *GenderCreate) Mutation() *GenderMutation {
	return gc.mutation
}

// Save creates the Gender in the database.
func (gc *GenderCreate) Save(ctx context.Context) (*Gender, error) {
	var (
		err  error
		node *Gender
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenderCreate) SaveX(ctx context.Context) *Gender {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenderCreate) check() error {
	if _, ok := gc.mutation.GenderValue(); !ok {
		return &ValidationError{Name: "genderValue", err: errors.New("ent: missing required field \"genderValue\"")}
	}
	if v, ok := gc.mutation.GenderValue(); ok {
		if err := gender.GenderValueValidator(v); err != nil {
			return &ValidationError{Name: "genderValue", err: fmt.Errorf("ent: validator failed for field \"genderValue\": %w", err)}
		}
	}
	return nil
}

func (gc *GenderCreate) sqlSave(ctx context.Context) (*Gender, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GenderCreate) createSpec() (*Gender, *sqlgraph.CreateSpec) {
	var (
		_node = &Gender{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gender.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.GenderValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldGenderValue,
		})
		_node.GenderValue = value
	}
	if nodes := gc.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.PatientsTable,
			Columns: []string{gender.PatientsColumn},
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

// GenderCreateBulk is the builder for creating many Gender entities in bulk.
type GenderCreateBulk struct {
	config
	builders []*GenderCreate
}

// Save creates the Gender entities in the database.
func (gcb *GenderCreateBulk) Save(ctx context.Context) ([]*Gender, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gender, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenderMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenderCreateBulk) SaveX(ctx context.Context) []*Gender {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
