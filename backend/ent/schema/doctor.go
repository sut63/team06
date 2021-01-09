package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Doctor holds the schema definition for the Patient entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("doctorName").NotEmpty(),
		field.String("doctorUsername").NotEmpty().Unique(),
		field.String("doctorPassword").NotEmpty(),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("DoctorToMedicalProcedure", MedicalProcedure.Type),
		edge.To("DoctorToAppointmentResults", AppointmentResults.Type),
		edge.To("DoctorToDiagnosis", Diagnosis.Type),
	}
}
