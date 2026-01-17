package database

import (
	"fmt"

	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"go.uber.org/fx"
)

type DBParams struct {
	fx.In
	Config *config.Config
}

type DBResult struct {
	fx.Out
	Database DB
}

func NewDatabase(lc fx.Lifecycle, p DBParams) (*DBResult, error) {
	db, err := newPostgres(p.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: db.Start,
		OnStop:  db.Stop,
	})

	return &DBResult{Database: db}, nil
}

func DBModule() fx.Option {
	return fx.Module("DB",
		fx.Provide(NewDatabase),
		fx.Invoke(StartDB),
	)
}

func StartDB(_ DB) {}
