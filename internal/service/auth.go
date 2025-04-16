package service

import (
	srvDto "skillsRockTodo/internal/service/dto"
	"skillsRockTodo/pkg/crypto"
	"skillsRockTodo/pkg/servererrors"
)

func (s *Service) Login(dto *srvDto.Login) (string, error) {
	passwordHash, err := s.store.Login(dto.Name)
	if err != nil {
		return "", err
	}
	if crypto.CheckHash(dto.Password, passwordHash) {
		token := crypto.GenerateSecureToken()
		return token, nil
	}

	return "", servererrors.InvalidUsernameOrPassword
}
