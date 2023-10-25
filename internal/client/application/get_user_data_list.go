package application

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/arseniy96/GophKeeper/internal/client/utils"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserDataList() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)

	req := &pb.UserDataListRequest{}
	res, err := c.ClientGRPC.GetUserDataList(ctx, req)
	if err != nil {
		return err
	}

	utils.SlowPrint("You have these saved data:")
	for _, el := range res.Data {
		fmt.Printf("id: %v, name: %v, type: %v, version: %v\n", el.Id, el.Name, el.DataType, el.Version)
	}

	return nil
}
