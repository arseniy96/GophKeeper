package grpcclient

import (
	"context"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) UpdateUserData(model *models.UserData) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.UpdateUserDataRequest{
		Id:      model.ID,
		Data:    model.Data,
		Version: model.Version,
	}
	_, err := c.gRPCClient.UpdateUserData(ctx, req)
	if err != nil {
		return fmt.Errorf("%w: gRPC UpdataUserData error: %v", ErrRequestErr, err)
	}

	return nil
}
