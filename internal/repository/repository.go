package repository

import (
	"skillsRockTodo/internal/entity"
)

type Repository interface {
	CreateTask(dto *dtoCreateTaskReq) error
	GetTasks() ([]*entity.Task, error)
	GetTask(Id int) (*entity.Task, error)
	UpdateTask(dto *dtoUpdateTaskReq) error
	DeleteTask(Id int) error
}
