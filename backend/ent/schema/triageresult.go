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
		field.Float("height").Positive(),
		field.Float("weight").Positive(),
		field.Float("pressure").Positive(),
		field.String("symptom").NotEmpty(),
		field.Time("triageDate").Default(time.Now),
	}
}

// Edges of the TriageResult.
func (TriageResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("urgencyLevel", UrgencyLevel.Type).Ref("triageResult").Unique(),
		edge.From("department", Department.Type).Ref("triageResult").Unique(),
		edge.From("nurse", Nurse.Type).Ref("triageResult").Unique(),
		edge.From("patient", Patient.Type).Ref("triageResult").Unique(),
	}
}
