package servererrors

import "errors"

var (
	RecordNotFound      = errors.New("record not found")
	InternalServerError = errors.New("internal server error")
)
