package apierror

import (
	"net/http"

	"github.com/go-chi/render"
)

func RespondWithApiErrorWithoutComponent(err error, info Information, w http.ResponseWriter, r *http.Request) {
	respondWithApiError(err, info, w, r, true)
}

func RespondWithApiError(err error, info Information, w http.ResponseWriter, r *http.Request) {
	respondWithApiError(err, info, w, r, false)
}

func respondWithApiError(err error, info Information, w http.ResponseWriter, r *http.Request, resetComponent bool) {
	apiError, ok := err.(ApiError)
	if !ok {
		InternalServerError(int(UnknownCode), string(UnknownComponent), err, info, w, r, resetComponent)
		return
	}

	switch apiError.ErrorType() {
	case ErrorTypeAuthorization:
		Unauthorised(int(apiError.Code()), string(apiError.Component()), apiError, info, w, r, resetComponent)
	case ErrorTypeIncorrectInput:
		BadRequest(int(apiError.Code()), string(apiError.Component()), apiError, info, w, r, resetComponent)
	case ErrorTypeNotFound:
		NotFound(int(apiError.Code()), string(apiError.Component()), apiError, info, w, r, resetComponent)
	default:
		InternalServerError(int(apiError.Code()), string(apiError.Component()), apiError, info, w, r, resetComponent)
	}
}

func InternalServerError(code int, component string, err error, info Information, w http.ResponseWriter, r *http.Request, resetComponent bool) {
	httpRespondWithError(err, code, component, info, w, r, "internal server error", http.StatusInternalServerError, resetComponent)
}

func Unauthorised(code int, component string, err error, info Information, w http.ResponseWriter, r *http.Request, resetComponent bool) {
	httpRespondWithError(err, code, component, info, w, r, err.Error(), http.StatusUnauthorized, resetComponent)
}

func BadRequest(code int, component string, err error, info Information, w http.ResponseWriter, r *http.Request, resetComponent bool) {
	httpRespondWithError(err, code, component, info, w, r, err.Error(), http.StatusBadRequest, resetComponent)
}

func NotFound(code int, component string, err error, info Information, w http.ResponseWriter, r *http.Request, resetComponent bool) {
	httpRespondWithError(err, code, component, info, w, r, err.Error(), http.StatusNotFound, resetComponent)
}

func httpRespondWithError(err error, code int, component string, info Information, w http.ResponseWriter, r *http.Request, logMsg string, status int, resetComponent bool) {

	if resetComponent {
		component = ""
	}
	resp := HttpErrorResponse{Info: info, Error: HttpError{code, component, logMsg}, httpStatus: status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type HttpErrorResponse struct {
	Info       Information `json:"info"`
	Error      HttpError   `json:"error"`
	httpStatus int
}

func (e HttpErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

type Information struct {
	Name      string  `json:"name"`
	Timestamp float32 `json:"timestamp"`
	Version   string  `json:"version"`
}

type HttpError struct {
	Code      int    `json:"code"`
	Component string `json:"component"`
	Message   string `json:"message"`
}
