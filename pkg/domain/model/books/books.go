package books

import (
	"time"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books/data"
)

type Books struct {
	id            int32
	title         string
	author        string
	genre         string
	publishedYear int16
	isbn          string
	price         float32
	status        string
	createdBy     int64
	createdAt     time.Time
	updatedBy     int64
	updatedAt     time.Time
}

func Unmarshal(lb data.LoadableBooks) (b Books, err error) {
	b.id = lb.Id
	b.title = lb.Title
	b.author = lb.Author
	b.genre = lb.Genre
	b.publishedYear = lb.PublishedYear
	b.isbn = lb.Isbn
	b.price = lb.Price
	b.status = lb.Status
	b.createdBy = lb.CreatedBy
	b.createdAt = lb.CreatedAt
	b.updatedBy = lb.UpdatedBy
	b.updatedAt = lb.UpdatedAt
	return
}

func UnmarshalSlice(loadables []data.LoadableBooks) (books []Books, err error) {
	for _, lb := range loadables {
		b, err := Unmarshal(lb)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return
}
