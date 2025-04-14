package service

import (
	"log/slog"
	"skillsRockTodo/internal/repository"
)

type Service struct {
	store repository.Repository
	log   *slog.Logger
}

func New(store repository.Repository, log *slog.Logger) *Service {
	return &Service{
		store,
		log,
	}
}
