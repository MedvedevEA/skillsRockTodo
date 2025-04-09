package storemap

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"
	"skillsRockTodo/pkg/servererrors"
	"time"

	"go.uber.org/zap"
)

type StoreMap struct {
	log    *zap.SugaredLogger
	serial int
	data   []*entity.Task
}

func New(log *zap.SugaredLogger) *StoreMap {
	return &StoreMap{
		log:    log,
		serial: 0,
		data:   []*entity.Task{},
	}
}

func (s *StoreMap) AddTask(dto *dto.AddTask) (*entity.Task, error) {

	s.serial++
	task := entity.Task{
		Id:          s.serial,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      "new",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	task2 := task
	s.data = append(s.data, &task)
	return &task2, nil
}
func (s *StoreMap) GetTasks() ([]*entity.Task, error) {
	tasks := []*entity.Task{}
	for i := range s.data {
		task := *s.data[i]
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
func (s *StoreMap) GetTask(Id int) (*entity.Task, error) {

	for i := range s.data {
		if s.data[i].Id == Id {
			task := *s.data[i]
			return &task, nil
		}
	}
	return nil, servererrors.ErrorRecordNotFound
}
func (s *StoreMap) UpdateTask(dto *dto.UpdateTask) error {
	for i := range s.data {
		if s.data[i].Id == dto.Id {
			if dto.Title != nil {
				s.data[i].Title = *dto.Title
			}
			if dto.Description != nil {
				s.data[i].Description = *dto.Description
			}
			if dto.Status != nil {
				s.data[i].Status = *dto.Status
			}
			s.data[i].UpdateAt = time.Now()
			return nil
		}
	}
	return servererrors.ErrorRecordNotFound
}
func (s *StoreMap) RemoveTask(Id int) error {
	for i := range s.data {
		if s.data[i].Id == Id {
			s.data[i] = s.data[len(s.data)-1]
			s.data = s.data[:len(s.data)-1]
			return nil
		}
	}
	return servererrors.ErrorRecordNotFound
}
