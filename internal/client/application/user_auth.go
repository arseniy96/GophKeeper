package application

import (
	"context"
	"fmt"
	"time"

	"github.com/arseniy96/GophKeeper/internal/client/utils"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func (c *Client) UserAuth() (string, error) {
	var (
		login, password string
		err             error
	)

	utils.SlowPrint("Please enter your login and password.")
	fmt.Printf("login: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return "", err
	}
	fmt.Printf("password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return "", err
	}

	token, err := c.signInRequest(login, password)
	if err != nil {
		fmt.Printf("login error: %v", err)
		return "", err
	}

	return token, nil
}

func (c *Client) signInRequest(login, pass string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SignInRequest{
		Login:    login,
		Password: pass,
	}
	res, err := c.ClientGRPC.SignIn(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Token, nil
}
