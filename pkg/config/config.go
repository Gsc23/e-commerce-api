package config

type EnvFile string

type Config struct {
	Server struct {
		Host string `env:"HOST"`
		Port int    `env:"PORT" default:"8080"`
		Env  string `env:"ENV" default:"development"`
	} `env:"SERVER"`
}
