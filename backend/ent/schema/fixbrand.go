package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Fixbrand holds the schema definition for the Fixbrand entity.
type Fixbrand struct {
	ent.Schema
}

// Fields of the Fixbrand.
func (Fixbrand) Fields() []ent.Field {
	return []ent.Field{
		field.String("fixbrandname").Unique(),
	}
}

// Edges of the Fixbrand.
func (Fixbrand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fix", Fix.Type).StorageKey(edge.Column("fixbrand_id")),
	}
}