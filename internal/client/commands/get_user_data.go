package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mailru/easyjson"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

func GetUserData(c Client) error {
	err := GetUserDataList(c)
	if err != nil {
		if errors.Is(err, ErrNoData) {
			//nolint:goconst,nolintlint // it's print
			utils.SlowPrint("You have no saved data")
			return nil
		}
		return err
	}

	utils.SlowPrint("Please enter data id")
	var dataID int64
	_, err = fmt.Scan(&dataID)
	if err != nil {
		return err
	}

	data, err := c.GetUserData(models.UserDataModel{ID: dataID})
	if err != nil {
		fmt.Printf("Something went wrong, error: %v", err)
		return nil
	}

	err = printData(data)
	if err != nil {
		return err
	}

	return nil
}

func printData(data *models.UserData) error {
	var pretty []byte
	dataType := data.DataType
	switch dataType {
	case PasswordDataType:
		passData := &PasswordData{}
		err := easyjson.Unmarshal(data.Data, passData)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(passData, "", "  ")
		if err != nil {
			return err
		}
	case CardDataType:
		cardData := &CardData{}
		err := easyjson.Unmarshal(data.Data, cardData)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(cardData, "", "  ")
		if err != nil {
			return err
		}
	case TextDataType:
		textData := &TextData{}
		err := easyjson.Unmarshal(data.Data, textData)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(textData, "", "  ")
		if err != nil {
			return err
		}
	default:
		return nil
	}
	fmt.Printf("\nYour data is:\n%s", pretty)

	return nil
}
