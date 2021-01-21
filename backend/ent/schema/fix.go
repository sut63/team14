package schema

import (
	"regexp"

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
		field.String("Productnumber").Match(regexp.MustCompile("[a-zA-Z0-9_]+$")).
			NotEmpty(),
		field.String("Problemtype").MaxLen(100).MinLen(5),
		field.String("Queue").Unique().MinLen(5).Match(regexp.MustCompile("[A-Z0-9_]+$")),
		field.Time("Date"),
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
