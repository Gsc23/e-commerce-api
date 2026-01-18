package config

import "go.uber.org/fx"

type ConfigResult struct {
	fx.Out

	Config Config
}

func ConfigModule() fx.Option {
	return fx.Module("config",
		fx.Provide(NewConfig),
	)
}
