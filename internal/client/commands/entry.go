package commands

import (
	"errors"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

const (
	getUserDataList = iota + 1
	getUserData
	saveUserData
	editUserData
)

func Start(c Client) error {
	utils.SlowPrint("Hello! I'm GophKeeper. I can save your private information.")

	token, err := UserAuth(c)
	if err != nil {
		return err
	}
	c.UpdateAuthToken(token)
	fmt.Printf("Your token is %v\n", token)

	return startSession(c)
}

func startSession(client Client) error {
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
		case getUserDataList:
			err := GetUserDataList(client)
			if err != nil {
				if errors.Is(err, ErrNoData) {
					//nolint:goconst,nolintlint // it's print
					utils.SlowPrint("You have no saved data")
					return nil
				}
				return err
			}
		case getUserData:
			err := GetUserData(client)
			if err != nil {
				return err
			}
		case saveUserData:
			err := SaveData(client)
			if err != nil {
				return err
			}
		case editUserData:
			err := EditData(client)
			if err != nil {
				return err
			}
		default:
			fmt.Println("Unknown command")
		}
		fmt.Printf("\n====================\n\n")
	}
}
