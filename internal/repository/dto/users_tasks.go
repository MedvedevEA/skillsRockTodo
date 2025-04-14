package dto

import "github.com/google/uuid"

type AddUserTask struct {
	UserId *uuid.UUID `json:"user_id"`
	TaskId *uuid.UUID `json:"task_id"`
}
type GetUserTasks struct {
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
	UserId   *uuid.UUID `json:"user_id"`
	TaskId   *uuid.UUID `json:"task_id"`
}
