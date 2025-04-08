package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

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

	store := storemap.New()

	service := service.New(store, logger)
	/*
		// Инициализация API
		app := api.NewRouters(&api.Routers{Service: serviceInstance}, cfg.Rest.Token)

		// Запуск HTTP-сервера в отдельной горутине
		go func() {
			logger.Infof("Starting server on %s", cfg.Rest.ListenAddress)
			if err := app.Listen(cfg.Rest.ListenAddress); err != nil {
				log.Fatal(errors.Wrap(err, "failed to start server"))
			}
		}()

		// Ожидание системных сигналов для корректного завершения работы
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		<-signalChan

		logger.Info("Shutting down gracefully...")
	*/
}
