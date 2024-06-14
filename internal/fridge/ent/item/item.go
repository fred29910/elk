// Code generated by ent, DO NOT EDIT.

package item

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCompartment holds the string denoting the compartment edge name in mutations.
	EdgeCompartment = "compartment"
	// Table holds the table name of the item in the database.
	Table = "items"
	// CompartmentTable is the table that holds the compartment relation/edge.
	CompartmentTable = "items"
	// CompartmentInverseTable is the table name for the Compartment entity.
	// It exists in this package in order to avoid circular dependency with the "compartment" package.
	CompartmentInverseTable = "compartments"
	// CompartmentColumn is the table column denoting the compartment relation/edge.
	CompartmentColumn = "compartment_contents"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"compartment_contents",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Item queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCompartmentField orders the results by compartment field.
func ByCompartmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompartmentStep(), sql.OrderByField(field, opts...))
	}
}
func newCompartmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompartmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompartmentTable, CompartmentColumn),
	)
}
