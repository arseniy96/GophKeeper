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
		s.Logger.Log.Errorf("find user error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	if err = s.Storage.SaveUserData(ctx, user.ID, in.Name, in.DataType, in.Data); err != nil {
		s.Logger.Log.Errorf("save data error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SaveDataResponse{
		Result: "OK",
	}, nil
}
