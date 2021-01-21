// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/patient"
)

// BloodTypeCreate is the builder for creating a BloodType entity.
type BloodTypeCreate struct {
	config
	mutation *BloodTypeMutation
	hooks    []Hook
}

// SetBloodValue sets the "bloodValue" field.
func (btc *BloodTypeCreate) SetBloodValue(s string) *BloodTypeCreate {
	btc.mutation.SetBloodValue(s)
	return btc
}

// AddPatientIDs adds the "patients" edge to the Patient entity by IDs.
func (btc *BloodTypeCreate) AddPatientIDs(ids ...int) *BloodTypeCreate {
	btc.mutation.AddPatientIDs(ids...)
	return btc
}

// AddPatients adds the "patients" edges to the Patient entity.
func (btc *BloodTypeCreate) AddPatients(p ...*Patient) *BloodTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return btc.AddPatientIDs(ids...)
}

// Mutation returns the BloodTypeMutation object of the builder.
func (btc *BloodTypeCreate) Mutation() *BloodTypeMutation {
	return btc.mutation
}

// Save creates the BloodType in the database.
func (btc *BloodTypeCreate) Save(ctx context.Context) (*BloodType, error) {
	var (
		err  error
		node *BloodType
	)
	if len(btc.hooks) == 0 {
		if err = btc.check(); err != nil {
			return nil, err
		}
		node, err = btc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BloodTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = btc.check(); err != nil {
				return nil, err
			}
			btc.mutation = mutation
			node, err = btc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(btc.hooks) - 1; i >= 0; i-- {
			mut = btc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, btc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (btc *BloodTypeCreate) SaveX(ctx context.Context) *BloodType {
	v, err := btc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (btc *BloodTypeCreate) check() error {
	if _, ok := btc.mutation.BloodValue(); !ok {
		return &ValidationError{Name: "bloodValue", err: errors.New("ent: missing required field \"bloodValue\"")}
	}
	if v, ok := btc.mutation.BloodValue(); ok {
		if err := bloodtype.BloodValueValidator(v); err != nil {
			return &ValidationError{Name: "bloodValue", err: fmt.Errorf("ent: validator failed for field \"bloodValue\": %w", err)}
		}
	}
	return nil
}

func (btc *BloodTypeCreate) sqlSave(ctx context.Context) (*BloodType, error) {
	_node, _spec := btc.createSpec()
	if err := sqlgraph.CreateNode(ctx, btc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (btc *BloodTypeCreate) createSpec() (*BloodType, *sqlgraph.CreateSpec) {
	var (
		_node = &BloodType{config: btc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bloodtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bloodtype.FieldID,
			},
		}
	)
	if value, ok := btc.mutation.BloodValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bloodtype.FieldBloodValue,
		})
		_node.BloodValue = value
	}
	if nodes := btc.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bloodtype.PatientsTable,
			Columns: []string{bloodtype.PatientsColumn},
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

// BloodTypeCreateBulk is the builder for creating many BloodType entities in bulk.
type BloodTypeCreateBulk struct {
	config
	builders []*BloodTypeCreate
}

// Save creates the BloodType entities in the database.
func (btcb *BloodTypeCreateBulk) Save(ctx context.Context) ([]*BloodType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(btcb.builders))
	nodes := make([]*BloodType, len(btcb.builders))
	mutators := make([]Mutator, len(btcb.builders))
	for i := range btcb.builders {
		func(i int, root context.Context) {
			builder := btcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BloodTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, btcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, btcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, btcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (btcb *BloodTypeCreateBulk) SaveX(ctx context.Context) []*BloodType {
	v, err := btcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
