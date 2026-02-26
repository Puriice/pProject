package main

import (
	"log"
	"pProject/internal/env"
	"pProject/internal/server"
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
