package main

import (
	"log"

	"github.com/puriice/pProject/internal/env"
	"github.com/puriice/pProject/internal/server"
)

func main() {
	env.InitEnv()

	host := env.GetEnv("HOST", "localhost")
	port := env.GetEnv("PORT", "8080")

	serv, err := server.NewServer(host, port)

	if err != nil {
		log.Fatal(err)
		return
	}

	server.Start(serv)
}
