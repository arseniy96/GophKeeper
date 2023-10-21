package handlers

import (
	"context"

	"github.com/arseniy96/GophKeeper/internal/server/storage"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

type Repository interface {
	HealthCheck() error
	CreateUser(context.Context, string, string) error
	UpdateUserToken(context.Context, string, string) error
	FindUserByLogin(context.Context, string) (*storage.User, error)
	FindUserByToken(context.Context, string) (*storage.User, error)
}

type Server struct {
	pb.UnimplementedGophKeeperServer
	Storage Repository
}

func NewServer(r Repository) *Server {
	return &Server{
		Storage: r,
	}
}
