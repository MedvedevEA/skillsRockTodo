package dto

import "github.com/google/uuid"

type AddTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetTasks struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
type UpdateTask struct {
	TaskId      *uuid.UUID `json:"task_id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
}
