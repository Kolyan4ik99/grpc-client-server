package main

import (
	"log"
	"time"

	"grpc-client-server/internal/client"
	"grpc-client-server/internal/config"
)

func main() {
	cfg, err := config.NewClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	cl := client.NewApp(cfg)

	for {
		err = cl.Start()
		log.Println(err)
		log.Println("Wait 10 second and try reconnect")
		time.Sleep(time.Second * 10)
	}
}
