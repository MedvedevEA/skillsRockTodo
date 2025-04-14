package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskId      *uuid.UUID `json:"task_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
}

type User struct {
	UserId   *uuid.UUID `json:"user_id"`
	Name     string     `json:"name"`
	Password string     `json:"password"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
}

type UserTask struct {
	UserTaskId *uuid.UUID `json:"user_task_id"`
	UserId     *uuid.UUID `json:"user_id"`
	TaskId     *uuid.UUID `json:"task_id"`
	CreateAt   time.Time  `json:"create_at"`
	UpdateAt   time.Time  `json:"update_at"`
}

type AccessPermission struct {
	AccessPermissionId *uuid.UUID `json:"access_permission_id"`
	UserId             *uuid.UUID `json:"user_id"`
	Route              string     `json:"route"`
	CreateAt           time.Time  `json:"create_at"`
	UpdateAt           time.Time  `json:"update_at"`
}
