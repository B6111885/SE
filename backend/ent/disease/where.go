// Code generated by entc, DO NOT EDIT.

package disease

import (
	"github.com/ADMIN/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DiseaseName applies equality check predicate on the "Disease_Name" field. It's identical to DiseaseNameEQ.
func DiseaseName(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameEQ applies the EQ predicate on the "Disease_Name" field.
func DiseaseNameEQ(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameNEQ applies the NEQ predicate on the "Disease_Name" field.
func DiseaseNameNEQ(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameIn applies the In predicate on the "Disease_Name" field.
func DiseaseNameIn(vs ...string) predicate.Disease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiseaseName), v...))
	})
}

// DiseaseNameNotIn applies the NotIn predicate on the "Disease_Name" field.
func DiseaseNameNotIn(vs ...string) predicate.Disease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Disease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiseaseName), v...))
	})
}

// DiseaseNameGT applies the GT predicate on the "Disease_Name" field.
func DiseaseNameGT(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameGTE applies the GTE predicate on the "Disease_Name" field.
func DiseaseNameGTE(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameLT applies the LT predicate on the "Disease_Name" field.
func DiseaseNameLT(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameLTE applies the LTE predicate on the "Disease_Name" field.
func DiseaseNameLTE(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameContains applies the Contains predicate on the "Disease_Name" field.
func DiseaseNameContains(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameHasPrefix applies the HasPrefix predicate on the "Disease_Name" field.
func DiseaseNameHasPrefix(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameHasSuffix applies the HasSuffix predicate on the "Disease_Name" field.
func DiseaseNameHasSuffix(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameEqualFold applies the EqualFold predicate on the "Disease_Name" field.
func DiseaseNameEqualFold(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiseaseName), v))
	})
}

// DiseaseNameContainsFold applies the ContainsFold predicate on the "Disease_Name" field.
func DiseaseNameContainsFold(v string) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiseaseName), v))
	})
}

// HasDiseaseDiagnose applies the HasEdge predicate on the "disease_diagnose" edge.
func HasDiseaseDiagnose() predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiseaseDiagnoseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DiseaseDiagnoseTable, DiseaseDiagnoseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiseaseDiagnoseWith applies the HasEdge predicate on the "disease_diagnose" edge with a given conditions (other predicates).
func HasDiseaseDiagnoseWith(preds ...predicate.Diagnose) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiseaseDiagnoseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DiseaseDiagnoseTable, DiseaseDiagnoseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
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
func Not(p predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		p(s.Not())
	})
}
