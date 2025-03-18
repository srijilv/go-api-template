package db

import (
	"context"
)

type DefaultPgSqlConnectionPoolProvider struct {
	GetPoolProvider func(context.Context) (PgSqlConnectionPoolProvider, error)
}

// func (p DefaultPgSqlConnectionPoolProvider) GetPgSqlConnectionPool(ctx context.Context, dbName string) (pool PgxIface, err error) {
// 	// if p.GetPoolProvider == nil {
// 	// 	p.GetPoolProvider = GetPgSqlPoolConnectionProvider
// 	// }
// 	provider, err := p.GetPoolProvider(ctx)
// 	if err != nil {
// 		return
// 	}

// 	pool, err = GetPoolTxWrapper(ctx)
// 	if err == nil {
// 		logrus.Info("there is transaction in context, wrapping it as pool and returning")
// 		return
// 	}

// 	pool, err = provider.GetPgSqlConnectionPool(ctx, dbName)
// 	return
// }
// func GetPoolTxWrapper(ctx context.Context) (pool PgxIface, err error) {
// 	var (
// 		tx pgx.Tx
// 		// ok bool
// 	)
// 	// if tx, ok = ctx.Value(ctxtypes.CtxTx).(pgx.Tx); !ok {
// 	// 	err = fmt.Errorf("there is no tx in context")
// 	// 	return
// 	// }

// 	pool = PoolTxWrapper{tx: tx}
// 	return
// }
