package main

import (
	"grpc-client-server/internal/config"
	"grpc-client-server/internal/server"
)

func main() {
	cfg := config.NewServerConfig()

	serv := server.NewApp(cfg)

	serv.Start()
}
