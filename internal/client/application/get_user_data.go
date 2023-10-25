package application

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/arseniy96/GophKeeper/internal/client/utils"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) GetUserData() error {
	err := c.GetUserDataList()
	if err != nil {
		return err
	}

	utils.SlowPrint("Please enter data id")
	var dataID int64
	_, err = fmt.Scan(&dataID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)
	req := &pb.UserDataRequest{Id: dataID}
	res, err := c.ClientGRPC.GetUserData(ctx, req)
	fmt.Printf("data is %v", res.Data)
	buf := bytes.NewBuffer(res.Data)

	var data PasswordStruct

	gob.NewDecoder(buf).Decode(&data)

	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("id: %v, name: %v, type: %v, version: %v, data: %v\n", res.Id, res.Name, res.DataType, res.Version, data)

	return nil
}
