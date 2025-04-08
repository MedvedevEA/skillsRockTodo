package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"skillsRockTodo/internal/config"
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

}
