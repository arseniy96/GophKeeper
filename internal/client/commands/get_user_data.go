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
		passStruct := &PasswordData{}
		err := easyjson.Unmarshal(data.Data, passStruct)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(passStruct, "", "  ")
		if err != nil {
			return err
		}
	case CardDataType:
		cardStruct := &CardData{}
		err := easyjson.Unmarshal(data.Data, cardStruct)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(cardStruct, "", "  ")
		if err != nil {
			return err
		}
	case FileDataType:
		fileStruct := &FileData{}
		err := easyjson.Unmarshal(data.Data, fileStruct)
		if err != nil {
			return nil
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(fileStruct, "", "  ")
		if err != nil {
			return err
		}
	case TextDataType:
		textStruct := &TextData{}
		err := easyjson.Unmarshal(data.Data, textStruct)
		if err != nil {
			return err
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(textStruct, "", "  ")
		if err != nil {
			return err
		}
	default:
		return nil
	}
	fmt.Printf("\nYour data is:\n%s", pretty)

	return nil
}
