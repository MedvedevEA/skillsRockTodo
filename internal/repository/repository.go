package repository

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"
)

type Repository interface {
	AddTask(dto *dto.AddTask) (*entity.Task, error)
	GetTasks() ([]*entity.Task, error)
	GetTask(Id int) (*entity.Task, error)
	UpdateTask(dto *dto.UpdateTask) error
	RemoveTask(Id int) error
}
