package grpcclient

import (
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/internal/client/interceptors"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

// Client – Объект gRPC клиента для общения с сервером.
type Client struct {
	gRPCClient pb.GophKeeperClient
	config     *config.Config
	authToken  string
	timeout    time.Duration
}

var ErrRequest = errors.New(`request error`)

// NewGRPCClient – функция инициализации gRPC клиента.
func NewGRPCClient(c *config.Config) (*Client, error) {
	client := &Client{
		config:  c,
		timeout: time.Duration(c.ConnectionTimeout) * time.Second,
	}

	conn, err := grpc.Dial(
		c.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor(client)),
	)
	if err != nil {
		return nil, fmt.Errorf("gRPC connection refused: %w", err)
	}
	gRPCClient := pb.NewGophKeeperClient(conn)
	client.gRPCClient = gRPCClient

	return client, nil
}

// GetAuthToken – метод получения AuthToken пользователя.
func (c *Client) GetAuthToken() string {
	return c.authToken
}
