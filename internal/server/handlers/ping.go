package handlers

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	time.Sleep(5 * time.Second) // for grasfull shutdown test
	err := s.Storage.HealthCheck()
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalBackendErrTxt)
	}

	return &pb.PingResponse{
		Result: "OK",
	}, nil
}
