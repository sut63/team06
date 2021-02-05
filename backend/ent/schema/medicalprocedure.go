package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"errors"
	"regexp"
)

// MedicalProcedure holds the schema definition for the User entity.
type MedicalProcedure struct {
	ent.Schema
}

// Fields of the MedicalProcedure.
func (MedicalProcedure) Fields() []ent.Field {
	return []ent.Field{
		field.String("procedureOrder").Validate(func(s string) error {
            match, _ := regexp.MatchString("[U]+[N]+[S]\\d{6}" ,s)
                if !match {
                    return errors.New("รูปแบบรหัสไม่ถูกต้อง")
                }
                return nil
            }),
		field.String("procedureRoom").MaxLen(4).MinLen(4),
		field.Time("Addtime"),
		field.String("procedureDescripe").NotEmpty(),
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
