package schema

import "github.com/facebook/ent"

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return nil
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return nil
}
