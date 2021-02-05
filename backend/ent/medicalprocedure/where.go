// Code generated by entc, DO NOT EDIT.

package medicalprocedure

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
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
func IDGT(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProcedureOrder applies equality check predicate on the "procedureOrder" field. It's identical to ProcedureOrderEQ.
func ProcedureOrder(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureRoom applies equality check predicate on the "procedureRoom" field. It's identical to ProcedureRoomEQ.
func ProcedureRoom(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureRoom), v))
	})
}

// Addtime applies equality check predicate on the "Addtime" field. It's identical to AddtimeEQ.
func Addtime(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// ProcedureDescripe applies equality check predicate on the "procedureDescripe" field. It's identical to ProcedureDescripeEQ.
func ProcedureDescripe(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureOrderEQ applies the EQ predicate on the "procedureOrder" field.
func ProcedureOrderEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderNEQ applies the NEQ predicate on the "procedureOrder" field.
func ProcedureOrderNEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderIn applies the In predicate on the "procedureOrder" field.
func ProcedureOrderIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProcedureOrder), v...))
	})
}

// ProcedureOrderNotIn applies the NotIn predicate on the "procedureOrder" field.
func ProcedureOrderNotIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProcedureOrder), v...))
	})
}

// ProcedureOrderGT applies the GT predicate on the "procedureOrder" field.
func ProcedureOrderGT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderGTE applies the GTE predicate on the "procedureOrder" field.
func ProcedureOrderGTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderLT applies the LT predicate on the "procedureOrder" field.
func ProcedureOrderLT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderLTE applies the LTE predicate on the "procedureOrder" field.
func ProcedureOrderLTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderContains applies the Contains predicate on the "procedureOrder" field.
func ProcedureOrderContains(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderHasPrefix applies the HasPrefix predicate on the "procedureOrder" field.
func ProcedureOrderHasPrefix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderHasSuffix applies the HasSuffix predicate on the "procedureOrder" field.
func ProcedureOrderHasSuffix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderEqualFold applies the EqualFold predicate on the "procedureOrder" field.
func ProcedureOrderEqualFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureOrderContainsFold applies the ContainsFold predicate on the "procedureOrder" field.
func ProcedureOrderContainsFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProcedureOrder), v))
	})
}

// ProcedureRoomEQ applies the EQ predicate on the "procedureRoom" field.
func ProcedureRoomEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomNEQ applies the NEQ predicate on the "procedureRoom" field.
func ProcedureRoomNEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomIn applies the In predicate on the "procedureRoom" field.
func ProcedureRoomIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProcedureRoom), v...))
	})
}

// ProcedureRoomNotIn applies the NotIn predicate on the "procedureRoom" field.
func ProcedureRoomNotIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProcedureRoom), v...))
	})
}

// ProcedureRoomGT applies the GT predicate on the "procedureRoom" field.
func ProcedureRoomGT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomGTE applies the GTE predicate on the "procedureRoom" field.
func ProcedureRoomGTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomLT applies the LT predicate on the "procedureRoom" field.
func ProcedureRoomLT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomLTE applies the LTE predicate on the "procedureRoom" field.
func ProcedureRoomLTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomContains applies the Contains predicate on the "procedureRoom" field.
func ProcedureRoomContains(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomHasPrefix applies the HasPrefix predicate on the "procedureRoom" field.
func ProcedureRoomHasPrefix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomHasSuffix applies the HasSuffix predicate on the "procedureRoom" field.
func ProcedureRoomHasSuffix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomEqualFold applies the EqualFold predicate on the "procedureRoom" field.
func ProcedureRoomEqualFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProcedureRoom), v))
	})
}

// ProcedureRoomContainsFold applies the ContainsFold predicate on the "procedureRoom" field.
func ProcedureRoomContainsFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProcedureRoom), v))
	})
}

// AddtimeEQ applies the EQ predicate on the "Addtime" field.
func AddtimeEQ(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddtime), v))
	})
}

// AddtimeNEQ applies the NEQ predicate on the "Addtime" field.
func AddtimeNEQ(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddtime), v))
	})
}

// AddtimeIn applies the In predicate on the "Addtime" field.
func AddtimeIn(vs ...time.Time) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddtime), v...))
	})
}

// AddtimeNotIn applies the NotIn predicate on the "Addtime" field.
func AddtimeNotIn(vs ...time.Time) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddtime), v...))
	})
}

// AddtimeGT applies the GT predicate on the "Addtime" field.
func AddtimeGT(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddtime), v))
	})
}

// AddtimeGTE applies the GTE predicate on the "Addtime" field.
func AddtimeGTE(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddtime), v))
	})
}

// AddtimeLT applies the LT predicate on the "Addtime" field.
func AddtimeLT(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddtime), v))
	})
}

// AddtimeLTE applies the LTE predicate on the "Addtime" field.
func AddtimeLTE(v time.Time) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddtime), v))
	})
}

// ProcedureDescripeEQ applies the EQ predicate on the "procedureDescripe" field.
func ProcedureDescripeEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeNEQ applies the NEQ predicate on the "procedureDescripe" field.
func ProcedureDescripeNEQ(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeIn applies the In predicate on the "procedureDescripe" field.
func ProcedureDescripeIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProcedureDescripe), v...))
	})
}

// ProcedureDescripeNotIn applies the NotIn predicate on the "procedureDescripe" field.
func ProcedureDescripeNotIn(vs ...string) predicate.MedicalProcedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProcedureDescripe), v...))
	})
}

// ProcedureDescripeGT applies the GT predicate on the "procedureDescripe" field.
func ProcedureDescripeGT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeGTE applies the GTE predicate on the "procedureDescripe" field.
func ProcedureDescripeGTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeLT applies the LT predicate on the "procedureDescripe" field.
func ProcedureDescripeLT(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeLTE applies the LTE predicate on the "procedureDescripe" field.
func ProcedureDescripeLTE(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeContains applies the Contains predicate on the "procedureDescripe" field.
func ProcedureDescripeContains(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeHasPrefix applies the HasPrefix predicate on the "procedureDescripe" field.
func ProcedureDescripeHasPrefix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeHasSuffix applies the HasSuffix predicate on the "procedureDescripe" field.
func ProcedureDescripeHasSuffix(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeEqualFold applies the EqualFold predicate on the "procedureDescripe" field.
func ProcedureDescripeEqualFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProcedureDescripe), v))
	})
}

// ProcedureDescripeContainsFold applies the ContainsFold predicate on the "procedureDescripe" field.
func ProcedureDescripeContainsFold(v string) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProcedureDescripe), v))
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProcedureType applies the HasEdge predicate on the "ProcedureType" edge.
func HasProcedureType() predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProcedureTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProcedureTypeTable, ProcedureTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProcedureTypeWith applies the HasEdge predicate on the "ProcedureType" edge with a given conditions (other predicates).
func HasProcedureTypeWith(preds ...predicate.ProcedureType) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProcedureTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProcedureTypeTable, ProcedureTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDoctor applies the HasEdge predicate on the "Doctor" edge.
func HasDoctor() predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorTable, DoctorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorWith applies the HasEdge predicate on the "Doctor" edge with a given conditions (other predicates).
func HasDoctorWith(preds ...predicate.Doctor) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorTable, DoctorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MedicalProcedure) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MedicalProcedure) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
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
func Not(p predicate.MedicalProcedure) predicate.MedicalProcedure {
	return predicate.MedicalProcedure(func(s *sql.Selector) {
		p(s.Not())
	})
}
