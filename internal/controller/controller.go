package controller

import (
	"log/slog"
	"skillsRockTodo/internal/service"
)

type Controller struct {
	service *service.Service
	lg      *slog.Logger
}

func New(service *service.Service, lg *slog.Logger) *Controller {
	return &Controller{
		service,
		lg,
	}
}
