package schema

import "github.com/facebookincubator/ent"

// Receipt holds the schema definition for the Receipt entity.
type Receipt struct {
	ent.Schema
}

// Fields of the Receipt.
func (Receipt) Fields() []ent.Field {
	return nil
}

// Edges of the Receipt.
func (Receipt) Edges() []ent.Edge {
	return nil
}
