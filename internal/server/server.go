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

	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
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
