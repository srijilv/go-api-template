package books

import "time"

func (b Books) Id() int32 {
	return b.id
}

func (b Books) Title() string {
	return b.title
}

func (b Books) Author() string {
	return b.author
}

func (b Books) Genre() string {
	return b.genre
}

func (b Books) PublishedYear() int16 {
	return b.publishedYear
}

func (b Books) ISBN() string {
	return b.isbn
}

func (b Books) Price() float32 {
	return b.price
}

func (b Books) Status() string {
	return b.status
}

func (b Books) CreatedBy() int64 {
	return b.createdBy
}

func (b Books) CreatedAt() time.Time {
	return b.createdAt
}

func (b Books) UpdatedBy() int64 {
	return b.updatedBy
}

func (b Books) UpdatedAt() time.Time {
	return b.updatedAt
}
