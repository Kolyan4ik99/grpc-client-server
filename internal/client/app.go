package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-client-server/internal/config"
	"grpc-client-server/rpc"
)

type App struct {
	cfg    *config.Client
	buffer *buffer
	log    log.Logger
}

func NewApp(cfg *config.Client) *App {
	return &App{
		cfg: cfg,
		buffer: newBuffer(
			cfg.Buffer.Size,
			cfg.Buffer.Threshold),
	}
}

func (a *App) Start() error {
	log.Println("Starting client application")
	ctx, _ := context.WithDeadline(context.Background(), a.cfg.DialDeadline)

	dial, err := grpc.Dial(a.cfg.ServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	dialClient := rpc.NewDialogClient(dial)
	log.Printf("Connection to server [%s]\n", a.cfg.ServerURL)

	cl, err := dialClient.AuthAndListen(ctx, &rpc.Info{
		Name:     a.cfg.UserName,
		Password: a.cfg.Password,
		Interval: a.cfg.DialInterval.Milliseconds(),
	})

	for {
		val, err := cl.Recv()
		if err != nil {
			a.buffer.flush()
			return fmt.Errorf("connection with server was closed")
		}
		go a.buffer.put(val.Index)
		log.Printf("Put value=[%d] to buffer\n", val.Index)
	}
}
