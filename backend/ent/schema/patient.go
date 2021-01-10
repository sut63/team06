package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.Int("personalID").Positive().Unique(),
		field.String("patientName").NotEmpty(),
		field.Int("age").Positive(),
		field.String("hospitalNumber").NotEmpty().Unique(),
		field.String("drugAllergy").NotEmpty(),
		field.Time("added_time"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Prefix", Prefix.Type).Ref("PrefixToPatient").Unique(),
		edge.From("Gender", Gender.Type).Ref("GenderToPatient").Unique(),
		edge.From("Bloodtype", BloodType.Type).Ref("BloodTypeToPatient").Unique(),
		edge.To("patientToTriageResult", TriageResult.Type),
		edge.To("PatientToAppointmentResults", AppointmentResults.Type),
		edge.To("PatientToMedicalProcedure", MedicalProcedure.Type),
		edge.To("PatientToRightToTreatment", RightToTreatment.Type),
	}
}
