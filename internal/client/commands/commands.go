package commands

import "github.com/arseniy96/GophKeeper/internal/client/models"

const (
	PasswordDataType = "password"
	CardDataType     = "card"
	FileDataType     = "file"
	TextDataType     = "text"
)

type Client interface {
	SignIn(auth models.AuthModel) (models.AuthToken, error)
	SignUp(auth models.AuthModel) (models.AuthToken, error)
	GetUserDataList() ([]models.UserDataList, error)
	GetUserData(models.UserDataModel) (*models.UserData, error)
	SaveUserData(*models.UserData) error
	UpdateAuthToken(token string)
}

type PasswordData struct {
	Site     string `json:"site"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CardData struct {
	Number     string `json:"number"`
	ExpDate    string `json:"exp_date"`
	CardHolder string `json:"card_holder"`
}

type FileData struct {
	Path string `json:"path"`
	Data []byte `json:"data"`
}

type TextData struct {
	Text string `json:"text"`
}
