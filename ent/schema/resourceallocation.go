package schema

import (

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ResourceAllocation holds the schema definition for the ResourceAllocation entity.
type ResourceAllocation struct {
	ent.Schema
}

// Fields of the ResourceAllocation.
func (ResourceAllocation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("status").Default("scheduled"),
	}
}

// Edges of the ResourceAllocation.
func (ResourceAllocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resource", KitchenResource.Type).
			Ref("allocations").
			Unique().
			Required(),
		edge.To("order", Order.Type).
			Unique().
			Required(),
		edge.From("shift", Shift.Type).
			Ref("allocations").
			Unique().
			Required(),
	}
}