package repository

import (
	"skillsRockTodo/internal/entity"
)

type Repository interface {
	CreateTask(task *entity.Task) error
	GetTasks() ([]*entity.Task, error)
	GetTask(taskId int) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTask(taskId int) error
}
