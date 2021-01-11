package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// BloodType holds the schema definition for the BloodType entity.
type BloodType struct {
	ent.Schema
}

// Fields of the BloodType.
func (BloodType) Fields() []ent.Field {
	return []ent.Field{
		field.String("blood").NotEmpty(),
	}
}

// Edges of the BloodType.
func (BloodType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("BloodTypeToPatient", Patient.Type),
	}
}
