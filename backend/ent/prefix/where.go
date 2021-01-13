// Code generated by entc, DO NOT EDIT.

package prefix

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PrefixValue applies equality check predicate on the "prefixValue" field. It's identical to PrefixValueEQ.
func PrefixValue(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueEQ applies the EQ predicate on the "prefixValue" field.
func PrefixValueEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueNEQ applies the NEQ predicate on the "prefixValue" field.
func PrefixValueNEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueIn applies the In predicate on the "prefixValue" field.
func PrefixValueIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrefixValue), v...))
	})
}

// PrefixValueNotIn applies the NotIn predicate on the "prefixValue" field.
func PrefixValueNotIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrefixValue), v...))
	})
}

// PrefixValueGT applies the GT predicate on the "prefixValue" field.
func PrefixValueGT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueGTE applies the GTE predicate on the "prefixValue" field.
func PrefixValueGTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueLT applies the LT predicate on the "prefixValue" field.
func PrefixValueLT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueLTE applies the LTE predicate on the "prefixValue" field.
func PrefixValueLTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueContains applies the Contains predicate on the "prefixValue" field.
func PrefixValueContains(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueHasPrefix applies the HasPrefix predicate on the "prefixValue" field.
func PrefixValueHasPrefix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueHasSuffix applies the HasSuffix predicate on the "prefixValue" field.
func PrefixValueHasSuffix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueEqualFold applies the EqualFold predicate on the "prefixValue" field.
func PrefixValueEqualFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrefixValue), v))
	})
}

// PrefixValueContainsFold applies the ContainsFold predicate on the "prefixValue" field.
func PrefixValueContainsFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrefixValue), v))
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
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
func Not(p predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		p(s.Not())
	})
}
