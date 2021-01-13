package schema

import (
	"time"

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
		field.Int("personalID").Positive(),
		field.String("patientName").NotEmpty(),
		field.Int("age").Positive(),
		field.String("hospitalNumber").NotEmpty(),
		field.String("drugAllergy").NotEmpty(),
		field.Time("addedDate").Default(time.Now),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prefix", Prefix.Type).Ref("patient").Unique(),
		edge.From("gender", Gender.Type).Ref("patient").Unique(),
		edge.From("bloodtype", BloodType.Type).Ref("patient").Unique(),
		edge.To("triageResult", TriageResult.Type),
		edge.To("PatientToAppointmentResults", AppointmentResults.Type),
		edge.To("PatientToMedicalProcedure", MedicalProcedure.Type),
		edge.To("PatientToRightToTreatment", RightToTreatment.Type),
		edge.To("PatientToDiagnosis", Diagnosis.Type),
	}
}
