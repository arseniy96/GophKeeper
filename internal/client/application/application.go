package application

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arseniy96/GophKeeper/internal/client/config"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
	"github.com/arseniy96/GophKeeper/src/logger"
)

type Client struct {
	AuthToken  string
	ClientGRPC pb.GophKeeperClient
	Config     *config.Config
	Logger     *logger.Logger
}

func NewClient(l *logger.Logger, c *config.Config) *Client {
	conn, err := grpc.Dial(c.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	gRPCClient := pb.NewGophKeeperClient(conn)

	return &Client{
		ClientGRPC: gRPCClient,
		Config:     c,
		Logger:     l,
	}
}

func (c *Client) UpdateAuthToken(token string) {
	c.AuthToken = token
}
