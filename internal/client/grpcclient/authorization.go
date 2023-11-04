package grpcclient

import (
	"context"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

// SignIn – метод логина пользователя на сервере.
func (c *Client) SignIn(model models.AuthModel) (models.AuthToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.SignInRequest{
		Login:    model.Login,
		Password: model.Password,
	}
	res, err := c.gRPCClient.SignIn(ctx, req)
	if err != nil {
		return "", fmt.Errorf("%w: gRPC SignIn error: %w", ErrRequest, err)
	}
	c.authToken = res.Token

	return models.AuthToken(res.Token), nil
}

// SignUp – метод регистрации пользователя на сервере.
func (c *Client) SignUp(model models.AuthModel) (models.AuthToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.SignUpRequest{
		Login:    model.Login,
		Password: model.Password,
	}
	res, err := c.gRPCClient.SignUp(ctx, req)
	if err != nil {
		return "", fmt.Errorf("%w: gRPC SignUp error: %w", ErrRequest, err)
	}
	c.authToken = res.Token

	return models.AuthToken(res.Token), nil
}
