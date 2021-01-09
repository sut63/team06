package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// MedicalProcedure holds the schema definition for the User entity.
type MedicalProcedure struct {
	ent.Schema
}

// Fields of the MedicalProcedure.
func (MedicalProcedure) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Addtime"),
	}
}

// Edges of the MedicalProcedure.
func (MedicalProcedure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Patient", Patient.Type).Ref("PatientToMedicalProcedure").Unique(),
		edge.From("ProcedureType", ProcedureType.Type).Ref("ProcedureToMedicalProcedure").Unique(),
		edge.From("Doctor", Doctor.Type).Ref("DoctorToMedicalProcedure").Unique(),
	}
}
