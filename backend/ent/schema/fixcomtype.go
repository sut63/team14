package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Fixcomtype holds the schema definition for the Fixcomtype entity.
type Fixcomtype struct {
	ent.Schema
}

// Fields of the Fixcomtype.
func (Fixcomtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("fixcomtypename").Unique(),
	}
}

// Edges of the Fixcomtype.
func (Fixcomtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fix", Fix.Type).StorageKey(edge.Column("fixcomtype_id")),
	}
}