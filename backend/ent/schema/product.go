package schema

import "github.com/facebookincubator/ent"

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return nil
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
