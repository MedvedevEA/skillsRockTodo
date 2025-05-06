package dto

import "github.com/google/uuid"

type AddTask struct {
	StatusId    *uuid.UUID `json:"status_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}
type GetTasks struct {
	Offset   int        `json:"offset"`
	Limit    int        `json:"limit"`
	StatusId *uuid.UUID `json:"status_id"`
}
type UpdateTask struct {
	TaskId      *uuid.UUID `json:"task_id"`
	StatusId    *uuid.UUID `json:"status_id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
}
