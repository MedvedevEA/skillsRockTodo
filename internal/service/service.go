package service

import (
	"log/slog"
	"skillsRockTodo/internal/repository"
)

type Service struct {
	store repository.Repository
	lg    *slog.Logger
}

func New(store repository.Repository, lg *slog.Logger) *Service {
	return &Service{
		store,
		lg,
	}
}
