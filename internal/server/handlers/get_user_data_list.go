package handlers

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) GetUserDataList(ctx context.Context, in *pb.UserDataListRequest) (*pb.UserDataListResponse, error) {
	userID := ctx.Value("user_id").(int64)

	userRecords, err := s.Storage.GetUserData(ctx, userID)
	if err != nil {
		s.Logger.Log.Errorf("get user data error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	if len(userRecords) == 0 {
		return nil, status.Error(codes.NotFound, http.StatusText(http.StatusNoContent))
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
