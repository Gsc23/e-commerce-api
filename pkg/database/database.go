package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	Pool() *pgxpool.Pool
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type PostgresImpl struct {
	pool *pgxpool.Pool
}

func newPostgres(config *config.Config) (DB, error) {
	connUrl := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.DB.User, config.DB.Pass),
		Host:   config.DB.Host,
		Path:   config.DB.Database,
	}

	if config.DB.Schema.Valid {
		q := connUrl.Query()
		q.Add("options", fmt.Sprintf("-csearch_path=%s", config.DB.Schema.String))
		connUrl.RawQuery = q.Encode()
	}

	poolConfig, err := pgxpool.ParseConfig(connUrl.String())
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &PostgresImpl{pool: conn}, nil
}

func (p *PostgresImpl) Start(ctx context.Context) error {
	if err := p.pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}

func (p *PostgresImpl) Stop(ctx context.Context) error {
	p.pool.Close()
	return nil
}

func (p *PostgresImpl) Pool() *pgxpool.Pool {
	return p.pool
}
