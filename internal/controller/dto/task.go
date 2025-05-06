package dto

import "github.com/google/uuid"

type AddTask struct {
	StatusId    *uuid.UUID `json:"status_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}
type GetTask struct {
	TaskId *uuid.UUID `uri:"taskId" validate:"required"`
}
type GetTasks struct {
	Offset   int        `query:"offset" validate:"gte=0"`
	Limit    int        `query:"limit" validate:"gte=0"`
	StatusId *uuid.UUID `json:"status_id"`
}
type UpdateTask struct {
	TaskId      *uuid.UUID `uri:"taskId" validate:"required"`
	StatusId    *uuid.UUID `json:"status_id" validate:"omitempty"`
	Title       *string    `json:"title" validate:"omitempty"`
	Description *string    `json:"description" validate:"omitempty"`
}
type RemoveTask struct {
	TaskId *uuid.UUID `uri:"taskId" validate:"required"`
}
