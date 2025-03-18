package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/srijilv/go-api-template.git/db"
	"github.com/srijilv/go-api-template.git/pkg/application"
	booksrepo "github.com/srijilv/go-api-template.git/pkg/infrastructure/sql/pgsql/books_repo"
	srvcgrpc "github.com/srijilv/go-api-template.git/pkg/interfaces/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Info("starting grpc server of books")
	initServer()
}

type localDbConn struct {
}

func (l localDbConn) GetPgSqlConnectionPool(ctx context.Context, dbName string) (db.PgxIface, error) {
	dsn := fmt.Sprintf("postgres://postgres:pass@localhost:5432/%s", dbName)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	return pool, nil
}
func initServer() {
	// ctx := context.Background()

	l := localDbConn{}
	bookRepo := booksrepo.NewPgsqlBookRepository(l, "learning")
	bookService := application.NewBooksService(bookRepo)

	serviceServer := srvcgrpc.NewServer(bookService, time.Now())

	// Start grpc server
	startServer(serviceServer)
}

func startServer(serviceServer srvcgrpc.Server) {
	log.Info("trying to register grpc server for push")

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	srvcgrpc.RegisterBooksServiceServer(grpcServer, serviceServer)
	log.Info("grpc server for books registered successfully")

	grpcPort := 50051
	if grpcPort == 0 {
		grpcPort = 5000
	}

	log.Infof("trying to listen grpc server for books in port %d: ", grpcPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("server listening at %v: ", listener.Addr())
	log.Info("starting grpc server for books")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
