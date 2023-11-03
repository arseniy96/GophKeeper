package grpcclient

import (
	"context"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) SaveUserData(model *models.UserData) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.SaveDataRequest{
		Name:     model.Name,
		Data:     model.Data,
		DataType: model.DataType,
	}
	_, err := c.gRPCClient.SaveData(ctx, req)
	if err != nil {
		return fmt.Errorf("%w: gRPC SaveUserData error: %v", ErrRequestErr, err)
	}

	return nil
}
