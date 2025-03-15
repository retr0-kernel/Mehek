package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Kitchen holds the schema definition for the Kitchen entity.
type Kitchen struct {
	ent.Schema
}

// Fields of the Kitchen.
func (Kitchen) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("location").NotEmpty(),
		field.Int("capacity").Positive(),
		field.JSON("operating_hours", map[string][]string{}),
	}
}

// Edges of the Kitchen.
func (Kitchen) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("resources", KitchenResource.Type),
		edge.To("staff", Staff.Type),
		edge.To("brands", Brand.Type),
		edge.To("inventory", InventoryItem.Type),
	}
}
