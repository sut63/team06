package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"errors"
	"regexp"
)

// Diagnosis holds the schema definition for the Diagnosis entity.
type Diagnosis struct {
	ent.Schema
}

// Fields of the Diagnosis.
func (Diagnosis) Fields() []ent.Field {
	return []ent.Field{
		field.String("symptom").NotEmpty(),
		field.String("Opinionresult").
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[ก-๘]", s)
				if !match {
					return errors.New("กรุณากรอกภาษาไทย")
				}
				return nil
			}),
		field.String("note").MaxLen(25),

		//field.DATE

		field.Time("diagnosisDate"),
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
