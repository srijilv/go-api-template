package books

import "context"

type BookStorageRepository interface {
	ListBooks(ctx context.Context) (err error)
}
