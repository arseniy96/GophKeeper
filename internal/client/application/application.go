package application

import (
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

const (
	AuthCommandType = "auth"
)

type Client struct {
	AuthToken      string
	ClientGRPC     pb.GophKeeperClient
	InputCommands  chan *Command
	OutputCommands chan *Result
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
		ClientGRPC:     client,
		InputCommands:  make(chan *Command, 1),
		OutputCommands: make(chan *Result, 1),
	}

	//go c.processCommands()

	return c
}

//func (c *Client) processCommands() {
//	for {
//		select {
//		case command := <-c.InputCommands:
//			commandType := command.Type
//			if commandType == AuthCommandType {
//				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//				req := &pb.SignInRequest{
//					Login:    "",
//					Password: "",
//				}
//
//				res, err := c.ClientGRPC.SignIn(ctx, req)
//				if err != nil {
//					log.Fatal(err)
//				}
//				c.OutputCommands <- &Result{
//					ErrorMessage: "",
//					Message:      "",
//				}
//				continue
//			}
//			c.OutputCommands <- &Result{
//				ErrorMessage: "unknown command",
//			}
//			continue
//		}
//	}
//}
