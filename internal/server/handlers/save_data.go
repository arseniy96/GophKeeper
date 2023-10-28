package handlers

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) SaveData(ctx context.Context, in *pb.SaveDataRequest) (*pb.SaveDataResponse, error) {
	userID := ctx.Value("user_id").(int64)

	if err := s.Storage.SaveUserData(ctx, userID, in.Name, in.DataType, in.Data); err != nil {
		s.Logger.Log.Errorf("save data error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SaveDataResponse{
		Result: "OK",
	}, nil
}
