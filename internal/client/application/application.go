package application

import (
	"errors"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/internal/client/grpcclient"
	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
	"github.com/arseniy96/GophKeeper/src/logger"
)

const (
	DataIDSyncChanSize = 5
)

type printer interface {
	Print(s string)
	Scan(a ...interface{}) (int, error)
}

type grpcClient interface {
	SignIn(model models.AuthModel) (models.AuthToken, error)
	SignUp(model models.AuthModel) (models.AuthToken, error)
	GetUserData(model models.UserDataModel) (*models.UserData, error)
	GetUserDataList() ([]models.UserDataList, error)
	SaveUserData(model *models.UserData) error
	UpdateUserData(model *models.UserData) error
}

type Client struct {
	gRPCClient grpcClient
	printer    printer
	Config     *config.Config
	Logger     *logger.Logger
}

type clientCache struct {
	token  string
	dataID int64
	data   *models.UserData
	actual bool
}

func NewClient(l *logger.Logger, c *config.Config) (*Client, error) {
	gRPCClient, err := grpcclient.NewGRPCClient(c)
	if err != nil {
		return nil, err
	}

	//go client.DataSyncWorker() TODO
	return &Client{
		gRPCClient: gRPCClient,
		printer:    &utils.Printer{},
		Config:     c,
		Logger:     l,
	}, nil
}

func (c *Client) Start() error {
	c.printer.Print("Hello! I'm GophKeeper. I can save your private information.")

	if err := c.userAuth(); err != nil {
		c.Logger.Log.Error(err)
		return err
	}
	return c.startSession()
}

func (c *Client) startSession() error {
	for {
		c.printer.Print("Choose command (enter number of command)")
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
