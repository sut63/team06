package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// PatientDetail holds the schema definition for the PatientDetail entity.
type PatientDetail struct {
	ent.Schema
}

// Fields of the PatientDetail.
func (PatientDetail) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the PatientDetail.
func (PatientDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prefix", Prefix.Type).Ref("patient_details").Unique(),
		edge.From("gender", Gender.Type).Ref("patient_details").Unique(),
		edge.From("bloodtype", BloodType.Type).Ref("patient_details").Unique(),
		edge.From("patient", Patient.Type).Ref("patient_details").Unique(),
	}
}
