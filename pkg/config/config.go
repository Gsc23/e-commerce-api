package config

import (
	"errors"
	"fmt"

	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type Config struct {
	Server struct {
		Host string `env:"HOST"`
		Port int    `env:"PORT" default:"8080"`
		Env  string `env:"ENV" default:"dev"`
	} `env:"SERVER"`
	DB struct {
		Host     string `env:"HOST"`
		Port     string `env:"PORT"`
		Database string `env:"DATABASE"`
		User     string `env:"USER"`
		Pass     string `env:"PASSWORD"`
	} `env:"DB"`
}

func loadEnvFile(envFile string) (*Config, error) {
	if err := godotenv.Load(envFile); err != nil {
		return nil, err
	}

	opts := &env.Options{NameSep: "_"}

	var config Config

	if err := env.Load(&config, opts); err != nil {
		return nil, fmt.Errorf("could not parse config: %w", err)
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

func (c *Config) validate() error {
	if c.Server.Port == 0 {
		return errors.New("server port is required")
	}
	if c.Server.Env == "" {
		return errors.New("server environment is required")
	}

	return nil
}
