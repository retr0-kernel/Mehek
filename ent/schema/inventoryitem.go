package schema

import (
	
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InventoryItem holds the schema definition for the InventoryItem entity.
type InventoryItem struct {
	ent.Schema
}

// Fields of the InventoryItem.
func (InventoryItem) Fields() []ent.Field {
	return []ent.Field{
		field.Float("quantity").Positive(),
		field.Time("expiration_date").Optional(),
	}
}

// Edges of the InventoryItem.
func (InventoryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("kitchen", Kitchen.Type).
			Ref("inventory").
			Unique().
			Required(),
		edge.To("ingredient", Ingredient.Type).
			Unique().
			Required(),
	}
}