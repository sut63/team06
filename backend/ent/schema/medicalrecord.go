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
		field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the MedicalRecord.
func (MedicalRecord) Edges() []ent.Edge {
	return nil
}
