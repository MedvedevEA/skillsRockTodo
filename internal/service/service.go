package service

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository"

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

func (s *Service) CreateTask(dto *repository.DtoCreateTaskReq) error {
	return s.store.CreateTask(dto)
}
func (s *Service) GetTasks() ([]*entity.Task, error) {
	return s.store.GetTasks()
}
func (s *Service) GetTask(Id int) (*entity.Task, error) {
	return s.store.GetTask(Id)
}
func (s *Service) UpdateTask(dto *repository.DtoUpdateTaskReq) error {
	return s.store.UpdateTask(dto)
}
func (s *Service) DeleteTask(Id int) error {
	return s.store.DeleteTask(Id)
}
