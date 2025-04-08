package storemap

import (
	"errors"
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository"
	"time"
)

type StoreMap struct {
	serial int
	data   []*entity.Task
}

func New() *StoreMap {
	return &StoreMap{
		serial: 0,
		data:   []*entity.Task{},
	}
}

func (s *StoreMap) CreateTask(dto *repository.DtoCreateTaskReq) error {

	s.serial++
	task := entity.Task{
		Id:          s.serial,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	s.data = append(s.data, &task)
	return nil
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
	return nil, errors.New("task not found")
}
func (s *StoreMap) UpdateTask(dto *repository.DtoUpdateTaskReq) error {
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
	return errors.New("task not found")
}
func (s *StoreMap) DeleteTask(Id int) error {
	for i := range s.data {
		if s.data[i].Id == Id {
			s.data[i] = s.data[len(s.data)-1]
			s.data = s.data[:len(s.data)-1]
			return nil
		}
	}
	return errors.New("task not found")
}
