package application

import (
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

type Client struct {
	AuthToken  string
	ClientGRPC pb.GophKeeperClient
}

type Command struct {
	Type string
	Args map[string]string
}

type Result struct {
	ErrorMessage string
	Message      string
}

func NewClient(client pb.GophKeeperClient) *Client {
	c := &Client{
		ClientGRPC: client,
	}

	return c
}
