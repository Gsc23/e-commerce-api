package http

import (
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In
	Config config.Config
}

type ServerResult struct {
	fx.Out
	Server Server
}

func NewServer(lc fx.Lifecycle, p ServerParams) ServerResult {
	srv := newServer(p.Config)
	lc.Append(fx.Hook{OnStart: srv.Run, OnStop: srv.Stop})

	return ServerResult{Server: srv}
}

func HTTPModule() fx.Option {
	return fx.Module("http",
		fx.Provide(NewServer),
		fx.Invoke(ResolveHTTPServer),
	)
}

func ResolveHTTPServer(_ Server) {}
