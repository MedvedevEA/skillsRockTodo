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
	Env        string     `yaml:"env" env:"TODO_ENV" env-default:"local"`
	Api        Api        `yaml:"api"`
	PostgreSQL PostgreSQL `yaml:"postgreSql"`
}

type Api struct {
	Addr         string        `yaml:"addr" env:"TODO_API_ADDR" env-required:"true"`
	WriteTimeout time.Duration `yaml:"writeTimeout" env:"TODO_API_WRITE_TIMEOUT" env-required:"true"`
	Name         string        `yaml:"name" env:"TODO_API_NAME" env-required:"true"`
}

type PostgreSQL struct {
	Host                string        `yaml:"host" env:"TODO_PG_HOST" env-required:"true"`
	Port                int           `yaml:"port" env:"TODO_PG_PORT" env-required:"true"`
	Name                string        `yaml:"name" env:"TODO_PG_NAME" env-required:"true"`
	User                string        `yaml:"user" env:"TODO_PG_USER" env-required:"true"`
	Password            string        `yaml:"password" env:"TODO_PG_PASSWORD" env-required:"true"`
	SSLMode             string        `yaml:"sslMode" env:"TODO_PG_SSL_MODE" env-default:"disable"`
	PoolMaxConns        int           `yaml:"poolMaxConns" env:"TODO_PG_POOL_MAX_CONNS" env-default:"5"`
	PoolMaxConnLifetime time.Duration `yaml:"poolMaxConnLifeTime" env:"TODO_PG_POOL_MAX_CONN_LIFETIME" env-default:"180s"`
	PoolMaxConnIdleTime time.Duration `yaml:"poolMaxConnIidleTime" env:"TODO_PG_POOL_MAX_CONN_IDLE_TIME" env-default:"100s"`
}

func Load() (*Config, error) {
	const op = "config.Load"

	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		log.Println("the 'config' flag is not set")
		configPath = os.Getenv("TODO_CONFIG_PATH")
	}
	var cfg Config

	if configPath == "" {
		log.Println("environment variable 'TODO_CONFIG_PATH' is not set")
		log.Println("creating a configuration file based on environment variables")
		err := cleanenv.ReadEnv(&cfg)
		return &cfg, err
	}
	log.Printf("creating a configuration file based on %s\n", configPath)
	err := cleanenv.ReadConfig(configPath, &cfg)
	return &cfg, errors.Wrap(err, op)
}
