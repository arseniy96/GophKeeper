package handlers

import (
	"context"
	"time"

	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	time.Sleep(5 * time.Second) // for grasfull shutdown test

	return &pb.PingResponse{
		Result: "OK",
	}, nil
}
