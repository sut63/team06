package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// TriageResult holds the schema definition for the TriageResult entity.
type TriageResult struct {
	ent.Schema
}

// Fields of the TriageResult.
func (TriageResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("symptom").NotEmpty(),
		field.Time("triageDate").Default(time.Now),
	}
}

// Edges of the TriageResult.
func (TriageResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("triageResultToUrgencyLevel", UrgencyLevel.Type).Ref("urgencyLevelToTriageResult").Unique(),
		edge.From("triageResultToDepartment", Department.Type).Ref("departmentToTriageResult").Unique(),
		edge.From("triageResultToNurse", Nurse.Type).Ref("nurseToTriageResult").Unique(),
		edge.From("triageResultToPatient", Patient.Type).Ref("patientToTriageResult").Unique(),
	}
}
