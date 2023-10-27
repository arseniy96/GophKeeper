package commands

import (
	"fmt"

	"github.com/mailru/easyjson"

	"github.com/arseniy96/GophKeeper/internal/client/models"
	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

func SaveData(c Client) error {
	var (
		DataTypes = [4]string{PasswordDataType, CardDataType, FileDataType, TextDataType}
		dti       int
		name      string
	)
	utils.SlowPrint("What data type do you want to save?")
	for i, dt := range DataTypes {
		fmt.Printf("%v. %v\n", i, dt)
	}
	_, err := fmt.Scanln(&dti)
	if err != nil {
		return err
	}

	model, err := buildData(dti)
	if err != nil {
		return err
	}

	utils.SlowPrint("What name to save the data with?")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return err
	}
	model.Name = name

	err = c.SaveUserData(model)
	if err != nil {
		return err
	}

	utils.SlowPrint("Saved!")

	return nil
}

func buildData(dti int) (*models.UserData, error) {
	switch dti {
	case 1:
		return buildPassword()
	//case 2:
	//	return buildCardData()
	default:
		return nil, fmt.Errorf("unknown data type")
	}
}

func buildPassword() (*models.UserData, error) {
	utils.SlowPrint("Please enter data")
	pass := &PasswordData{}
	fmt.Printf("site: ")
	_, err := fmt.Scanln(&pass.Site)
	if err != nil {
		return nil, err
	}
	fmt.Printf("login: ")
	_, err = fmt.Scanln(&pass.Login)
	if err != nil {
		return nil, err
	}
	fmt.Printf("password: ")
	_, err = fmt.Scanln(&pass.Password)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(pass)
	return &models.UserData{
		DataType: PasswordDataType,
		Data:     byteData,
	}, nil
}
