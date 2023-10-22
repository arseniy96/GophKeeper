package handlers

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/logger"
	"github.com/arseniy96/GophKeeper/internal/server/storage"
	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) GetUserData(ctx context.Context, in *pb.UserDataRequest) (*pb.UserDataResponse, error) {
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

	record, err := s.Storage.FindUserRecord(ctx, in.Id, user.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNowRows) {
			return nil, status.Error(codes.NotFound, DataNotFound)
		}
		logger.Log.Errorf("find data error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}

	return &pb.UserDataResponse{
		Id:       record.ID,
		Name:     record.Name,
		Data:     record.Data,
		DataType: record.DataType,
		Version:  record.Version,
		CreateAt: record.CreatedAt,
	}, nil
}
