package openapi

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	ct time.Time
}

type Config struct {
	Persistance Persistance
	HTTP        HTTP
}
type HTTP struct {
	Port   int
	Prefix string
}
type Persistance struct {
	Driver   string
	Server   string
	Port     int
	User     string
	Password string
}

func NewServer(ct time.Time) Server {
	return Server{
		ct: ct,
	}
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
