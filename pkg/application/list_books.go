package application

import (
	"context"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
)

func (b BookServiceImpl) ListBooks(ctx context.Context, page, limit int32) (books []books.Books, err error) {
	return b.bookRepo.ListBooks(ctx, page, limit)
}
