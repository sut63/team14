package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("Productname").
			NotEmpty(),
		field.String("Numberofproduct").
			NotEmpty(),
		field.String("Price").MinLen(3),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Adminrepair.Type).StorageKey(edge.Column("product_id")),
		edge.From("brand", Brand.Type).Ref("product").Unique(),
		edge.From("typeproduct", Typeproduct.Type).Ref("product").Unique(),
		edge.From("personal", Personal.Type).Ref("product").Unique(),
		edge.To("receipt", Receipt.Type).StorageKey(edge.Column("product_id")),
	}
}
