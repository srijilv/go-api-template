package openapi

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/srijilv/go-api-template.git/pkg/application"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/common"
	apierror "github.com/srijilv/go-api-template.git/utils/api_error"
)

const (
	ApiVersion    = "1.0"
	layer         = "interfaces"
	interfaceType = "openapi"
)

type Renderer func(w http.ResponseWriter, r *http.Request, v render.Renderer) error

type Server struct {
	booksService application.BookService
	renderer     Renderer
	ct           time.Time
}

func NewServer(booksService application.BookService, renderer Renderer, ct time.Time) Server {
	if booksService == nil {
		panic("book service is nil")
	}

	if renderer == nil {
		renderer = render.Render
	}

	return Server{
		booksService: booksService,
		renderer:     renderer,
		ct:           ct,
	}
}
func createInfo(t time.Time, apiName string) (info common.Information, apiErrorInfo apierror.Information) {
	info = common.Information{
		Name:      apiName,
		Timestamp: float32(t.Unix()),
		Version:   ApiVersion,
	}

	apiErrorInfo = apierror.Information(info)
	return
}
