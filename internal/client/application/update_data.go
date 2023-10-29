package application

import (
	"context"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) UpdateUserData(model *models.UserData) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.GetTimeout())
	defer cancel()

	req := &pb.UpdateUserDataRequest{
		Id:      model.ID,
		Data:    model.Data,
		Version: model.Version,
	}
	_, err := c.ClientGRPC.UpdateUserData(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("update user data error: %v", err)
		return err
	}

	return nil
}
