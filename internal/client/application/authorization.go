package application

import (
	"context"
	"time"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) SignIn(model models.AuthModel) (models.AuthToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SignInRequest{
		Login:    model.Login,
		Password: model.Password,
	}
	res, err := c.ClientGRPC.SignIn(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("SignIn error: %v", err)
		return "", err
	}

	return models.AuthToken(res.Token), nil
}

func (c *Client) SignUp(model models.AuthModel) (models.AuthToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SignUpRequest{
		Login:    model.Login,
		Password: model.Password,
	}
	res, err := c.ClientGRPC.SignUp(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("SignUp error: %v", err)
		return "", err
	}

	return models.AuthToken(res.Token), nil
}
