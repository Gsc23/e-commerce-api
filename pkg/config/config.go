package config

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/guregu/null/v6"
	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type Config interface {
	ServerPort() int
	ServerHost() string
	Env() string
	DBConnString() string
	LoggerLevel() string
	LoggerColors() bool
	LoggerTrace() bool
}

type config struct {
	Server struct {
		Host string `env:"HOST"`
		Port int    `env:"PORT" default:"8080"`
		Env  string `env:"ENV" default:"dev"`
	} `env:"SERVER"`
	DB struct {
		Host     string      `env:"HOST"`
		Port     string      `env:"PORT"`
		Database string      `env:"DATABASE"`
		User     string      `env:"USER"`
		Pass     string      `env:"PASSWORD"`
		Schema   null.String `env:"SCHEMA" default:"app"`
	} `env:"DB"`
	Logger struct {
		Level      string `env:"LEVEL"`
		IsColorful bool   `env:"COLORFULL"`
		Trace      bool   `env:"TRACE"`
	} `env:"LOGGER"`
}

func NewConfig() (ConfigResult, error) {
	config := loadEnvFile(".env")
	if config == nil {
		return ConfigResult{}, errors.New("failed to load dotenv")
	}

	return ConfigResult{
		Config: config,
	}, nil
}

func loadEnvFile(envFile string) *config {
	if err := godotenv.Load(envFile); err != nil {
		return nil
	}

	opts := &env.Options{NameSep: "_"}

	var config config

	if err := env.Load(&config, opts); err != nil {
		return nil
	}

	if err := config.validate(); err != nil {
		return nil
	}

	return &config
}

func (c *config) validate() error {
	if c.Server.Port == 0 {
		return errors.New("server port is required")
	}
	if c.Server.Env == "" {
		return errors.New("server environment is required")
	}

	return nil
}

func (c *config) ServerPort() int {
	return c.Server.Port
}

func (c *config) ServerHost() string {
	return c.Server.Host
}

func (c *config) Env() string {
	return c.Server.Env
}

func (c *config) DBConnString() string {
	connUrl := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.DB.User, c.DB.Pass),
		Host:   c.DB.Host,
		Path:   c.DB.Database,
	}

	if c.DB.Schema.Valid {
		q := connUrl.Query()
		q.Add("options", fmt.Sprintf("-csearch_path=%s", c.DB.Schema.String))
		connUrl.RawQuery = q.Encode()
	}

	return connUrl.String()
}

func (c *config) LoggerLevel() string {
	return c.Logger.Level
}

func (c *config) LoggerColors() bool {
	return c.Logger.IsColorful
}

func (c *config) LoggerTrace() bool {
	return c.Logger.Trace
}
