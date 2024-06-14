// Code generated by ent, DO NOT EDIT.

package http

import (
	"github.com/google/uuid"
)

// Payload of a ent.User create request.
type UserCreateRequest struct {
	UUID *uuid.UUID `json:"uuid"`
}

// Payload of a ent.User update request.
type UserUpdateRequest struct {
	UUID *uuid.UUID `json:"uuid"`
}
