package handlers

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) GetUserDataList(ctx context.Context, in *pb.UserDataListRequest) (*pb.UserDataListResponse, error) {
	var token string
	if meta, ok := metadata.FromIncomingContext(ctx); ok {
		values := meta.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	encryptedToken := mycrypto.HashFunc(token)
	user, err := s.Storage.FindUserByToken(ctx, encryptedToken)
	if err != nil {
		s.Logger.Log.Errorf("find user error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	userRecords, err := s.Storage.GetUserData(ctx, user.ID)
	if err != nil {
		s.Logger.Log.Errorf("get user data error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	if len(userRecords) == 0 {
		return nil, status.Error(codes.NotFound, http.StatusText(http.StatusNoContent))
	}

	var records []*pb.UserDataNested
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
