package logger

import (
	"github.com/Gsc23/e-commerce-api/pkg/config"
	"go.uber.org/fx"
)

type LoggerParams struct {
	fx.In

	Config config.Config
}

type LoggerResult struct {
	fx.Out

	GlobalLogger Logger
	Factory      *LoggerFactory
}

func LoggerModule() fx.Option {
	return fx.Module("Logger",
		fx.Provide(NewGlobalLogger),
	)
}

func NewGlobalLogger(p LoggerParams) (LoggerResult, error) {
	loggerConfig, err := newLoggerConfig(p.Config)
	if err != nil {
		return LoggerResult{}, err
	}

	factory := NewLoggerFactory(loggerConfig)

	return LoggerResult{
		Factory:      factory,
		GlobalLogger: factory.NewLoggerNamed("app"),
	}, nil
}
