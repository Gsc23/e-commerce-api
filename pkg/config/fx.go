package config

import "go.uber.org/fx"

type ConfigResult struct {
	fx.Out
	Config *Config
}

func NewConfig() ConfigResult {
	cfg, err := loadEnvFile(".env")
	if err != nil {
		panic(err)
	}

	return ConfigResult{Config: cfg}
}

func ConfigModule() fx.Option {
	return fx.Module("config",
		fx.Provide(NewConfig),
		fx.Invoke(StartConfig),
	)
}

func StartConfig(_ *Config) {}
