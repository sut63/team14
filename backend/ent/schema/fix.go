package schema

import "github.com/facebookincubator/ent"

// Fix holds the schema definition for the Fix entity.
type Fix struct {
	ent.Schema
}

// Fields of the Fix.
func (Fix) Fields() []ent.Field {
	return nil
}

// Edges of the Fix.
func (Fix) Edges() []ent.Edge {
	return nil
}
