package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// TreatmentType holds the schema definition for the TreatmentType entity.
type TreatmentType struct {
	ent.Schema
}

// Fields of the TreatmentType.
func (TreatmentType) Fields() []ent.Field {
	return []ent.Field{

		field.String("Type").NotEmpty(),
	}
}

// Edges of the TreatmentType.
func (TreatmentType) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("TreatmentTypeToDiagnosis", Diagnosis.Type),
	}
}
