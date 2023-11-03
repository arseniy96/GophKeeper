package handlers

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/storage"
	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	login := in.Login
	pass := in.Password
	encryptedPass, err := mycrypto.HashFunc(pass)
	if err != nil {
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	userID, err := s.Storage.CreateUser(ctx, login, encryptedPass)
	if err != nil {
		if errors.Is(err, storage.ErrConflict) {
			s.Logger.Log.Debug("user already exists")
			return nil, status.Error(codes.AlreadyExists, http.StatusText(http.StatusConflict))
		}
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	authToken, err := mycrypto.BuildJWT(userID, s.Config.SecretKey)
	if err != nil {
		s.Logger.Log.Errorf("update user token error: %v", err) //nolint:goconst,nolintlint // it's format
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SignUpResponse{
		Token: authToken,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	login := in.Login
	pass := in.Password
	user, err := s.Storage.FindUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, storage.ErrNowRows) {
			s.Logger.Log.Debugf("user with login `%v` not found", login)
			return nil, status.Error(codes.NotFound, http.StatusText(http.StatusUnauthorized))
		}
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	if err := mycrypto.CompareHash(pass, user.Password); err != nil {
		s.Logger.Log.Debugf("authorization failed, login: %v", login)
		return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusUnauthorized))
	}

	authToken, err := mycrypto.BuildJWT(user.ID, s.Config.SecretKey)
	if err != nil {
		s.Logger.Log.Errorf("update user token error: %v", err) //nolint:goconst,nolintlint // it's format
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SignInResponse{
		Token: authToken,
	}, nil
}
