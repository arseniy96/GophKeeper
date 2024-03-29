package application

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// Start – функция для начала сценария.
// Сначала пользователь должен авторизоваться.
// Затем пользователю предлагают выбрать одну из команд:
//  1. Получение всех сохранённых данных (мета-данных).
//  2. Получение данных.
//  3. Сохранение данных.
//  4. Редактирование данных.
func (c *Client) Start() error {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	c.printer.Print("Hello! I'm GophKeeper. I can save your private information.")

	if err := c.UserAuth(); err != nil {
		c.Logger.Log.Error(err)
		return err
	}
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return c.startSession(ctx)
	})

	return grp.Wait()
}

func (c *Client) startSession(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.Logger.Log.Info("client has been shutdown")
			return nil
		default:
			c.printer.Print("Choose command (enter number of command)")
			fmt.Println("1. Get all saved data")
			fmt.Println("2. Get some saved data")
			fmt.Println("3. Save some data")
			fmt.Println("4. Edit saved data")

			var commandNumber int
			_, err := c.printer.Scan(&commandNumber)
			if err != nil {
				return err
			}

			switch commandNumber {
			case getUserDataList:
				err := c.GetUserDataList()
				if err != nil {
					if errors.Is(err, ErrNoData) {
						//nolint:goconst,nolintlint // it's print
						c.printer.Print("You have no saved data")
						continue
					}
					c.Logger.Log.Error(err)
					continue
				}
			case getUserData:
				err := c.GetUserData()
				if err != nil {
					c.Logger.Log.Error(err)
					continue
				}
			case saveUserData:
				err := c.SaveData()
				if err != nil {
					c.Logger.Log.Error(err)
					continue
				}
			case editUserData:
				err := c.EditData()
				if err != nil {
					c.Logger.Log.Error(err)
					continue
				}
			default:
				fmt.Println("Unknown command")
			}
			fmt.Printf("\n====================\n\n")
		}
	}
}
