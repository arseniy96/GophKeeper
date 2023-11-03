package application

import (
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

func (c *Client) userAuth() error {
	var ans string
	c.printer.Print("Do you have an account? (y/n)")
	_, err := c.printer.Scan(&ans)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %v", ErrInternal, err)
	}

	switch ans {
	case "y":
		return c.userSignIn()
	case "n":
		return c.userSignUp()
	default:
		return c.userAuth()
	}
}

func (c *Client) userSignIn() error {
	var (
		login, password string
		err             error
	)

	c.printer.Print("Please enter your login and password.")
	fmt.Print(loginInput)
	_, err = c.printer.Scan(&login)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %v", ErrInternal, err)
	}
	fmt.Print(passwordInput)
	_, err = c.printer.Scan(&password)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %v", ErrInternal, err)
	}

	token, err := c.gRPCClient.SignIn(models.AuthModel{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("%w: SignIn error: %v", ErrUserNotAuthorized, err)
	}
	c.Logger.Log.Debugf("AuthToken: %v", string(token))

	return nil
}

func (c *Client) userSignUp() error {
	var (
		login, password string
		err             error
	)

	c.printer.Print("Please enter login and password.")
	fmt.Print(loginInput)
	_, err = c.printer.Scan(&login)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %v", ErrInternal, err)
	}
	fmt.Print(passwordInput)
	_, err = c.printer.Scan(&password)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %v", ErrInternal, err)
	}

	token, err := c.gRPCClient.SignUp(models.AuthModel{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("%w: SignUp error: %v", ErrUserNotAuthorized, err)
	}
	c.Logger.Log.Debugf("AuthToken: %v", string(token))

	return nil
}
