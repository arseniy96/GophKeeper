package main

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	"github.com/arseniy96/GophKeeper/internal/server/config"
	"github.com/arseniy96/GophKeeper/internal/server/handlers"
	"github.com/arseniy96/GophKeeper/internal/server/interceptors"
	"github.com/arseniy96/GophKeeper/internal/server/logger"
	"github.com/arseniy96/GophKeeper/internal/server/storage"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	err := config.Initialize()
	if err != nil {
		return err
	}
	err = logger.Initialize(config.Settings.LogLevel)
	if err != nil {
		return err
	}

	rep, err := storage.NewStorage(config.Settings.DatabaseDSN)
	if err != nil {
		logger.Log.Errorf("database error: %v", err)
	}
	defer func() {
		if err = rep.Close(); err != nil {
			logger.Log.Errorf("database close error: %v", err)
		}
	}()

	serverGRPC := handlers.NewServer(rep)

	listen, err := net.Listen("tcp", config.Settings.Host)
	if err != nil {
		return err
	}
	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(interceptors.AuthInterceptor(rep)))
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
