package grpc

import (
	"time"

	"github.com/srijilv/go-api-template.git/pkg/application"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/common"
)

const (
	ApiVersion    = "1.0"
	layer         = "interfaces"
	interfaceType = "openapi"
)

type Server struct {
	booksService application.BookService
	ct           time.Time
	UnimplementedBooksServiceServer
}

func NewServer(booksService application.BookService, ct time.Time) Server {
	if booksService == nil {
		panic("book service is nil")
	}

	return Server{
		booksService: booksService,
		ct:           ct,
	}
}

func createInformation(t time.Time, apiInfoName string) (info *common.Information) {
	info = &common.Information{
		Name:      apiInfoName,
		Timestamp: float32(t.Unix()),
		Version:   ApiVersion,
	}

	// grpcErr = &grpcerr.Information{
	// 	Version:   info.Version,
	// 	Name:      info.Name,
	// 	Timestamp: info.Timestamp,
	// }

	return
}
