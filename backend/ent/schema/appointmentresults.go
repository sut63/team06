package schema

import "github.com/facebookincubator/ent"

// AppointmentResults holds the schema definition for the AppointmentResults entity.
type AppointmentResults struct {
	ent.Schema
}

// Fields of the AppointmentResults.
func (AppointmentResults) Fields() []ent.Field {
	return nil
}

// Edges of the AppointmentResults.
func (AppointmentResults) Edges() []ent.Edge {
	return nil
}
