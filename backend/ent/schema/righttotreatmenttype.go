package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// RightToTreatmentType holds the schema definition for the RightToTreatmentType entity.
type RightToTreatmentType struct {
	ent.Schema
}

// Fields of the RightToTreatmentType.
func (RightToTreatmentType) Fields() []ent.Field {
	return []ent.Field{
		field.String("TypeName").NotEmpty().Unique(),
	}
}

// Edges of the RightToTreatmentType.
func (RightToTreatmentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("type", RightToTreatment.Type),
	}
}
