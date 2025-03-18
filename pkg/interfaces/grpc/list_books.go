package grpc

import (
	context "context"

	"github.com/sirupsen/logrus"
	"github.com/srijilv/go-api-template.git/pkg/components"
	listbooks "github.com/srijilv/go-api-template.git/pkg/interfaces/grpc/list_books"
)

func (s Server) Listbooks(ctx context.Context, in *listbooks.ListBooksRequest) (out *listbooks.ListBooksResponse, err error) {
	logrus.WithFields(logrus.Fields{
		"components": components.ApiComponent,
		"layer":      layer,
	}).Info("Request came to fetch the books")

	info := createInformation(s.ct, "List books")

	records, err := s.booksService.ListBooks(ctx, in.Page, in.Limit)
	if err != nil {
		return
	}

	resp := listbooks.Unmarshal(info, records)
	out = &resp

	return
}
