package repository

import (
	"skillsRockTodo/internal/entity"
)

type Repository interface {
	CreateTask(dto *DtoCreateTaskReq) error
	GetTasks() ([]*entity.Task, error)
	GetTask(Id int) (*entity.Task, error)
	UpdateTask(dto *DtoUpdateTaskReq) error
	DeleteTask(Id int) error
}
