package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return []ent.Field{
		field.String("HospitalName").NotEmpty().Unique(),
	}
}

// Edges of the Hospitalt.
func (Hospital) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hospital", RightToTreatment.Type),
	}
}
