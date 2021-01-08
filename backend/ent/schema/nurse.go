package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("NurseToTriageResult", TriageResult.Type),
		edge.To("NurseToAppointmentResults", AppointmentResults.Type),
	}
}
