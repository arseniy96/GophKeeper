package application

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserDataList() ([]models.UserDataList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)

	req := &pb.UserDataListRequest{}
	res, err := c.ClientGRPC.GetUserDataList(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("get user data list error: %v", err)
		return nil, err
	}

	var records []models.UserDataList
	for _, el := range res.Data {
		rec := models.UserDataList{
			ID:       el.Id,
			Name:     el.Name,
			DataType: el.DataType,
			Version:  el.Version,
		}
		records = append(records, rec)
	}

	return records, nil
}
