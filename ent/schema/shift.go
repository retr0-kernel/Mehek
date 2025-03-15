package schema

import (
	
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Shift holds the schema definition for the Shift entity.
type Shift struct {
	ent.Schema
}

// Fields of the Shift.
func (Shift) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
	}
}

// Edges of the Shift.
func (Shift) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("staff", Staff.Type).
			Ref("shifts").
			Unique().
			Required(),
		edge.To("allocations", ResourceAllocation.Type),
	}
}