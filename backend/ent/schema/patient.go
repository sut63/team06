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
		field.String("hospitalNumber").NotEmpty(),
		field.String("patientName").NotEmpty(),
		field.String("drugAllergy").NotEmpty(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patient_details", PatientDetail.Type),
		edge.To("triageResult", TriageResult.Type),
		edge.To("PatientToAppointmentResults", AppointmentResults.Type),
		edge.To("PatientToMedicalProcedure", MedicalProcedure.Type),
		edge.To("PatientToRightToTreatment", RightToTreatment.Type),
		edge.To("PatientToDiagnosis", Diagnosis.Type),
	}
}
