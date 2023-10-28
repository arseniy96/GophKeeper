package application

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)

	req := &pb.UserDataRequest{Id: model.ID}
	res, err := c.ClientGRPC.GetUserData(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("get user data error: %v", err)
		return nil, err
	}

	return &models.UserData{
		ID:       res.Id,
		Name:     res.Name,
		DataType: res.DataType,
		Version:  res.Version,
		Data:     res.Data,
	}, nil
}
