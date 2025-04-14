package service

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository"
	"skillsRockTodo/internal/repository/dto"

	"go.uber.org/zap"
)

type Service struct {
	store repository.Repository
	log   *zap.SugaredLogger
}

func New(store repository.Repository, log *zap.SugaredLogger) *Service {
	return &Service{
		store,
		log,
	}
}

func (s *Service) AddTask(dto *dto.AddTask) (*entity.Task, error) {
	return s.store.AddTask(dto)
}
func (s *Service) GetTasks() ([]*entity.Task, error) {
	return s.store.GetTasks()
}
func (s *Service) GetTask(Id int) (*entity.Task, error) {
	return s.store.GetTask(Id)
}
func (s *Service) UpdateTask(dto *dto.UpdateTask) error {
	return s.store.UpdateTask(dto)
}
func (s *Service) RemoveTask(Id int) error {
	return s.store.RemoveTask(Id)
}
