package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Badge holds the schema definition for the Badge entity.
type Badge struct {
	ent.Schema
}

// Fields of the Badge.
func (Badge) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("aws"),
	}
}

// Mixin of the Badge.
func (Badge) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Edges of the Badge.
func (Badge) Edges() []ent.Edge {
	return nil
}
