package openapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	ct time.Time
}

func NewServer(ct time.Time) Server {
	return Server{
		ct: ct,
	}
}

func (s Server) PostAccountsRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostAccountsRegister api  works")

}

func RunHTTPServer(addr string, s Server, prefixPath string) (err error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	handler := HandlerFromMux(s, r)
	rootRouter := chi.NewRouter()
	rootRouter.Mount(prefixPath, handler)

	err = http.ListenAndServe(addr, rootRouter)
	return
}
