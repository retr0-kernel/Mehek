package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("role").NotEmpty(),
		field.String("skills"),
		field.JSON("availability", map[string][]string{}),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("kitchen", Kitchen.Type).
			Ref("staff").
			Unique().
			Required(),
		edge.To("shifts", Shift.Type),
	}
}