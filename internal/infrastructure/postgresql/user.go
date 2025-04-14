package postgresql

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"

	"github.com/google/uuid"
)

func (p *PostgreSql) AddUser(dto *dto.AddUser) (*entity.User, error) {
	return nil, nil
}
func (p *PostgreSql) GetUser(userId *uuid.UUID) (*entity.User, error) {
	return nil, nil
}
func (p *PostgreSql) GetUsers(dto *dto.GetUsers) ([]*entity.User, error) {
	return nil, nil
}
func (p *PostgreSql) UpdateUser(dto *dto.UpdateUser) (*entity.User, error) {
	return nil, nil
}
func (p *PostgreSql) RemoveUser(userId *uuid.UUID) error {
	return nil
}
