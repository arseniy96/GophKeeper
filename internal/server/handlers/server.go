package handlers

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"sync"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/arseniy96/GophKeeper/internal/server/config"
	"github.com/arseniy96/GophKeeper/internal/server/interceptors"
	"github.com/arseniy96/GophKeeper/internal/server/storage"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
	"github.com/arseniy96/GophKeeper/src/logger"
)

type Repository interface {
	HealthCheck() error
	CreateUser(ctx context.Context, login, password string) error
	UpdateUserToken(ctx context.Context, login, token string) error
	FindUserByLogin(ctx context.Context, login string) (*storage.User, error)
	FindUserByToken(ctx context.Context, token string) (*storage.User, error)
	SaveUserData(ctx context.Context, userID int64, name, dataType string, data []byte) error
	GetUserData(ctx context.Context, userID int64) ([]storage.ShortRecord, error)
	FindUserRecord(ctx context.Context, id, userID int64) (*storage.Record, error)
}

type Server struct {
	pb.UnimplementedGophKeeperServer
	Storage Repository
	Config  *config.Config
	Logger  *logger.Logger
}

func NewServer(r Repository, c *config.Config, l *logger.Logger) *Server {
	return &Server{
		Storage: r,
		Config:  c,
		Logger:  l,
	}
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", s.Config.Host)
	if err != nil {
		s.Logger.Log.Error(err)
		return fmt.Errorf("tcp connection failed")
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.AuthInterceptor(s.Storage, s.Logger),
		logging.UnaryServerInterceptor(interceptors.LoggerInterceptor()),
	))

	pb.RegisterGophKeeperServer(gRPCServer, s)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer s.Logger.Log.Info("server has been shutdown")
		defer wg.Done()
		<-ctx.Done()

		s.Logger.Log.Info("app got a signal")
		gRPCServer.GracefulStop()
	}()

	s.Logger.Log.Info("gRPC server is running")

	return gRPCServer.Serve(listen)
}
