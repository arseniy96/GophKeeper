package grpcclient

import (
	"context"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserDataList() ([]models.UserDataList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.UserDataListRequest{}
	res, err := c.gRPCClient.GetUserDataList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: gRPC GetUserDataList error: %v", ErrRequestErr, err)
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
	//c.SyncCache(records) //TODO

	return records, nil
}
