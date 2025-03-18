package listbooks

import (
	"fmt"
	"net/http"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/common"
)

func (l ListBooksResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (l *ListBooksResponse) Unmarshal(records []books.Books, info common.Information) {
	l.Info = info
	l.Payload = unmarshalPayload(records)
}

func unmarshalPayload(records []books.Books) (payload []ListBooksResponsePayload) {
	for _, b := range records {
		payload = append(payload, ListBooksResponsePayload{
			Id:            b.Id(),
			Title:         b.Title(),
			Author:        b.Author(),
			Genre:         b.Genre(),
			Isbn:          b.ISBN(),
			Price:         fmt.Sprintf("%.2f", b.Price()),
			PublishedYear: b.PublishedYear(),
			Status:        b.Status(),
		})
	}
	return
}
