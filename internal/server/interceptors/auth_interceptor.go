package interceptors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/handlers"
	"github.com/arseniy96/GophKeeper/internal/server/logger"
	"github.com/arseniy96/GophKeeper/internal/server/storage"
	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	"github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func AuthInterceptor(s *storage.Database) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == gophkeeper.GophKeeper_SignUp_FullMethodName || info.FullMethod == gophkeeper.GophKeeper_SignIn_FullMethodName {
			return handler(ctx, req)
		}

		var token string
		if meta, ok := metadata.FromIncomingContext(ctx); ok {
			values := meta.Get("token")
			if len(values) > 0 {
				token = values[0]
			}
		}
		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, handlers.MissingTokenTxt)
		}

		encryptedToken := mycrypto.HashFunc(token)
		if _, err := s.FindUserByToken(ctx, encryptedToken); err != nil {
			if errors.Is(err, storage.ErrNowRows) {
				logger.Log.Debugf("invalid token: %v", token)
				return nil, status.Error(codes.Unauthenticated, handlers.InvalidTokenTxt)
			}
			logger.Log.Errorf("find user error: %v", err)
			return nil, status.Error(codes.Internal, handlers.InternalBackendErrTxt)
		}

		return handler(ctx, req)
	}
}
