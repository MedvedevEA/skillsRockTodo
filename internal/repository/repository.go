package repository

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"

	"github.com/google/uuid"
)

type Repository interface {
	AddTask(dto *dto.AddTask) (*entity.Task, error)
	GetTask(taskId *uuid.UUID) (*entity.Task, error)
	GetTasks(dto *dto.GetTasks) ([]*entity.Task, error)
	UpdateTask(dto *dto.UpdateTask) (*entity.Task, error)
	RemoveTask(taskId *uuid.UUID) error

	AddUser(dto *dto.AddUser) (*entity.User, error)
	GetUser(userId *uuid.UUID) (*entity.User, error)
	GetUsers(dto *dto.GetUsers) ([]*entity.User, error)
	UpdateUser(dto *dto.UpdateUser) (*entity.User, error)
	RemoveUser(userId *uuid.UUID) error

	AddUserTask(dto *dto.AddUserTask) (*entity.UserTask, error)
	GetUserTasks(dto *dto.GetUserTasks) ([]*entity.UserTask, error)
	RemoveUserTask(userTaskID *uuid.UUID) error

	GetAccessPermissions() []*entity.AccessPermission
	Login(userName string) (string, error)
}
