// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/masseelch/elk/internal/client_gen/ent/collar"
	"github.com/masseelch/elk/internal/client_gen/ent/owner"
	"github.com/masseelch/elk/internal/client_gen/ent/pet"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges        PetEdges `json:"edges"`
	owner_pets   *uuid.UUID
	selectValues sql.SelectValues
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Collar holds the value of the collar edge.
	Collar *Collar `json:"collar,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Owner `json:"owner,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*Pet `json:"friends,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CollarOrErr returns the Collar value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) CollarOrErr() (*Collar, error) {
	if e.Collar != nil {
		return e.Collar, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: collar.Label}
	}
	return nil, &NotLoadedError{edge: "collar"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*Owner, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: owner.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) FriendsOrErr() ([]*Pet, error) {
	if e.loadedTypes[3] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldAge:
			values[i] = new(sql.NullInt64)
		case pet.FieldID, pet.FieldName:
			values[i] = new(sql.NullString)
		case pet.ForeignKeys[0]: // owner_pets
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pe.ID = value.String
			}
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case pet.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pe.Age = int(value.Int64)
			}
		case pet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field owner_pets", values[i])
			} else if value.Valid {
				pe.owner_pets = new(uuid.UUID)
				*pe.owner_pets = *value.S.(*uuid.UUID)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pet.
// This includes values selected through modifiers, order, etc.
func (pe *Pet) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryCollar queries the "collar" edge of the Pet entity.
func (pe *Pet) QueryCollar() *CollarQuery {
	return NewPetClient(pe.config).QueryCollar(pe)
}

// QueryCategories queries the "categories" edge of the Pet entity.
func (pe *Pet) QueryCategories() *CategoryQuery {
	return NewPetClient(pe.config).QueryCategories(pe)
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *OwnerQuery {
	return NewPetClient(pe.config).QueryOwner(pe)
}

// QueryFriends queries the "friends" edge of the Pet entity.
func (pe *Pet) QueryFriends() *PetQuery {
	return NewPetClient(pe.config).QueryFriends(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return NewPetClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet
