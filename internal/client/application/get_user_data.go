package application

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mailru/easyjson"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

func (c *Client) GetUserData() error {
	err := c.GetUserDataList()
	if err != nil {
		if errors.Is(err, ErrNoData) {
			//nolint:goconst,nolintlint // it's print
			c.printer.Print("You have no saved data")
			return nil
		}
		return err
	}

	c.printer.Print("Please enter data id")
	var (
		data   *models.UserData
		dataID int64
	)
	_, err = c.printer.Scan(&dataID)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	m := models.UserDataModel{ID: dataID}
	data, err = c.cache.GetUserData(m) // сначала пытаемся достать из кеша
	if err != nil {
		c.Logger.Log.Warnf("get data from cache error: %w", err)
		data, err = c.gRPCClient.GetUserData(m) // если в кеше нет, идём на сервер
		if err != nil {
			return err
		}
		c.cache.Append(data) // складываем в кеш
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
	case passwordDataType:
		passStruct := &PasswordData{}
		err := easyjson.Unmarshal(data.Data, passStruct)
		if err != nil {
			return fmt.Errorf("easyjson unmarshal error: %w", err)
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(passStruct, "", "  ")
		if err != nil {
			return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
		}
	case cardDataType:
		cardStruct := &CardData{}
		err := easyjson.Unmarshal(data.Data, cardStruct)
		if err != nil {
			return fmt.Errorf("easyjson unmarshal error: %w", err)
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(cardStruct, "", "  ")
		if err != nil {
			return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
		}
	case fileDataType:
		fileStruct := &FileData{}
		err := easyjson.Unmarshal(data.Data, fileStruct)
		if err != nil {
			return fmt.Errorf("easyjson unmarshal error: %w", err)
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(fileStruct, "", "  ")
		if err != nil {
			return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
		}
	case textDataType:
		textStruct := &TextData{}
		err := easyjson.Unmarshal(data.Data, textStruct)
		if err != nil {
			return fmt.Errorf("easyjson unmarshal error: %w", err)
		}
		//nolint:goconst,nolintlint // it's format
		pretty, err = json.MarshalIndent(textStruct, "", "  ")
		if err != nil {
			return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
		}
	default:
		return nil
	}
	fmt.Printf("\nYour data is:\n%s", pretty)

	return nil
}
