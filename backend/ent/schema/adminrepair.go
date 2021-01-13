package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Adminrepair holds the schema definition for the Adminrepair entity.
type Adminrepair struct {
	ent.Schema
}

// Fields of the Adminrepair.
func (Adminrepair) Fields() []ent.Field {
	return []ent.Field{
		field.String("equipmentdamate").NotEmpty(),
	}
}

// Edges of the Adminrepair.
func (Adminrepair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("AdminrepairPersonal", Personal.Type).Ref("personal").Unique(),
		edge.From("AdminrepairFix", Fix.Type).Ref("fix").Unique(),
		edge.From("AdminrepairProduct", Product.Type).Ref("product").Unique(),
	}
}
