package main

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	"github.com/arseniy96/GophKeeper/internal/handlers"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	serverGRPC := handlers.NewServer(nil)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterGophKeeperServer(gRPCServer, serverGRPC)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer fmt.Println("server has been shutdown")
		defer wg.Done()
		<-ctx.Done()

		fmt.Println("app got a signal")
		gRPCServer.GracefulStop()
	}()

	fmt.Println("gRPC server is running")
	return gRPCServer.Serve(listen)
}
