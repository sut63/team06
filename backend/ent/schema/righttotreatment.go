package schema

import (
    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/edge"
    "github.com/facebook/ent/schema/field"
    "regexp"

)

// RightToTreatment holds the schema definition for the RightToTreatment entity.
type RightToTreatment struct {
    ent.Schema
}

// Fields of the RightToTreatment.
func (RightToTreatment) Fields() []ent.Field {
    return []ent.Field{
        field.Time("StartTime"),
        field.Time("EndTime"),
        field.String("tel").Match(regexp.MustCompile("[0]\\d{9}")),
        field.String("idennum").MaxLen(13).MinLen(13),
        field.Int("age").Min(0),
    }
}

// Edges of the RightToTreatment.
func (RightToTreatment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("Hospital", Hospital.Type).
            Ref("hospital").
            Unique(),
        edge.From("RightToTreatmentType", RightToTreatmentType.Type).
            Ref("type").
            Unique(),
        edge.From("Patient", Patient.Type).
            Ref("PatientToRightToTreatment").
            Unique(),
    }
}