package dto

import (
	"github.com/google/uuid"
)

type AddUserTask struct {
	UserId *uuid.UUID `uri:"userId" validate:"required"`
	TaskId *uuid.UUID `uri:"taskId" validate:"required"`
}

type GetUserTasks struct {
	Offset int        `query:"offset" validate:"gte=0"`
	Limit  int        `query:"limit" validate:"gte=0"`
	UserId *uuid.UUID `query:"userId"`
	TaskId *uuid.UUID `query:"taskId"`
}
type RemoveUserTask struct {
	UserTaskId *uuid.UUID `uri:"userTaskId" validate:"required"`
}
