package books

import (
	"context"
)

type BookStorageRepository interface {
	ListBooks(ctx context.Context, page, limit int32) (books []Books, err error)
}
