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
		fmt.Printf("%v. %v\n", i+1, dt)
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
	case 2:
		return buildCardData()
	case 4:
		return buildTextData()
	default:
		return nil, UnknownDataType
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

func buildCardData() (*models.UserData, error) {
	utils.SlowPrint("Please enter data")
	card := &CardData{}
	fmt.Printf("card number: ")
	_, err := fmt.Scanln(&card.Number)
	if err != nil {
		return nil, err
	}
	fmt.Printf("expired date (mm/yy): ")
	_, err = fmt.Scanln(&card.ExpDate)
	if err != nil {
		return nil, err
	}
	fmt.Printf("card holder: ")
	_, err = fmt.Scanln(&card.CardHolder)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(card)
	return &models.UserData{
		DataType: CardDataType,
		Data:     byteData,
	}, nil
}

func buildTextData() (*models.UserData, error) {
	utils.SlowPrint("Please enter text")
	text := &TextData{}
	_, err := fmt.Scan(&text.Text)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(text)
	return &models.UserData{
		DataType: TextDataType,
		Data:     byteData,
	}, nil
}
