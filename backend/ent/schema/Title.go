package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Title holds the schema definition for the Title entity.
type Title struct {
	ent.Schema
}

// Fields of the Title.
func (Title) Fields() []ent.Field {
	return []ent.Field{
		field.String("titlename").Unique(),
	}
}

// Edges of the Title.
func (Title) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personal", Personal.Type).StorageKey(edge.Column("title_id")),
		edge.To("customer", Customer.Type).StorageKey(edge.Column("title_id")),
	}
}
