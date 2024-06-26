// Code generated by ent, DO NOT EDIT.

package toy

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the toy type in the database.
	Label = "toy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldMaterial holds the string denoting the material field in the database.
	FieldMaterial = "material"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the toy in the database.
	Table = "toys"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "toys"
	// OwnerInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	OwnerInverseTable = "pets"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "pet_toys"
)

// Columns holds all SQL columns for toy fields.
var Columns = []string{
	FieldID,
	FieldColor,
	FieldMaterial,
	FieldTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "toys"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"pet_toys",
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

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Color defines the type for the "color" enum field.
type Color string

// Color values.
const (
	ColorRed    Color = "red"
	ColorOrange Color = "orange"
	ColorYellow Color = "yellow"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorIndigo Color = "indigo"
	ColorViolet Color = "violet"
	ColorPurple Color = "purple"
	ColorPink   Color = "pink"
	ColorSilver Color = "silver"
	ColorGold   Color = "gold"
	ColorBeige  Color = "beige"
	ColorBrown  Color = "brown"
	ColorGrey   Color = "grey"
	ColorBlack  Color = "black"
	ColorWhite  Color = "white"
)

func (c Color) String() string {
	return string(c)
}

// ColorValidator is a validator for the "color" field enum values. It is called by the builders before save.
func ColorValidator(c Color) error {
	switch c {
	case ColorRed, ColorOrange, ColorYellow, ColorGreen, ColorBlue, ColorIndigo, ColorViolet, ColorPurple, ColorPink, ColorSilver, ColorGold, ColorBeige, ColorBrown, ColorGrey, ColorBlack, ColorWhite:
		return nil
	default:
		return fmt.Errorf("toy: invalid enum value for color field: %q", c)
	}
}

// Material defines the type for the "material" enum field.
type Material string

// Material values.
const (
	MaterialLeather Material = "leather"
	MaterialPlastic Material = "plastic"
	MaterialFabric  Material = "fabric"
)

func (m Material) String() string {
	return string(m)
}

// MaterialValidator is a validator for the "material" field enum values. It is called by the builders before save.
func MaterialValidator(m Material) error {
	switch m {
	case MaterialLeather, MaterialPlastic, MaterialFabric:
		return nil
	default:
		return fmt.Errorf("toy: invalid enum value for material field: %q", m)
	}
}

// OrderOption defines the ordering options for the Toy queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByMaterial orders the results by the material field.
func ByMaterial(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaterial, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
