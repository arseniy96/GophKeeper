package commands

import (
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

func UserAuth(c Client) (string, error) {
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

	token, err := c.SignIn(models.AuthModel{
		Login:    login,
		Password: password,
	})
	if err != nil {
		fmt.Printf("login error: %v", err)
		return "", err
	}

	return string(token), nil
}
