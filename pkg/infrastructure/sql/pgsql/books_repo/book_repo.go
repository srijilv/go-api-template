package booksrepo

import (
	"github.com/srijilv/go-api-template.git/db"
	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
)

const (
	layer  = "infrastructure"
	dbType = "postgres"

	booksTable = "books"
)

type PgsqlBookRepository struct {
	poolProvider db.PgSqlConnectionPoolProvider
	dbName       string
}

func NewPgsqlBookRepository(poolProvider db.PgSqlConnectionPoolProvider, dbName string) books.BookStorageRepository {
	if poolProvider == nil {
		panic("pool provider is nil")
	}
	if dbName == "" {
		panic("db name is empty")
	}
	return PgsqlBookRepository{
		poolProvider: poolProvider,
		dbName:       dbName,
	}
}
