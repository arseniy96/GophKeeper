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
	encryptedPass := mycrypto.HashFunc(pass)
	err := s.Storage.CreateUser(ctx, login, encryptedPass)
	if err != nil {
		if errors.Is(err, storage.ErrConflict) {
			s.Logger.Log.Debug("user already exists")
			return nil, status.Error(codes.AlreadyExists, http.StatusText(http.StatusConflict))
		}
		s.Logger.Log.Errorf("create user error: %v", err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	authToken, err := s.updateUserToken(ctx, login)
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
	}
	encryptedPass := mycrypto.HashFunc(pass)
	if user.Password != encryptedPass {
		s.Logger.Log.Debugf("authorization failed, login: %v", login)
		return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusUnauthorized))
	}

	authToken, err := s.updateUserToken(ctx, login)
	if err != nil {
		s.Logger.Log.Errorf("update user token error: %v", err) //nolint:goconst,nolintlint // it's format
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SignInResponse{
		Token: authToken,
	}, nil
}

func (s *Server) updateUserToken(ctx context.Context, login string) (string, error) {
	//authToken, err := mycrypto.GenRandomToken()
	//if err != nil {
	//	s.Logger.Log.Errorf("create random token error: %v", err)
	//	return "", err
	//}
	// TODO: implement JWT
	authToken := "882811db55e55c140e728b36ca540e39"
	encryptedToken := mycrypto.HashFunc(authToken)
	err := s.Storage.UpdateUserToken(ctx, login, encryptedToken)
	if err != nil {
		s.Logger.Log.Errorf("update user token error: %v", err) //nolint:goconst,nolintlint // it's format
		return "", err
	}

	return authToken, nil
}
