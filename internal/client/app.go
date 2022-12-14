package client

import (
	"context"
	"fmt"
	"log"
	"time"

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

	dial, err := grpc.Dial(a.cfg.ServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	dialClient := rpc.NewDialogClient(dial)
	log.Printf("Connection to server [%s]\n", a.cfg.ServerURL)

	go a.stopConnectionAfterTime(dialClient)

	func() {
		err := a.listenAndPutToBuffer(dialClient)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (a *App) listenAndPutToBuffer(client rpc.DialogClient) error {
	cl, err := client.Listen(context.Background(), &rpc.Info{
		Name:     a.cfg.UserName,
		Password: a.cfg.Password,
		Interval: a.cfg.DialInterval.Milliseconds(),
	})
	if err != nil {
		return fmt.Errorf("server is not avaiable")
	}

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

func (a *App) stopConnectionAfterTime(client rpc.DialogClient) {
	log.Printf("Wait %s to stop connection", a.cfg.DialDeadline)
	time.Sleep(a.cfg.DialDeadline)
	log.Println("Connection successfully stopped")
	_, err := client.StopListen(context.Background(), &rpc.Empty{})
	if err != nil {
		log.Println(err)
	}
}
