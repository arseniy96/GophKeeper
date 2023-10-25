package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arseniy96/GophKeeper/internal/client/application"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
	pb "github.com/arseniy96/GophKeeper/src/grpc/gophkeeper"
)

func main() {
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	gRPCClient := pb.NewGophKeeperClient(conn)

	client := application.NewClient(gRPCClient)

	err = run(client)
	if err != nil {
		panic(err)
	}
}

func run(client *application.Client) error {
	utils.SlowPrint("Hello! I'm GophKeeper. I can save your private information.")

	token, err := client.UserAuth()
	if err != nil {
		return err
	}
	client.AuthToken = token

	fmt.Printf("Your token is %v\n", token)

	err = startSession(client)
	if err != nil {
		return err
	}

	return nil
}

func startSession(client *application.Client) error {
	for {
		utils.SlowPrint("Choose command (enter number of command)")
		fmt.Println("1. Get all saved data")
		fmt.Println("2. Get some saved data")
		fmt.Println("3. Save some data")
		fmt.Println("4. Edit saved data")

		var commandNumber int
		_, err := fmt.Scanln(&commandNumber)
		if err != nil {
			return err
		}

		switch commandNumber {
		case 1:
			err := client.GetUserDataList()
			if err != nil {
				return err
			}
		case 2:
			err := client.GetUserData()
			if err != nil {
				return err
			}
		case 3:
			err := client.SaveData()
			if err != nil {
				return err
			}
		}
		fmt.Printf("\n\n====================\n\n")
	}
}
