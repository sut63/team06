package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// UrgencyLevel holds the schema definition for the UrgencyLevel entity.
type UrgencyLevel struct {
	ent.Schema
}

// Fields of the UrgencyLevel.
func (UrgencyLevel) Fields() []ent.Field {
	return []ent.Field{
		field.String("urgencyName").NotEmpty().Unique(),
	}
}

// Edges of the UrgencyLevel.
func (UrgencyLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("urgencyLevelToTriageResult", TriageResult.Type),
	}
}
