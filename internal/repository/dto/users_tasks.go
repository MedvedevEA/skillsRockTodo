package dto

import "github.com/google/uuid"

type AddUserTask struct {
	UserId *uuid.UUID `json:"user_id"`
	TaskId *uuid.UUID `json:"task_id"`
}
type GetUserTasks struct {
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
	UserId *uuid.UUID `json:"user_id"`
	TaskId *uuid.UUID `json:"task_id"`
}
