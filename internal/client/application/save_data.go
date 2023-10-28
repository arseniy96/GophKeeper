package application

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) SaveUserData(model *models.UserData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)

	req := &pb.SaveDataRequest{
		Name:     model.Name,
		Data:     model.Data,
		DataType: model.DataType,
	}
	_, err := c.ClientGRPC.SaveData(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("save user data error: %v", err)
		return err
	}

	return nil
}
