package http

import (
	"context"
	"fmt"

	"github.com/Gsc23/e-commerce-api/pkg/config"
	"github.com/Gsc23/e-commerce-api/pkg/logger"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In

	Config config.Config
	Logger logger.Logger
}

type ServerResult struct {
	fx.Out
	Server Server
}

func NewServer(lc fx.Lifecycle, p ServerParams) ServerResult {
	log := p.Logger

	srv := newServer(p.Config)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info(ctx,
				"starting http server",
				"addr", fmt.Sprintf(":%d", p.Config.ServerPort()),
				"env", p.Config.Env(),
			)
			return srv.Run(ctx)
		},
		OnStop: func(ctx context.Context) error {
			log.Info(ctx, "stopping http server")
			return srv.Stop(ctx)
		},
	})

	return ServerResult{Server: srv}
}

func HTTPModule() fx.Option {
	return fx.Module("http",
		fx.Provide(NewServer),
		fx.Invoke(ResolveHTTPServer),
	)
}

func ResolveHTTPServer(_ Server) {}
