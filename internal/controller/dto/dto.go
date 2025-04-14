package dto

import "github.com/google/uuid"

type AddTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type GetTasks struct {
	Page     int `validate:"gte=0,default:0"`
	PageSize int `validate:"gte=0,default:10"`
}
type GetTask struct {
	TaskId *uuid.UUID `validate:"required"`
}
type UpdateTask struct {
	TaskId      *uuid.UUID `validate:"required"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status" validate:"omitempty,oneof=new in_progress done"`
}
type RemoveTask struct {
	TaskId *uuid.UUID `validate:"required"`
}
