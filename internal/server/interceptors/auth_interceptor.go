package interceptors

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/server/storage"
	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	"github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
	"github.com/arseniy96/GophKeeper/src/logger"
)

type store interface {
	FindUserByToken(context.Context, string) (*storage.User, error)
}

func AuthInterceptor(s store, l *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == gophkeeper.GophKeeper_SignUp_FullMethodName ||
			info.FullMethod == gophkeeper.GophKeeper_SignIn_FullMethodName ||
			info.FullMethod == gophkeeper.GophKeeper_Ping_FullMethodName {
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
			return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusForbidden))
		}

		encryptedToken := mycrypto.HashFunc(token)
		if _, err := s.FindUserByToken(ctx, encryptedToken); err != nil {
			if errors.Is(err, storage.ErrNowRows) {
				l.Log.Debugf("invalid token: %v", token)
				return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusForbidden))
			}
			l.Log.Errorf("find user error: %v", err)
			return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
		}

		return handler(ctx, req)
	}
}
