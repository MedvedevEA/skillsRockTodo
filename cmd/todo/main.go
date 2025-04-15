package main

import (
	"context"

	"skillsRockTodo/internal/apiserver"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/infrastructure/postgresql"
	"skillsRockTodo/internal/logger"
	"skillsRockTodo/internal/service"
)

func main() {
	cfg := config.MustLoad()

	lg := logger.MustNew(cfg.Env)

	store := postgresql.MustNew(context.Background(), lg, &cfg.PostgreSQL)

	service := service.New(store, lg)

	apiServer := apiserver.New(service, lg, &cfg.Api)
	apiServer.MustRun()

}
