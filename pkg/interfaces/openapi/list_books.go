package openapi

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/srijilv/go-api-template.git/pkg/components"
	listbooks "github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/list_books"
	apierror "github.com/srijilv/go-api-template.git/utils/api_error"
)

const (
	ErrCodeRenderError apierror.Code = 100 + iota
)

func (s Server) ListBooks(w http.ResponseWriter, r *http.Request, params ListBooksParams) {
	logrus.WithFields(
		logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
		},
	)

	info, apiInfoErr := createInfo(s.ct, "List books")

	records, err := s.booksService.ListBooks(r.Context(), params.Page, params.Limit)
	if err != nil {
		apierror.RespondWithApiError(err, apiInfoErr, w, r)
		return
	}

	listBooks := listbooks.ListBooksResponse{}
	listBooks.Unmarshal(records, info)

	render.Status(r, http.StatusOK)
	if err = s.renderer(w, r, listBooks); err != nil {
		err = apierror.NewUnknownError(ErrCodeRenderError, components.ApiComponent, err)
		apierror.RespondWithApiError(err, apiInfoErr, w, r)
		return
	}

	log.Info("list books successfully rendered")
}
