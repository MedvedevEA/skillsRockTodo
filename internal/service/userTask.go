package service

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"

	"github.com/google/uuid"
)

func (s *Service) AddUserTask(dto *dto.AddUserTask) (*entity.UserTask, error) {
	return s.store.AddUserTask(dto)
}
func (s *Service) GetUserTask(userTaskId *uuid.UUID) (*entity.UserTask, error) {
	return s.store.GetUserTask(userTaskId)
}
func (s *Service) GetUserTasks(dto *dto.GetUserTasks) ([]*entity.UserTask, error) {
	return s.store.GetUserTasks(dto)
}
func (s *Service) RemoveUserTask(userTaskId *uuid.UUID) error {
	return s.store.RemoveUserTask(userTaskId)
}
