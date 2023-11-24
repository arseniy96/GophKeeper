package application

import (
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

// UserAuth – функция авторизации пользователя.
func (c *Client) UserAuth() error {
	var ans string
	c.printer.Print("Do you have an account? (y/n)")
	_, err := c.printer.Scan(&ans)
	if err != nil {
		return fmt.Errorf(InternalErrTemplate, ErrInternal, err)
	}

	switch ans {
	case "y":
		authM, err := buildAuthData(c.printer)
		if err != nil {
			return err
		}
		return c.userSignIn(*authM)
	case "n":
		authM, err := buildAuthData(c.printer)
		if err != nil {
			return err
		}
		return c.userSignUp(*authM)
	default:
		return c.UserAuth()
	}
}

func (c *Client) userSignIn(authM models.AuthModel) error {
	_, err := c.gRPCClient.SignIn(authM)
	if err != nil {
		return fmt.Errorf("%w: SignIn error: %w", ErrUserNotAuthorized, err)
	}

	return nil
}

func (c *Client) userSignUp(authM models.AuthModel) error {
	_, err := c.gRPCClient.SignUp(authM)
	if err != nil {
		return fmt.Errorf("%w: SignUp error: %w", ErrUserNotAuthorized, err)
	}

	return nil
}

func buildAuthData(p printer) (*models.AuthModel, error) {
	var (
		login, password string
		err             error
	)

	p.Print("Please enter your login and password.")
	fmt.Print(loginInput)
	_, err = p.Scan(&login)
	if err != nil {
		return nil, fmt.Errorf(InternalErrTemplate, ErrInternal, err)
	}
	fmt.Print(passwordInput)
	_, err = p.Scan(&password)
	if err != nil {
		return nil, fmt.Errorf(InternalErrTemplate, ErrInternal, err)
	}

	return &models.AuthModel{
		Login:    login,
		Password: password,
	}, nil
}
