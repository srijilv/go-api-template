package listbooks

import (
	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/grpc/common"
)

func Unmarshal(info *common.Information, records []books.Books) (response ListBooksResponse) {
	response.Info = info
	response.Payload = unmarshalPayload(records)
	return
}

func unmarshalPayload(records []books.Books) (appointmnetResp []*ListBooksPayload) {
	for _, b := range records {
		appointmnetResp = append(appointmnetResp, &ListBooksPayload{
			Id:            b.Id(),
			Title:         b.Title(),
			Author:        b.Author(),
			Genre:         b.Genre(),
			Isbn:          b.ISBN(),
			Price:         b.Price(),
			PublishedYear: int32(b.PublishedYear()),
			Status:        b.Status(),
		})

	}
	return
}
