package schema

import (
	"errors"
	"regexp"
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
		field.String("personalID").MaxLen(13).MinLen(13),
		field.String("hospitalNumber").Validate(func(h string) error {
			match, _ := regexp.MatchString("HN\\d", h)
			if !match {
				return errors.New("รูปแบบหมายเลขผู้ป่วยไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("patientName").NotEmpty(),
		field.String("drugAllergy").NotEmpty(),
		field.String("mobileNumber").MaxLen(10).MinLen(10),
		field.Time("added").Default(time.Now),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prefix", Prefix.Type).Ref("patients").Unique(),
		edge.From("gender", Gender.Type).Ref("patients").Unique(),
		edge.From("bloodtype", BloodType.Type).Ref("patients").Unique(),

		edge.To("triageResult", TriageResult.Type),
		edge.To("PatientToAppointmentResults", AppointmentResults.Type),
		edge.To("PatientToMedicalProcedure", MedicalProcedure.Type),
		edge.To("PatientToRightToTreatment", RightToTreatment.Type),
		edge.To("PatientToDiagnosis", Diagnosis.Type),
	}
}
