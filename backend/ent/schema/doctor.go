package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Doctor holds the schema definition for the Doctor entity.
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

		edge.To("DoctorToDiagnosis", Diagnosis.Type),
		edge.To("DoctorToAppointmentResults", AppointmentResults.Type),
		edge.To("DoctorToMedicalProcedure", MedicalProcedure.Type),
	}
}
