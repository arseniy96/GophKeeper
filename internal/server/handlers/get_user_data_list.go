package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/logger"
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
		logger.Log.Errorf("find user error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}
	userRecords, err := s.Storage.GetUserData(ctx, user.ID)
	if err != nil {
		logger.Log.Errorf("get user data error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}
	if len(userRecords) == 0 {
		return nil, status.Error(codes.NotFound, DataNotFound)
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
