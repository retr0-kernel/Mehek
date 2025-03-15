package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Brand holds the schema definition for the Brand entity.
type Brand struct {
	ent.Schema
}

// Fields of the Brand.
func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("cuisine_type"),
		field.String("logo_url"),
	}
}

// Edges of the Brand.
func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("kitchen", Kitchen.Type).
			Ref("brands").
			Unique().
			Required(),
		edge.To("menus", Menu.Type),
		edge.To("orders", Order.Type),
	}
}