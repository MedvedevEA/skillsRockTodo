package service

import (
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"
	crypto "skillsRockTodo/pkg/crypto"

	"github.com/google/uuid"
)

func (s *Service) AddUser(dto *dto.AddUser) (*entity.User, error) {
	hash := crypto.GetHash(dto.Password)
	dto.Password = hash
	return s.store.AddUser(dto)
}
func (s *Service) GetUser(userId *uuid.UUID) (*entity.User, error) {
	return s.store.GetUser(userId)
}
func (s *Service) GetUsers(dto *dto.GetUsers) ([]*entity.User, error) {
	return s.store.GetUsers(dto)
}
func (s *Service) UpdateUser(dto *dto.UpdateUser) (*entity.User, error) {
	if dto.Password != nil {
		hash := crypto.GetHash(*dto.Password)
		dto.Password = &hash
	}
	return s.store.UpdateUser(dto)
}
func (s *Service) RemoveUser(userId *uuid.UUID) error {
	return s.store.RemoveUser(userId)
}
