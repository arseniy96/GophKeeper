package commands

import (
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

func UserAuth(c Client) (string, error) {
	var ans string
	utils.SlowPrint("Do you have an account? (y/n)")
	_, err := fmt.Scanln(&ans)
	if err != nil {
		return "", err
	}

	switch ans {
	case "y":
		return userSignIn(c)
	case "n":
		return userSignUp(c)
	default:
		return UserAuth(c)
	}
}

func userSignIn(c Client) (string, error) {
	var (
		login, password string
		err             error
	)

	utils.SlowPrint("Please enter your login and password.")
	fmt.Print(loginInput)
	_, err = fmt.Scan(&login)
	if err != nil {
		return "", err
	}
	fmt.Print(passwordInput)
	_, err = fmt.Scan(&password)
	if err != nil {
		return "", err
	}

	token, err := c.SignIn(models.AuthModel{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", ErrUserNotAuthorized
	}

	return string(token), nil
}

func userSignUp(c Client) (string, error) {
	var (
		login, password string
		err             error
	)

	utils.SlowPrint("Please enter login and password.")
	fmt.Print(loginInput)
	_, err = fmt.Scan(&login)
	if err != nil {
		return "", err
	}
	fmt.Print(passwordInput)
	_, err = fmt.Scan(&password)
	if err != nil {
		return "", err
	}

	token, err := c.SignUp(models.AuthModel{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", ErrUserNotAuthorized
	}

	return string(token), nil
}
