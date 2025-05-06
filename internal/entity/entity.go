package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageId *uuid.UUID `json:"message_id"`
	TaskId    *uuid.UUID `json:"task_id"`
	UserId    *uuid.UUID `json:"user_id"`
	Text      string     `json:"text"`
	CreateAt  time.Time  `json:"create_at"`
	UpdateAt  time.Time  `json:"update_at"`
}
type Status struct {
	StatusId *uuid.UUID `json:"status_id"`
	Name     string     `json:"name"`
}
type Task struct {
	TaskId      *uuid.UUID `json:"task_id"`
	StatusId    *uuid.UUID `json:"status_is"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}
type TaskUser struct {
	TaskUserId *uuid.UUID `json:"task_user_id"`
	TaskId     *uuid.UUID `json:"task_id"`
	UserId     *uuid.UUID `json:"user_id"`
}
type User struct {
	UserId *uuid.UUID `json:"user_id"`
	Name   string     `json:"name"`
}
