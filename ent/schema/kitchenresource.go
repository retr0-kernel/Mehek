package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// KitchenResource holds the schema definition for the KitchenResource entity.
type KitchenResource struct {
	ent.Schema
}

// Fields of the KitchenResource.
func (KitchenResource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("type").NotEmpty(),
		field.Int("capacity").Positive(),
		field.Bool("available").Default(true),
	}
}

// Edges of the KitchenResource.
func (KitchenResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("kitchen", Kitchen.Type).
			Ref("resources").
			Unique().
			Required(),
		edge.To("allocations", ResourceAllocation.Type),
	}
}