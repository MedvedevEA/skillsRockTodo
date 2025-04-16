package dto

import "github.com/google/uuid"

type AddUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type GetUsers struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Name   string `json:"name"`
}
type UpdateUser struct {
	UserId   *uuid.UUID `json:"user_id"`
	Name     *string    `json:"name"`
	Password *string    `json:"password"`
}
