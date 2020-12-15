package schema

import (
	"github.com/facebook/ent"
)

// Register holds the schema definition for the Register entity.
type Register struct {
	ent.Schema
}

// Fields of the Register.
func (Register) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Register.
func (Register) Edges() []ent.Edge {
	return []ent.Edge{}
}

