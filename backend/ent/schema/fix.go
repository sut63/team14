package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Fix holds the schema definition for the Fix entity.
type Fix struct {
	ent.Schema
}
// Fields of the Fix.
func (Fix) Fields() []ent.Field {
	return []ent.Field{
		field.String("productnumber").
		NotEmpty(),
		field.String("problemtype").
		NotEmpty(),
		field.String("queue").Unique().
		NotEmpty(),
		field.Time("date"),
	}
}

// Edges of the Fix.
func (Fix) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fix", Adminrepair.Type).StorageKey(edge.Column("fix_id")).Unique(),
		edge.From("fixbrand", Fixbrand.Type).Ref("fix").Unique(),
		edge.From("personal", Personal.Type).Ref("fix").Unique(),
		edge.From("customer", Customer.Type).Ref("fix").Unique(),
		edge.From("fixcomtype", Fixcomtype.Type).Ref("fix").Unique(),
	}
}