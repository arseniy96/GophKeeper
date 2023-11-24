package grpcclient

import (
	"context"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

// GetUserData – метод получения данных пользователя с сервера.
func (c *Client) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.UserDataRequest{Id: model.ID}
	res, err := c.gRPCClient.GetUserData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: gRPC GetUserData error: %w", ErrRequest, err)
	}

	return &models.UserData{
		ID:       res.Id,
		Name:     res.Name,
		DataType: res.DataType,
		Version:  res.Version,
		Data:     res.Data,
	}, nil
}
