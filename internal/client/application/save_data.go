package application

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/mailru/easyjson"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

var dataTypes = [4]string{passwordDataType, cardDataType, fileDataType, textDataType}

// SaveData – сохранение данных пользователя.
// Сценарий:
//  1. Приложение предлагает пользователю выбрать тип данных, которые надо сохранить.
//  2. Приложение просит пользователя ввести необходимые данные для сохранения.
//  3. Приложение преобразует данные в байты и отправляет на сервер.
func (c *Client) SaveData() error {
	var (
		dti  int
		name string
	)
	c.printer.Print("What data type do you want to save?")
	for i, dt := range dataTypes {
		fmt.Printf("%v. %v\n", i+1, dt)
	}
	_, err := c.printer.Scan(&dti)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	model, err := buildData(dti, c.printer)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	c.printer.Print("What name to save the data with?")
	_, err = c.printer.Scan(&name)
	if err != nil {
		return fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	model.Name = name

	err = c.gRPCClient.SaveUserData(model)
	if err != nil {
		return err
	}
	c.UpdateDataInCache(model)

	c.printer.Print("Saved!")

	return nil
}

func buildData(dti int, p printer) (*models.UserData, error) {
	switch dti {
	case passwordData:
		return buildPassword(p)
	case cardData:
		return buildCardData(p)
	case fileData:
		return buildFileData(p)
	case textData:
		return buildTextData(p)
	default:
		return nil, ErrUnknownDataType
	}
}

func buildPassword(p printer) (*models.UserData, error) {
	p.Print("Please enter password data")
	pass := &PasswordData{}
	fmt.Print(siteInput)
	_, err := p.Scan(&pass.Site)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	fmt.Print(loginInput)
	_, err = p.Scan(&pass.Login)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	fmt.Print(passwordInput)
	_, err = p.Scan(&pass.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	byteData, err := easyjson.Marshal(pass)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	return &models.UserData{
		DataType: passwordDataType,
		Data:     byteData,
	}, nil
}

func buildCardData(p printer) (*models.UserData, error) {
	p.Print("Please enter card details")
	card := &CardData{}
	fmt.Print(cardNumberInput)
	_, err := p.Scan(&card.Number)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	fmt.Print(cardExpDateInput)
	_, err = p.Scan(&card.ExpDate)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	fmt.Print(cardHolderInput)
	_, err = p.Scan(&card.CardHolder)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	byteData, err := easyjson.Marshal(card)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	return &models.UserData{
		DataType: cardDataType,
		Data:     byteData,
	}, nil
}

func buildTextData(p printer) (*models.UserData, error) {
	text := &TextData{}
	p.Print("Please enter text")
	_, err := p.Scan(&text.Text)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	byteData, err := easyjson.Marshal(text)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	return &models.UserData{
		DataType: textDataType,
		Data:     byteData,
	}, nil
}

func buildFileData(p printer) (*models.UserData, error) {
	file := &FileData{}
	p.Print("Please enter path to file")
	_, err := p.Scan(&file.Path)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	openedFile, err := os.Open(file.Path)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	defer func() {
		_ = openedFile.Close()
	}()

	stat, err := openedFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(openedFile).Read(bs)
	if err != nil && errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}
	file.Data = bs

	byteData, err := easyjson.Marshal(file)
	if err != nil {
		return nil, fmt.Errorf("%w: something went wrong: %w", ErrInternal, err)
	}

	return &models.UserData{
		DataType: fileDataType,
		Data:     byteData,
	}, nil
}
