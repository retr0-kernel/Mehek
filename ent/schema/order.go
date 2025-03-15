package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("required_by"),
		field.String("status").Default("pending"),
		field.Float("total_price").Default(0),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("brand", Brand.Type).
			Ref("orders").
			Unique().
			Required(),
		edge.To("items", OrderItem.Type),
		edge.From("resource_allocations", ResourceAllocation.Type).
			Ref("order"),
	}
}