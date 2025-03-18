package data

import (
	"time"
)

type LoadableBooks struct {
	Id            int32
	Title         string
	Author        string
	Genre         string
	PublishedYear int16
	Isbn          string
	Price         float32
	Status        string
	CreatedBy     int64
	CreatedAt     time.Time
	UpdatedBy     int64
	UpdatedAt     time.Time
}
