package dto

import (
	"github.com/google/uuid"
)

type AddUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GetUsers struct {
	Offset int `query:"offset" validate:"gte=0"`
	Limit  int `query:"limit" validate:"gte=0"`
}
type GetUser struct {
	UserId *uuid.UUID `uri:"userId" validate:"required"`
}
type UpdateUser struct {
	UserId   *uuid.UUID `uri:"userId" validate:"required"`
	Name     *string    `json:"name" validate:"omitempty"`
	Password *string    `json:"password" validate:"omitempty"`
}
type RemoveUser struct {
	UserId *uuid.UUID `uri:"userId" validate:"required"`
}
