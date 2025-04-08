package storemap

import "skillsRockTodo/internal/entity"

type StoreMap struct {
	data []*entity.Task
}

func New() *StoreMap {
	return &StoreMap{
		data: []*entity.Task{},
	}
}

func (s *StoreMap) CreateTask(task *entity.Task) error {
	return nil
}
func (s *StoreMap) GetTasks() ([]*entity.Task, error) {
	return nil, nil
}
func (s *StoreMap) GetTask(taskId int) (*entity.Task, error) {
	return nil, nil
}
func (s *StoreMap) UpdateTask(task *entity.Task) error {
	return nil
}
func (s *StoreMap) DeleteTask(taskId int) error {
	return nil
}
