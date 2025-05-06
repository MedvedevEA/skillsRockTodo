package service

import (
	"skillsRockTodo/internal/entity"
	repoStoreDto "skillsRockTodo/internal/repository/repostore/dto"
)

func (s *Service) GetUsers(dto *repoStoreDto.GetUsers) ([]*entity.User, error) {
	return s.store.GetUsers(dto)
}
