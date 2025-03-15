package listbooks

import (
	"net/http"

	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/common"
)

func (l ListBooksResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (l *ListBooksResponse) Unmarshal(books []books.Books, info common.Information) {
	l.Info = info
	l.Payload = unmarshalPayload(books)
}

func unmarshalPayload(books []books.Books) (payload []ListBooksResponsePayload) {
	for _, b := range books {
		payload = append(payload, ListBooksResponsePayload{
			Id:   b.Id(),
			Name: b.Name(),
		})
	}
	return
}
