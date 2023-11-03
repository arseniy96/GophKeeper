package handlers

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/src"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) GetUserDataList(ctx context.Context, in *pb.UserDataListRequest) (*pb.UserDataListResponse, error) {
	userID := ctx.Value(src.UserIDContextKey).(int64)

	userRecords, err := s.Storage.GetUserData(ctx, userID)
	if err != nil {
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	records := make([]*pb.UserDataNested, 0, len(userRecords))
	for _, rec := range userRecords {
		data := &pb.UserDataNested{
			Id:       rec.ID,
			Name:     rec.Name,
			DataType: rec.DataType,
			Version:  rec.Version,
			CreateAt: rec.CreatedAt,
		}
		records = append(records, data)
	}

	return &pb.UserDataListResponse{
		Data: records,
	}, nil
}
