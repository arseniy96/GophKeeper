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

var (
	DataTypes = [4]string{"password", "card", "file", "text"}
)

type PasswordStruct struct {
	Site     string
	Login    string
	Password string
}

func (c *Client) SaveData() error {
	utils.SlowPrint("What data type do you want to save?")
	for i, dt := range DataTypes {
		fmt.Printf("%v. %v", i, dt)
	}
	var dti int
	_, err := fmt.Scanln(&dti)
	if err != nil {
		return err
	}

	switch dti {
	case 1:
		return c.savePassword()
	case 2:
		return c.saveCardData()
	}

	return nil
}

func (c *Client) savePassword() error {
	var pass PasswordStruct
	var name string
	utils.SlowPrint("Please enter login and password")
	fmt.Printf("site: ")
	_, err := fmt.Scan(&pass.Site)
	if err != nil {
		return err
	}
	fmt.Printf("login: ")
	_, err = fmt.Scan(&pass.Login)
	if err != nil {
		return err
	}
	fmt.Printf("password: ")
	_, err = fmt.Scan(&pass.Password)
	if err != nil {
		return err
	}

	fmt.Println(pass)
	var data bytes.Buffer
	gob.NewEncoder(&data).Encode(pass)
	bytesArr := data.Bytes()
	fmt.Printf("%v", bytesArr)

	fmt.Println("enter name of data ")
	_, err = fmt.Scan(&name)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", c.AuthToken)

	req := &pb.SaveDataRequest{
		Name:     name,
		Data:     bytesArr,
		DataType: DataTypes[1],
	}
	_, err = c.ClientGRPC.SaveData(ctx, req)
	return err
}

func (c *Client) saveCardData() error {
	var login, password, site string
	utils.SlowPrint("Please enter card data")
	fmt.Printf("site: ")
	_, err := fmt.Scan(&site)
	if err != nil {
		return err
	}
	fmt.Printf("login: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return err
	}
	fmt.Printf("password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return err
	}

	return nil
}
