package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"skillsRockTodo/internal/apiserver"
	"skillsRockTodo/internal/config"
	"skillsRockTodo/internal/infrastructure/storemap"
	customLogger "skillsRockTodo/internal/logger"
	"skillsRockTodo/internal/service"
)

func main() {

	flagConfig := flag.String("config", "./../../config/todo.env", "Configuration file")
	flag.Parse()
	if *flagConfig != "" {
		err := godotenv.Load(*flagConfig)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to load .env file"))
		}
	}

	var cfg config.TodoConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "failed to load configuration"))
	}

	logger, err := customLogger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error initializing logger"))
	}

	store := storemap.New(logger)

	service := service.New(store, logger)

	logger.Infof("API Server '%s' is started in addr:[%s]", cfg.Rest.ServerName, cfg.Rest.ListenAddress)
	apiServer := apiserver.New(service, logger, cfg.Rest)
	if err := apiServer.Run(); err != nil {
		logger.Fatalf("API Server '%s' error: %s", cfg.Rest.ServerName, err)
	}
	logger.Infof("API Server '%s' is stoped", cfg.Rest.ServerName)

}
