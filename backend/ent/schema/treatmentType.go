package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
