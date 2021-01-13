package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PaymentType holds the schema definition for the PaymentType entity.
type PaymentType struct {
	ent.Schema
}

// Fields of the PaymentType.
func (PaymentType) Fields() []ent.Field {
	return []ent.Field{
		field.String("Typename").NotEmpty(),
	}
}

// Edges of the PaymentType.
func (PaymentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receipt", Receipt.Type).StorageKey(edge.Column("paymenttype_id")),
	}
}
