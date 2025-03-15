package pgsql

import (
	"github.com/srijilv/go-api-template.git/db"
	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
)

type PgsqlBookRepository struct {
	poolProvider db.PgSqlConnectionPoolProvider
}

func NewPgsqlBookRepository(poolProvider db.PgSqlConnectionPoolProvider) books.BookStorageRepository {
	if poolProvider == nil {
		panic("pool provider is nil")
	}
	return PgsqlBookRepository{
		poolProvider: poolProvider,
	}
}
