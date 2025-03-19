package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sirupsen/logrus"
	"github.com/srijilv/go-api-template.git/db"
	"github.com/srijilv/go-api-template.git/pkg/application"
	booksrepo "github.com/srijilv/go-api-template.git/pkg/infrastructure/sql/pgsql/books_repo"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi"
)

type localDbConn struct {
}

func (l localDbConn) GetPgSqlConnectionPool(ctx context.Context, dbName string) (db.PgxIface, error) {
	dsn := fmt.Sprintf("postgres://postgres:pass@postgresdb:5432/%s", dbName)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	return pool, nil
}

func main() {
	l := localDbConn{}
	bookRepo := booksrepo.NewPgsqlBookRepository(l, "learning")
	bookService := application.NewBooksService(bookRepo)
	serviceServer := openapi.NewServer(bookService, render.Render, time.Now())

	RunHTTPServerOnAddr(
		"0.0.0.0:8080",
		func(router chi.Router) http.Handler {
			return openapi.HandlerFromMux(serviceServer, router)
		},
		"/",
	)
}

type Middleware func(http.Handler) http.Handler

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler, apiPrefix string) {

	apiRouter := chi.NewRouter()

	rootRouter := chi.NewRouter()
	rootRouter.Mount(apiPrefix, createHandler(apiRouter))

	logrus.Info("Starting HTTP server")

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func grpcClientDial(key string) (conn *grpc.ClientConn) {
	log.Info("trying to dial client grpc for app clients")
	grpcClients := getGRPCClient(key)

	var interceptor grpc.DialOption

	conn, err := grpc.NewClient(grpcClients,
		grpc.WithTransportCredentials(insecure.NewCredentials()), interceptor,
		// grpc.WithStatsHandler(otelgrpc.NewClientHandler())
	)
	if err != nil {
		panic(err)
	}

	return
}

func getGRPCClient(key string) (client string) {
	return
}
