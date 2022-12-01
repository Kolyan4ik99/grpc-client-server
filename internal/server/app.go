package server

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"grpc-client-server/internal/config"
	"grpc-client-server/rpc"
)

type App struct {
	cfg *config.Server
}

func (a *App) AuthAndListen(info *rpc.Info, server rpc.Dialog_AuthAndListenServer) error {
	fmt.Printf("password=%s, username=%s\n", info.Password, info.Name)
	dur := time.Duration(info.Interval)
	ticker := time.NewTicker(dur * time.Millisecond)
	index := int64(1)

	for {
		select {
		case <-ticker.C:
			err := server.Send(&rpc.Value{Index: index})
			if err != nil {
				return err
			}
			index++

		case <-server.Context().Done():
			fmt.Printf("Client %s close connection\n", info.Name)
			return nil
		}
	}
}

func NewApp(cfg *config.Server) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() {
	s := grpc.NewServer()
	rpc.RegisterDialogServer(s, a)

	listen, err := net.Listen("tcp", a.cfg.URL)
	if err != nil {
		return
	}
	fmt.Println("Service is successful start, address:", a.cfg.URL)

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
