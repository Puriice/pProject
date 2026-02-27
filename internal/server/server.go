package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/httplibs/pkg/middleware"
	"github.com/puriice/httplibs/pkg/middleware/cors"
	"github.com/puriice/pProject/internal/hander/project"
)

type Server struct {
	Host string
	Port string
	db   *pgxpool.Pool
}

func NewServer(host string, port string) (*Server, error) {
	var connectionString = os.Getenv("DB_URL")

	if connectionString == "" {
		return nil, errors.New("Invalid connection string.")
	}

	db, err := pgxpool.New(context.Background(), connectionString)

	if err != nil {
		return nil, err
	}

	return &Server{
		Host: host,
		Port: port,
		db:   db,
	}, nil
}

func Start(server *Server) {
	address := fmt.Sprintf("%s:%s", server.Host, server.Port)

	router := http.NewServeMux()

	projectModel := project.NewModel(server.db)
	projectHandler := project.NewHandler(projectModel)

	projectHandler.RegisterRoute(router)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	corsOption := cors.NewCorsOptions()
	corsOption.AllowOrigins = cors.Wildcard()
	corsOption.AllowNoOrigin = true
	corsOption.AllowCredentials = true

	pipeline := middleware.Pipe(
		middleware.Logger,
		middleware.Cors(*corsOption),
	)

	httpServer := &http.Server{
		Addr:    address,
		Handler: pipeline(mux),
	}

	log.Println("server listening on", address)

	go httpServer.ListenAndServe()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpServer.Shutdown(ctx)
}
