package main

import (
	"context"
	"fmt"
	"log"

	"skillsRockTodo/internal/apiserver"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/infrastructure/postgresql"
	"skillsRockTodo/internal/logger"
	"skillsRockTodo/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	lg, err := logger.New(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	store, err := postgresql.New(context.Background(), &cfg.PostgreSQL, lg)
	if err != nil {
		log.Fatal(err)
	}

	service := service.New(store, lg)

	lg.Info(fmt.Sprintf("API Server '%s' is started in addr:[%s]", cfg.Api.Name, cfg.Api.Addr))
	apiServer := apiserver.New(service, lg, &cfg.Api)
	if err := apiServer.Run(); err != nil {
		lg.Error(fmt.Sprintf("API Server '%s' error: %s", cfg.Api.Name, err))
		return
	}
	lg.Info(fmt.Sprintf("API Server '%s' is stoped", cfg.Api.Name))
	return

}
