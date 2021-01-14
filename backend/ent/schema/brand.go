package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Brand holds the schema definition for the Brand entity.
type Brand struct {
	ent.Schema
}

// Fields of the Brand.
func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("Brandname").Unique(),
	}
}

// Edges of the Brand.
func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).StorageKey(edge.Column("Brand")),
	}
}