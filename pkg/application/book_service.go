package application

import (
	"context"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
)

type BookService interface {
	ListBooks(ctx context.Context, page, limit int32) (books []books.Books, err error)
}

type BookServiceImpl struct {
	bookRepo books.BookStorageRepository
}

func NewBooksService(bookRepo books.BookStorageRepository) BookService {
	if bookRepo == nil {
		panic("book repo is nil")
	}
	return BookServiceImpl{
		bookRepo: bookRepo,
	}
}
