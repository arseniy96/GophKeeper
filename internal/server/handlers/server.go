package handlers

import (
	"context"

	"github.com/arseniy96/GophKeeper/internal/server/storage"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
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
}

func NewServer(r Repository) *Server {
	return &Server{
		Storage: r,
	}
}
