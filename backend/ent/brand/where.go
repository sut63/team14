// Code generated by entc, DO NOT EDIT.

package brand

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Brandname applies equality check predicate on the "Brandname" field. It's identical to BrandnameEQ.
func Brandname(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrandname), v))
	})
}

// BrandnameEQ applies the EQ predicate on the "Brandname" field.
func BrandnameEQ(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrandname), v))
	})
}

// BrandnameNEQ applies the NEQ predicate on the "Brandname" field.
func BrandnameNEQ(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrandname), v))
	})
}

// BrandnameIn applies the In predicate on the "Brandname" field.
func BrandnameIn(vs ...string) predicate.Brand {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Brand(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBrandname), v...))
	})
}

// BrandnameNotIn applies the NotIn predicate on the "Brandname" field.
func BrandnameNotIn(vs ...string) predicate.Brand {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Brand(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBrandname), v...))
	})
}

// BrandnameGT applies the GT predicate on the "Brandname" field.
func BrandnameGT(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBrandname), v))
	})
}

// BrandnameGTE applies the GTE predicate on the "Brandname" field.
func BrandnameGTE(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBrandname), v))
	})
}

// BrandnameLT applies the LT predicate on the "Brandname" field.
func BrandnameLT(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBrandname), v))
	})
}

// BrandnameLTE applies the LTE predicate on the "Brandname" field.
func BrandnameLTE(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBrandname), v))
	})
}

// BrandnameContains applies the Contains predicate on the "Brandname" field.
func BrandnameContains(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBrandname), v))
	})
}

// BrandnameHasPrefix applies the HasPrefix predicate on the "Brandname" field.
func BrandnameHasPrefix(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBrandname), v))
	})
}

// BrandnameHasSuffix applies the HasSuffix predicate on the "Brandname" field.
func BrandnameHasSuffix(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBrandname), v))
	})
}

// BrandnameEqualFold applies the EqualFold predicate on the "Brandname" field.
func BrandnameEqualFold(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBrandname), v))
	})
}

// BrandnameContainsFold applies the ContainsFold predicate on the "Brandname" field.
func BrandnameContainsFold(v string) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBrandname), v))
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFix applies the HasEdge predicate on the "fix" edge.
func HasFix() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixWith applies the HasEdge predicate on the "fix" edge with a given conditions (other predicates).
func HasFixWith(preds ...predicate.Fix) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		p(s.Not())
	})
}