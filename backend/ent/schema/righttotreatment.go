package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// RightToTreatment holds the schema definition for the RightToTreatment entity.
type RightToTreatment struct {
	ent.Schema
}

// Fields of the RightToTreatment.
func (RightToTreatment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Addedtime"),
	}
}

// Edges of the RightToTreatment.
func (RightToTreatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Hospital", Hospital.Type).
			Ref("hospital").
			Unique(),
		edge.From("RightToTreatmentType", RightToTreatmentType.Type).
			Ref("type").
			Unique(),
		edge.From("Patient", Patient.Type).
			Ref("PatientToRightToTreatment").
			Unique(),
	}
}