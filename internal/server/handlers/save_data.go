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

func (s *Server) SaveData(ctx context.Context, in *pb.SaveDataRequest) (*pb.SaveDataResponse, error) {
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
	if err = s.Storage.SaveUserData(ctx, user.ID, in.Name, in.DataType, in.Data); err != nil {
		logger.Log.Errorf("save data error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}

	return &pb.SaveDataResponse{
		Result: "OK",
	}, nil
}
