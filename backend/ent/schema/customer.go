package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("Customername").NotEmpty(),
		field.String("Address").NotEmpty(),
		field.String("Phonenumber").MaxLen(10).MinLen(10),
		field.String("Identificationnumber").MaxLen(13).MinLen(13),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("customer").Unique(),
		edge.From("personal", Personal.Type).Ref("customer").Unique(),
		edge.From("title", Title.Type).Ref("customer").Unique(),
		edge.To("fix", Fix.Type).StorageKey(edge.Column("customer_id")),
		edge.To("receipt", Receipt.Type).StorageKey(edge.Column("customer_id")),
	}
}
