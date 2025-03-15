package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/gommon/log"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type PgSqlConnectionPoolProvider interface {
	GetPgSqlConnectionPool(ctx context.Context, dbName string) (PgxIface, error)
}

type PoolTxWrapper struct {
	tx pgx.Tx
}

func (p PoolTxWrapper) Begin(ctx context.Context) (pgx.Tx, error) {
	log.Info("Beginning nested transaction")
	tx, err := p.tx.Begin(ctx)
	return tx, err
}

func (p PoolTxWrapper) Close() {}
