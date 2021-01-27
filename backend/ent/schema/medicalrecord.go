package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// MedicalRecord holds the schema definition for the MedicalRecord entity.
type MedicalRecord struct {
	ent.Schema
}

// Fields of the MedicalRecord.
func (MedicalRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty().Unique(),
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the MedicalRecord.
func (MedicalRecord) Edges() []ent.Edge {
	return nil
}
