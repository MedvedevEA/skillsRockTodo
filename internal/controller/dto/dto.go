package dto

import "github.com/google/uuid"

type AddTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetTasks struct {
	Offset int `query:"offset" validate:"gte=0"`
	Limit  int `query:"limit" validate:"gte=0"`
}
type GetTask struct {
	TaskId *uuid.UUID `uri:"taskId" validate:"required"`
}
type UpdateTask struct {
	TaskId      *uuid.UUID `uri:"taskId" validate:"required"`
	Title       *string    `json:"title" validate:"omitempty"`
	Description *string    `json:"description" validate:"omitempty"`
	Status      *string    `json:"status" validate:"omitempty,oneof=new in_progress done"`
}
type RemoveTask struct {
	TaskId *uuid.UUID `uri:"taskId" validate:"required"`
}
