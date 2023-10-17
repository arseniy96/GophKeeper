package handlers

import (
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

type Repository interface {
	HealthCheck() error
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
