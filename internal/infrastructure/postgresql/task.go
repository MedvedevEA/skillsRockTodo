package postgresql

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"

	"github.com/google/uuid"
)

func (p *PostgreSql) AddTask(dto *dto.AddTask) (*entity.Task, error) {
	return nil, nil
}
func (p *PostgreSql) GetTask(taskId *uuid.UUID) (*entity.Task, error) {
	return nil, nil
}
func (p *PostgreSql) GetTasks(dto *dto.GetTasks) ([]*entity.Task, error) {
	return nil, nil
}
func (p *PostgreSql) UpdateTask(dto *dto.UpdateTask) (*entity.Task, error) {
	return nil, nil
}
func (p *PostgreSql) RemoveTask(taskId *uuid.UUID) error {
	return nil
}
