package application

import (
	"context"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserDataList() ([]models.UserDataList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.GetTimeout())
	defer cancel()

	req := &pb.UserDataListRequest{}
	res, err := c.ClientGRPC.GetUserDataList(ctx, req)
	if err != nil {
		c.Logger.Log.Errorf("get user data list error: %v", err)
		return nil, err
	}

	records := make([]models.UserDataList, 0, len(res.Data))
	for _, el := range res.Data {
		rec := models.UserDataList{
			ID:       el.Id,
			Name:     el.Name,
			DataType: el.DataType,
			Version:  el.Version,
		}
		records = append(records, rec)
	}
	c.SyncCache(records)

	return records, nil
}
