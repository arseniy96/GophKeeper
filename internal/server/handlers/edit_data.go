package handlers

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/storage"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) UpdateUserData(ctx context.Context, in *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	userID := ctx.Value("user_id").(int64)

	record, err := s.Storage.FindUserRecord(ctx, in.Id, userID)
	if err != nil {
		if errors.Is(err, storage.ErrNowRows) {
			return nil, status.Error(codes.NotFound, http.StatusText(http.StatusNoContent))
		}
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	if record.Version != in.Version {
		s.Logger.Log.Errorf("data version conflict, current version is %v", record.Version)
		return nil, status.Error(codes.FailedPrecondition, http.StatusText(http.StatusConflict))
	}

	record.Data = in.Data
	err = s.Storage.UpdateUserRecord(ctx, record)
	if err != nil {
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return &pb.UpdateUserDataResponse{
		Result: "OK",
	}, nil
}
