package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Typeproduct holds the schema definition for the Typeproduct entity.
type Typeproduct struct {
	ent.Schema
}

// Fields of the Typeproduct.
func (Typeproduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("Typeproductname").Unique(),
	}
}

// Edges of the Typeproduct.
func (Typeproduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).StorageKey(edge.Column("Typeproduct")),
	}
}