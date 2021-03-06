// Code generated by entc, DO NOT EDIT.

package fix

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Productnumber applies equality check predicate on the "Productnumber" field. It's identical to ProductnumberEQ.
func Productnumber(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductnumber), v))
	})
}

// Problemtype applies equality check predicate on the "Problemtype" field. It's identical to ProblemtypeEQ.
func Problemtype(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProblemtype), v))
	})
}

// Queue applies equality check predicate on the "Queue" field. It's identical to QueueEQ.
func Queue(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQueue), v))
	})
}

// Date applies equality check predicate on the "Date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// ProductnumberEQ applies the EQ predicate on the "Productnumber" field.
func ProductnumberEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductnumber), v))
	})
}

// ProductnumberNEQ applies the NEQ predicate on the "Productnumber" field.
func ProductnumberNEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductnumber), v))
	})
}

// ProductnumberIn applies the In predicate on the "Productnumber" field.
func ProductnumberIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductnumber), v...))
	})
}

// ProductnumberNotIn applies the NotIn predicate on the "Productnumber" field.
func ProductnumberNotIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductnumber), v...))
	})
}

// ProductnumberGT applies the GT predicate on the "Productnumber" field.
func ProductnumberGT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductnumber), v))
	})
}

// ProductnumberGTE applies the GTE predicate on the "Productnumber" field.
func ProductnumberGTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductnumber), v))
	})
}

// ProductnumberLT applies the LT predicate on the "Productnumber" field.
func ProductnumberLT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductnumber), v))
	})
}

// ProductnumberLTE applies the LTE predicate on the "Productnumber" field.
func ProductnumberLTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductnumber), v))
	})
}

// ProductnumberContains applies the Contains predicate on the "Productnumber" field.
func ProductnumberContains(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProductnumber), v))
	})
}

// ProductnumberHasPrefix applies the HasPrefix predicate on the "Productnumber" field.
func ProductnumberHasPrefix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProductnumber), v))
	})
}

// ProductnumberHasSuffix applies the HasSuffix predicate on the "Productnumber" field.
func ProductnumberHasSuffix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProductnumber), v))
	})
}

// ProductnumberEqualFold applies the EqualFold predicate on the "Productnumber" field.
func ProductnumberEqualFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProductnumber), v))
	})
}

// ProductnumberContainsFold applies the ContainsFold predicate on the "Productnumber" field.
func ProductnumberContainsFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProductnumber), v))
	})
}

// ProblemtypeEQ applies the EQ predicate on the "Problemtype" field.
func ProblemtypeEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeNEQ applies the NEQ predicate on the "Problemtype" field.
func ProblemtypeNEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeIn applies the In predicate on the "Problemtype" field.
func ProblemtypeIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProblemtype), v...))
	})
}

// ProblemtypeNotIn applies the NotIn predicate on the "Problemtype" field.
func ProblemtypeNotIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProblemtype), v...))
	})
}

// ProblemtypeGT applies the GT predicate on the "Problemtype" field.
func ProblemtypeGT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeGTE applies the GTE predicate on the "Problemtype" field.
func ProblemtypeGTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeLT applies the LT predicate on the "Problemtype" field.
func ProblemtypeLT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeLTE applies the LTE predicate on the "Problemtype" field.
func ProblemtypeLTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeContains applies the Contains predicate on the "Problemtype" field.
func ProblemtypeContains(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeHasPrefix applies the HasPrefix predicate on the "Problemtype" field.
func ProblemtypeHasPrefix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeHasSuffix applies the HasSuffix predicate on the "Problemtype" field.
func ProblemtypeHasSuffix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeEqualFold applies the EqualFold predicate on the "Problemtype" field.
func ProblemtypeEqualFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProblemtype), v))
	})
}

// ProblemtypeContainsFold applies the ContainsFold predicate on the "Problemtype" field.
func ProblemtypeContainsFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProblemtype), v))
	})
}

// QueueEQ applies the EQ predicate on the "Queue" field.
func QueueEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQueue), v))
	})
}

// QueueNEQ applies the NEQ predicate on the "Queue" field.
func QueueNEQ(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQueue), v))
	})
}

// QueueIn applies the In predicate on the "Queue" field.
func QueueIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQueue), v...))
	})
}

// QueueNotIn applies the NotIn predicate on the "Queue" field.
func QueueNotIn(vs ...string) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQueue), v...))
	})
}

// QueueGT applies the GT predicate on the "Queue" field.
func QueueGT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQueue), v))
	})
}

// QueueGTE applies the GTE predicate on the "Queue" field.
func QueueGTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQueue), v))
	})
}

// QueueLT applies the LT predicate on the "Queue" field.
func QueueLT(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQueue), v))
	})
}

// QueueLTE applies the LTE predicate on the "Queue" field.
func QueueLTE(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQueue), v))
	})
}

// QueueContains applies the Contains predicate on the "Queue" field.
func QueueContains(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldQueue), v))
	})
}

// QueueHasPrefix applies the HasPrefix predicate on the "Queue" field.
func QueueHasPrefix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldQueue), v))
	})
}

// QueueHasSuffix applies the HasSuffix predicate on the "Queue" field.
func QueueHasSuffix(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldQueue), v))
	})
}

// QueueEqualFold applies the EqualFold predicate on the "Queue" field.
func QueueEqualFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldQueue), v))
	})
}

// QueueContainsFold applies the ContainsFold predicate on the "Queue" field.
func QueueContainsFold(v string) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldQueue), v))
	})
}

// DateEQ applies the EQ predicate on the "Date" field.
func DateEQ(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "Date" field.
func DateNEQ(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "Date" field.
func DateIn(vs ...time.Time) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "Date" field.
func DateNotIn(vs ...time.Time) predicate.Fix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "Date" field.
func DateGT(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "Date" field.
func DateGTE(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "Date" field.
func DateLT(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "Date" field.
func DateLTE(v time.Time) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasFix applies the HasEdge predicate on the "fix" edge.
func HasFix() predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixWith applies the HasEdge predicate on the "fix" edge with a given conditions (other predicates).
func HasFixWith(preds ...predicate.Adminrepair) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixbrand applies the HasEdge predicate on the "fixbrand" edge.
func HasFixbrand() predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixbrandTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixbrandTable, FixbrandColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixbrandWith applies the HasEdge predicate on the "fixbrand" edge with a given conditions (other predicates).
func HasFixbrandWith(preds ...predicate.Fixbrand) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixbrandInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixbrandTable, FixbrandColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Personal) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixcomtype applies the HasEdge predicate on the "fixcomtype" edge.
func HasFixcomtype() predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixcomtypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixcomtypeTable, FixcomtypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixcomtypeWith applies the HasEdge predicate on the "fixcomtype" edge with a given conditions (other predicates).
func HasFixcomtypeWith(preds ...predicate.Fixcomtype) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixcomtypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixcomtypeTable, FixcomtypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Fix) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Fix) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
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
func Not(p predicate.Fix) predicate.Fix {
	return predicate.Fix(func(s *sql.Selector) {
		p(s.Not())
	})
}
