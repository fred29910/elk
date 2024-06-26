// Code generated by ent, DO NOT EDIT.

package compartment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the compartment type in the database.
	Label = "compartment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFridge holds the string denoting the fridge edge name in mutations.
	EdgeFridge = "fridge"
	// EdgeContents holds the string denoting the contents edge name in mutations.
	EdgeContents = "contents"
	// Table holds the table name of the compartment in the database.
	Table = "compartments"
	// FridgeTable is the table that holds the fridge relation/edge.
	FridgeTable = "compartments"
	// FridgeInverseTable is the table name for the Fridge entity.
	// It exists in this package in order to avoid circular dependency with the "fridge" package.
	FridgeInverseTable = "fridges"
	// FridgeColumn is the table column denoting the fridge relation/edge.
	FridgeColumn = "fridge_compartments"
	// ContentsTable is the table that holds the contents relation/edge.
	ContentsTable = "items"
	// ContentsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ContentsInverseTable = "items"
	// ContentsColumn is the table column denoting the contents relation/edge.
	ContentsColumn = "compartment_contents"
)

// Columns holds all SQL columns for compartment fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "compartments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"fridge_compartments",
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

// OrderOption defines the ordering options for the Compartment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFridgeField orders the results by fridge field.
func ByFridgeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFridgeStep(), sql.OrderByField(field, opts...))
	}
}

// ByContentsCount orders the results by contents count.
func ByContentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContentsStep(), opts...)
	}
}

// ByContents orders the results by contents terms.
func ByContents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFridgeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FridgeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FridgeTable, FridgeColumn),
	)
}
func newContentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContentsTable, ContentsColumn),
	)
}
