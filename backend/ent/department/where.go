// Code generated by entc, DO NOT EDIT.

package department

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DepartmentName applies equality check predicate on the "departmentName" field. It's identical to DepartmentNameEQ.
func DepartmentName(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameEQ applies the EQ predicate on the "departmentName" field.
func DepartmentNameEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameNEQ applies the NEQ predicate on the "departmentName" field.
func DepartmentNameNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameIn applies the In predicate on the "departmentName" field.
func DepartmentNameIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepartmentName), v...))
	})
}

// DepartmentNameNotIn applies the NotIn predicate on the "departmentName" field.
func DepartmentNameNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepartmentName), v...))
	})
}

// DepartmentNameGT applies the GT predicate on the "departmentName" field.
func DepartmentNameGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameGTE applies the GTE predicate on the "departmentName" field.
func DepartmentNameGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameLT applies the LT predicate on the "departmentName" field.
func DepartmentNameLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameLTE applies the LTE predicate on the "departmentName" field.
func DepartmentNameLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameContains applies the Contains predicate on the "departmentName" field.
func DepartmentNameContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameHasPrefix applies the HasPrefix predicate on the "departmentName" field.
func DepartmentNameHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameHasSuffix applies the HasSuffix predicate on the "departmentName" field.
func DepartmentNameHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameEqualFold applies the EqualFold predicate on the "departmentName" field.
func DepartmentNameEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDepartmentName), v))
	})
}

// DepartmentNameContainsFold applies the ContainsFold predicate on the "departmentName" field.
func DepartmentNameContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDepartmentName), v))
	})
}

// HasTriageResult applies the HasEdge predicate on the "triageResult" edge.
func HasTriageResult() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TriageResultTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TriageResultTable, TriageResultColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTriageResultWith applies the HasEdge predicate on the "triageResult" edge with a given conditions (other predicates).
func HasTriageResultWith(preds ...predicate.TriageResult) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TriageResultInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TriageResultTable, TriageResultColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		p(s.Not())
	})
}
