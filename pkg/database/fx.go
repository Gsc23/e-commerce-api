package database

import (
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"go.uber.org/fx"
)

type DBParams struct {
	fx.In

	Config config.Config
}

type DBResult struct {
	fx.Out

	Database DB
}

func NewDatabase(lc fx.Lifecycle, p DBParams) (DBResult, error) {
	db, err := newPostgres(p.Config)
	if err != nil {
		return DBResult{}, err
	}

	lc.Append(fx.Hook{
		OnStart: db.Start,
		OnStop:  db.Stop,
	})

	return DBResult{Database: db}, nil
}

func DBModule() fx.Option {
	return fx.Module("DB",
		fx.Provide(NewDatabase),
		fx.Invoke(ResolveDB),
	)
}

func ResolveDB(_ DB) {}
