package application

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/internal/client/interceptors"
	"github.com/arseniy96/GophKeeper/internal/client/models"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
	"github.com/arseniy96/GophKeeper/src/logger"
)

const (
	DataIDSyncChanSize = 5
)

type Client struct {
	ClientGRPC     pb.GophKeeperClient
	Config         *config.Config
	Logger         *logger.Logger
	Cache          []clientCache
	DataIDSyncChan chan int64
	AuthToken      string
}

type clientCache struct {
	token  string
	dataID int64
	data   *models.UserData
	actual bool
}

func NewClient(l *logger.Logger, c *config.Config) (*Client, func() error) {
	client := &Client{
		Config:         c,
		Logger:         l,
		DataIDSyncChan: make(chan int64, DataIDSyncChanSize),
	}

	conn, err := grpc.Dial(
		c.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor(client)),
	)
	if err != nil {
		log.Fatal(err)
	}
	gRPCClient := pb.NewGophKeeperClient(conn)
	client.ClientGRPC = gRPCClient

	go client.DataSyncWorker()

	return client, conn.Close
}

func (c *Client) UpdateAuthToken(token string) {
	c.AuthToken = token
}

func (c *Client) GetAuthToken() string {
	return c.AuthToken
}

func (c *Client) GetTimeout() time.Duration {
	return time.Duration(c.Config.ConnectionTimeout) * time.Second
}
