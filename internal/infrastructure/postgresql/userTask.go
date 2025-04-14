package postgresql

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"

	"github.com/google/uuid"
)

func (p *PostgreSql) AddUserTask(dto *dto.AddUserTask) (*entity.UserTask, error) {
	return nil, nil
}
func (p *PostgreSql) GetUserTask(userTaskId *uuid.UUID) (*entity.UserTask, error) {
	return nil, nil
}
func (p *PostgreSql) GetUserTasks(dto *dto.GetUserTasks) ([]*entity.UserTask, error) {
	return nil, nil
}
func (p *PostgreSql) RemoveUserTask(userTaskId *uuid.UUID) error {
	return nil
}
