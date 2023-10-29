package commands

import (
	"errors"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

var dataTypes = [4]string{PasswordDataType, CardDataType, FileDataType, TextDataType}

func EditData(c Client) error {
	err := GetUserDataList(c)
	if err != nil {
		if errors.Is(err, ErrNoData) {
			//nolint:goconst,nolintlint // it's print
			utils.SlowPrint("You have no saved data")
			return nil
		}
		return err
	}

	utils.SlowPrint("Please enter data id for update")
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
	dti := 0
	switch data.DataType {
	case PasswordDataType:
		dti = 1
	case CardDataType:
		dti = 2
	case FileDataType:
		dti = 3
	case TextDataType:
		dti = 4
	}

	model, err := buildData(dti)
	if err != nil {
		return err
	}
	model.ID = data.ID
	model.Version = data.Version
	err = c.UpdateUserData(model)
	return err
}
