package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ProcedureType holds the schema definition for the User entity.
type ProcedureType struct {
	ent.Schema
}

// Fields of the ProcedureType.
func (ProcedureType) Fields() []ent.Field {
	return []ent.Field{
		field.String("ProcedureType"),
	}
}

// Edges of the ProcedureType.
func (ProcedureType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ProcedureToMedicalProcedure", MedicalProcedure.Type),
	}
}
