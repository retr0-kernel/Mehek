package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MenuItem holds the schema definition for the MenuItem entity.
type MenuItem struct {
	ent.Schema
}

// Fields of the MenuItem.
func (MenuItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Float("price").Positive(),
		field.Int("prep_time").Positive(),
		field.String("equipment_needed"),
	}
}

// Edges of the MenuItem.
func (MenuItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu", Menu.Type).
			Ref("items").
			Unique().
			Required(),
		edge.To("ingredients", Ingredient.Type),
		edge.From("order_items", OrderItem.Type).
			Ref("menu_item"),
	}
}