package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Diagnosis holds the schema definition for the Diagnosis entity.
type Diagnosis struct {
	ent.Schema
}

// Fields of the section.
func (Diagnosis) Fields() []ent.Field {
	return []ent.Field{

		field.String("symptom").NotEmpty(),
		field.String("Opinionresult").NotEmpty(),
		field.Time("diagnosisDate").Default(time.Now),
	}
}

// Edges of the section.
func (Diagnosis) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("Doctor_name", Doctor.Type).
			Ref("DoctorToDiagnosis").
			Unique(),
		edge.From("Patient", Patient.Type).
			Ref("PatientToDiagnosis").
			Unique(),
		edge.From("type", TreatmentType.Type).
			Ref("TreatmentTypeToDiagnosis").
			Unique(),
	}
}
