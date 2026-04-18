package database

import (
	"context"

	"github.com/Gsc23/e-commerce-api/pkg/config"
	"github.com/Gsc23/e-commerce-api/pkg/logger"
	"go.uber.org/fx"
)

type DBParams struct {
	fx.In

	Config config.Config
	Logger logger.Logger
}

type DBResult struct {
	fx.Out

	Database DB
}

func NewDatabase(lc fx.Lifecycle, p DBParams) (DBResult, error) {
	log := p.Logger

	db, err := newPostgres(p.Config)
	if err != nil {
		log.Error(context.Background(), "failed to create database config", "error", err)
		return DBResult{}, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info(ctx, "starting database")
			if err := db.Start(ctx); err != nil {
				log.Error(ctx, "database start failed", "error", err)
				return err
			}
			log.Info(ctx, "database started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info(ctx, "stopping database")
			return db.Stop(ctx)
		},
	})

	return DBResult{Database: db}, nil
}

func DBModule() fx.Option {
	return fx.Module("DB",
		fx.Provide(NewDatabase),
	)
}
