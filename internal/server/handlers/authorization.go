package handlers

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/logger"
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
			logger.Log.Debug("user already exists")
			return nil, status.Error(codes.AlreadyExists, UserAlreadyExistsTxt)
		}
		logger.Log.Errorf("create user error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}

	authToken, err := s.updateUserToken(ctx, login)
	if err != nil {
		logger.Log.Errorf("update user token error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}

	return &pb.SignUpResponse{
		Token: authToken,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	// достаём данные из запроса
	// хэшируем пароль пользователя
	// ищем пользователя по login
	// сверяем с хэшированным паролем пароль из БД
	// если всё ок, создаём токен и отдаём пользователю
	// токен сохраняем в БД
	login := in.Login
	pass := in.Password
	user, err := s.Storage.FindUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, storage.ErrNowRows) {
			logger.Log.Debugf("user with login `%v` not found", login)
			return nil, status.Error(codes.NotFound, UserNotFoundTxt)
		}

	}
	encryptedPass := mycrypto.HashFunc(pass)
	if user.Password != encryptedPass {
		logger.Log.Debugf("authorization failed, login: %v", login)
		return nil, status.Error(codes.Unauthenticated, UserUnauthorizedTxt)
	}

	authToken, err := s.updateUserToken(ctx, login)
	if err != nil {
		logger.Log.Errorf("update user token error: %v", err)
		return nil, status.Error(codes.Internal, InternalBackendErrTxt)
	}

	return &pb.SignInResponse{
		Token: authToken,
	}, nil
}

func (s *Server) updateUserToken(ctx context.Context, login string) (string, error) {
	authToken, err := mycrypto.GenRandomToken()
	if err != nil {
		logger.Log.Errorf("create random token error: %v", err)
		return "", err
	}
	encryptedToken := mycrypto.HashFunc(authToken)
	err = s.Storage.UpdateUserToken(ctx, login, encryptedToken)
	if err != nil {
		logger.Log.Errorf("update user token error: %v", err)
		return "", err
	}

	return authToken, nil
}
