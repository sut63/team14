package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Receipt holds the schema definition for the Receipt entity.
type Receipt struct {
	ent.Schema
}

// Fields of the Receipt.
func (Receipt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time").
			Default(time.Now),
	}
}

// Edges of the Receipt.
func (Receipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("paymenttype", PaymentType.Type).Ref("receipt").Unique(),
		edge.From("adminrepair", Adminrepair.Type).Ref("receipt").Unique(),
		edge.From("personal", Personal.Type).Ref("receipt").Unique(),
		edge.From("customer", Customer.Type).Ref("receipt").Unique(),
		edge.From("product", Product.Type).Ref("receipt").Unique(),
	}
}
