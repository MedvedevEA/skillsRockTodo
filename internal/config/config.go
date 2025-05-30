package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env   string `yaml:"env" env:"TODO_ENV" env-default:"local"`
	Api   Api    `yaml:"api"`
	Store Store  `yaml:"store"`
}

type Api struct {
	PublicKeyPath string        `yaml:"publicKeyPath" env:"TODO_API_PUBLIC_KEY_PATH" env-required:"true"`
	Addr          string        `yaml:"addr" env:"TODO_API_ADDR" env-required:"true"`
	WriteTimeout  time.Duration `yaml:"writeTimeout" env:"TODO_API_WRITE_TIMEOUT" env-required:"true"`
	Name          string        `yaml:"name" env:"TODO_API_NAME" env-required:"true"`
}

type Store struct {
	Host                string        `yaml:"host" env:"TODO_STORE_HOST" env-required:"true"`
	Port                int           `yaml:"port" env:"TODO_STORE_PORT" env-required:"true"`
	Name                string        `yaml:"name" env:"TODO_STORE_NAME" env-required:"true"`
	User                string        `yaml:"user" env:"TODO_STORE_USER" env-required:"true"`
	Password            string        `yaml:"password" env:"TODO_STORE_PASSWORD" env-required:"true"`
	SSLMode             string        `yaml:"sslMode" env:"TODO_STORE_SSL_MODE" env-default:"disable"`
	PoolMaxConns        int           `yaml:"poolMaxConns" env:"TODO_STORE_POOL_MAX_CONNS" env-default:"5"`
	PoolMaxConnLifetime time.Duration `yaml:"poolMaxConnLifeTime" env:"TODO_STORE_POOL_MAX_CONN_LIFETIME" env-default:"180s"`
	PoolMaxConnIdleTime time.Duration `yaml:"poolMaxConnIidleTime" env:"TODO_STORE_POOL_MAX_CONN_IDLE_TIME" env-default:"100s"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"

	configPath := ""
	cfg := new(Config)

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
	configPath = "./../../config/local.yml" //for debug
	if configPath != "" {
		log.Printf("%s: the value of the 'config' flag: %s\n", op, configPath)
		if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
			log.Fatal(errors.Wrap(err, op))
		}
		return cfg
	}
	log.Printf("%s: the 'config' flag is not set\n", op)

	configPath = os.Getenv("TODO_CONFIG_PATH")
	if configPath != "" {
		log.Printf("%s: the value of the environment variable: %s\n", op, configPath)
		if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
			log.Fatal(errors.Wrap(err, op))
		}
		return cfg
	}
	log.Printf("%s: environment variable 'TODO_CONFIG_PATH' is not set\n", op)

	log.Printf("%s: the parameter file is not defined. Loading the application configuration from the environment variables\n", op)
	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatal(errors.Wrap(err, op))
	}
	log.Printf("%s: configuration file %+v", op, cfg)
	return cfg
}
