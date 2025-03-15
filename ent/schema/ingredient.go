package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("unit").NotEmpty(),
		field.Float("cost_per_unit").Positive(),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu_items", MenuItem.Type).
			Ref("ingredients"),
		edge.From("inventory_items", InventoryItem.Type).
			Ref("ingredient"),
	}
}