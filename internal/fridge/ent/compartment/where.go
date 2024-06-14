// Code generated by ent, DO NOT EDIT.

package compartment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/masseelch/elk/internal/fridge/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Compartment {
	return predicate.Compartment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Compartment {
	return predicate.Compartment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Compartment {
	return predicate.Compartment(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Compartment {
	return predicate.Compartment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Compartment {
	return predicate.Compartment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Compartment {
	return predicate.Compartment(sql.FieldContainsFold(FieldName, v))
}

// HasFridge applies the HasEdge predicate on the "fridge" edge.
func HasFridge() predicate.Compartment {
	return predicate.Compartment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FridgeTable, FridgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFridgeWith applies the HasEdge predicate on the "fridge" edge with a given conditions (other predicates).
func HasFridgeWith(preds ...predicate.Fridge) predicate.Compartment {
	return predicate.Compartment(func(s *sql.Selector) {
		step := newFridgeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContents applies the HasEdge predicate on the "contents" edge.
func HasContents() predicate.Compartment {
	return predicate.Compartment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContentsTable, ContentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContentsWith applies the HasEdge predicate on the "contents" edge with a given conditions (other predicates).
func HasContentsWith(preds ...predicate.Item) predicate.Compartment {
	return predicate.Compartment(func(s *sql.Selector) {
		step := newContentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Compartment) predicate.Compartment {
	return predicate.Compartment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Compartment) predicate.Compartment {
	return predicate.Compartment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Compartment) predicate.Compartment {
	return predicate.Compartment(sql.NotPredicates(p))
}
