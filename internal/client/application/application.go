package application

import (
	"github.com/arseniy96/GophKeeper/internal/client/clientcache"
	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/internal/client/grpcclient"
	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
	"github.com/arseniy96/GophKeeper/src/logger"
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

type clientCache interface {
	Append(data *models.UserData)
	GetUserData(model models.UserDataModel) (*models.UserData, error)
	GetUserDataList() []models.UserDataList
}

// Client – структура, которая служит для работы с сервером. Отвечает за сценарий приложения.
type Client struct {
	gRPCClient grpcClient
	printer    printer
	cache      clientCache
	Config     *config.Config
	Logger     *logger.Logger
}

// NewClient – функция для инициализации клиента.
// Принимает логгер, для внутреннего логгирования и конфигурацию для клиента.
func NewClient(l *logger.Logger, c *config.Config) (*Client, error) {
	gRPCClient, err := grpcclient.NewGRPCClient(c)
	if err != nil {
		return nil, err
	}
	cache := clientcache.NewCache()

	return &Client{
		gRPCClient: gRPCClient,
		printer:    &utils.Printer{},
		cache:      cache,
		Config:     c,
		Logger:     l,
	}, nil
}
