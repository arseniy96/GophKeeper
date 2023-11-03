package interceptors

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/GophKeeper/internal/services/mycrypto"
	"github.com/arseniy96/GophKeeper/src"
	"github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
	"github.com/arseniy96/GophKeeper/src/logger"
)

func AuthInterceptor(l *logger.Logger, secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, r interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
		if i.FullMethod == gophkeeper.GophKeeper_SignUp_FullMethodName ||
			i.FullMethod == gophkeeper.GophKeeper_SignIn_FullMethodName ||
			i.FullMethod == gophkeeper.GophKeeper_Ping_FullMethodName {
			return h(ctx, r)
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

		userID, err := mycrypto.GetUserID(token, secret)
		if err != nil {
			l.Log.Debugf("invalid token: %v", token)
			return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusForbidden))
		}

		ctx = context.WithValue(ctx, src.UserIDContextKey, userID)

		return h(ctx, r)
	}
}
