// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/simple/ent"
	collar "github.com/masseelch/elk/internal/simple/ent/collar"
)

// Basic HTTP Error Response
type ErrResponse struct {
	Code   int         `json:"code"`             // http response status code
	Status string      `json:"status"`           // user-level status message
	Errors interface{} `json:"errors,omitempty"` // application-level error
}

func (e ErrResponse) MarshalToHTTPResponseWriter(w http.ResponseWriter) (int, error) {
	d, err := easyjson.Marshal(e)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(d)))
	w.WriteHeader(e.Code)
	return w.Write(d)
}

func BadRequest(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Conflict(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusConflict,
		Status: http.StatusText(http.StatusConflict),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Forbidden(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func InternalServerError(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func NotFound(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

type (
	// Category4094953247View represents the data serialized for the following serialization group combinations:
	// []
	// [owner pet pet:owner]
	Category4094953247View struct {
		ID   uint64 `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	Category4094953247Views []*Category4094953247View
)

func NewCategory4094953247View(e *ent.Category) *Category4094953247View {
	if e == nil {
		return nil
	}
	return &Category4094953247View{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewCategory4094953247Views(es []*ent.Category) Category4094953247Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Category4094953247Views, len(es))
	for i, e := range es {
		r[i] = NewCategory4094953247View(e)
	}
	return r
}

type (
	// Collar1522160880View represents the data serialized for the following serialization group combinations:
	// []
	// [owner pet pet:owner]
	Collar1522160880View struct {
		ID    int          `json:"id,omitempty"`
		Color collar.Color `json:"color,omitempty"`
	}
	Collar1522160880Views []*Collar1522160880View
)

func NewCollar1522160880View(e *ent.Collar) *Collar1522160880View {
	if e == nil {
		return nil
	}
	return &Collar1522160880View{
		ID:    e.ID,
		Color: e.Color,
	}
}

func NewCollar1522160880Views(es []*ent.Collar) Collar1522160880Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Collar1522160880Views, len(es))
	for i, e := range es {
		r[i] = NewCollar1522160880View(e)
	}
	return r
}

type (
	// Owner139708381View represents the data serialized for the following serialization group combinations:
	// []
	// [owner pet pet:owner]
	Owner139708381View struct {
		ID   uuid.UUID `json:"id,omitempty"`
		Name string    `json:"name,omitempty"`
		Age  int       `json:"age,omitempty"`
	}
	Owner139708381Views []*Owner139708381View
)

func NewOwner139708381View(e *ent.Owner) *Owner139708381View {
	if e == nil {
		return nil
	}
	return &Owner139708381View{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewOwner139708381Views(es []*ent.Owner) Owner139708381Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Owner139708381Views, len(es))
	for i, e := range es {
		r[i] = NewOwner139708381View(e)
	}
	return r
}

type (
	// Pet1876743790View represents the data serialized for the following serialization group combinations:
	// [owner pet pet:owner]
	Pet1876743790View struct {
		ID      string              `json:"id,omitempty"`
		Name    string              `json:"name,omitempty"`
		Age     int                 `json:"age,omitempty"`
		Owner   *Owner139708381View `json:"owner,omitempty"`
		Friends Pet359800019Views   `json:"friends,omitempty"`
	}
	Pet1876743790Views []*Pet1876743790View
)

func NewPet1876743790View(e *ent.Pet) *Pet1876743790View {
	if e == nil {
		return nil
	}
	return &Pet1876743790View{
		ID:      e.ID,
		Name:    e.Name,
		Age:     e.Age,
		Owner:   NewOwner139708381View(e.Edges.Owner),
		Friends: NewPet359800019Views(e.Edges.Friends),
	}
}

func NewPet1876743790Views(es []*ent.Pet) Pet1876743790Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Pet1876743790Views, len(es))
	for i, e := range es {
		r[i] = NewPet1876743790View(e)
	}
	return r
}

type (
	// Pet359800019View represents the data serialized for the following serialization group combinations:
	// []
	Pet359800019View struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}
	Pet359800019Views []*Pet359800019View
)

func NewPet359800019View(e *ent.Pet) *Pet359800019View {
	if e == nil {
		return nil
	}
	return &Pet359800019View{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewPet359800019Views(es []*ent.Pet) Pet359800019Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Pet359800019Views, len(es))
	for i, e := range es {
		r[i] = NewPet359800019View(e)
	}
	return r
}
