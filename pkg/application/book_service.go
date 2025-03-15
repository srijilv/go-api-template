package application

import (
	"context"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
)

type BookService interface {
	ListBooks(ctx context.Context) (books []books.Books, err error)
}
