package schema

import "github.com/facebookincubator/ent"

// MedicineProcuder holds the schema definition for the MedicineProcuder entity.
type MedicineProcuder struct {
	ent.Schema
}

// Fields of the MedicineProcuder.
func (MedicineProcuder) Fields() []ent.Field {
	return nil
}

// Edges of the MedicineProcuder.
func (MedicineProcuder) Edges() []ent.Edge {
	return nil
}
