package main

import (
	"log"

	"grpc-client-server/internal/client"
	"grpc-client-server/internal/config"
)

func main() {
	cfg, err := config.NewClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	cl := client.NewApp(cfg)

	err = cl.Start()
	log.Println(err)
}
