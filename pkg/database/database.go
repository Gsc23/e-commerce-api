package database

import (
	"context"

	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	GetPool() *pgxpool.Pool
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type PostgresImpl struct {
	DBConfig *pgxpool.Config
	Pool     *pgxpool.Pool
}

func newPostgres(cfg config.Config) (*PostgresImpl, error) {
	DBConfig, err := pgxpool.ParseConfig(cfg.DBConnString())
	if err != nil {
		return nil, err
	}

	return &PostgresImpl{
		DBConfig: DBConfig,
	}, nil
}

func (p *PostgresImpl) Start(ctx context.Context) error {
	pool, err := pgxpool.NewWithConfig(ctx, p.DBConfig)
	if err != nil {
		return err
	}

	p.Pool = pool
	return p.Pool.Ping(ctx)
}

func (p *PostgresImpl) Stop(ctx context.Context) error {
	p.Pool.Close()
	return nil
}

func (p *PostgresImpl) GetPool() *pgxpool.Pool {
	if p.Pool == nil {
		panic("database not started")
	}

	return p.Pool
}
