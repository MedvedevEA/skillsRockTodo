package dto

type GetUsers struct {
	Name   *string `query:"name" validate:"omitempty"`
	Offset int     `query:"offset" validate:"gte=0"`
	Limit  int     `query:"limit" validate:"gte=0"`
}
