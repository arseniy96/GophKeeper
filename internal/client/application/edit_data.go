package application

import (
	"errors"
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

// EditData – метод клиента, который позволяет редактировать ранее сохранённые данные.
// Сценарий:
//  1. Приложение забирает с сервера все сохранённые данные пользователя (мета данные).
//  2. Приложение предлагает пользователю выбрать из сохранённых данных те, которые пользователь хочет отредактировать.
//  3. Приложение достаёт нужные данные и предлагает пользователю ввести заново данные.
//  4. Приложение обновляем данные, сохраняя название и id, и отправляет на сервер.
func (c *Client) EditData() error {
	err := c.GetUserDataList()
	if err != nil {
		if errors.Is(err, ErrNoData) {
			//nolint:goconst,nolintlint // it's print
			c.printer.Print("You have no saved data")
			return nil
		}
		return err
	}

	c.printer.Print("Please enter data id for update")
	var dataID int64
	_, err = c.printer.Scan(&dataID)
	if err != nil {
		return fmt.Errorf(InternalErrTemplate, ErrInternal, err)
	}

	data, err := c.gRPCClient.GetUserData(models.UserDataModel{ID: dataID})
	if err != nil {
		return err
	}
	dti := 0
	switch data.DataType {
	case passwordDataType:
		dti = 1
	case cardDataType:
		dti = 2
	case fileDataType:
		dti = 3
	case textDataType:
		dti = 4
	}

	model, err := buildData(dti, c.printer)
	if err != nil {
		return err
	}
	model.ID = data.ID
	model.Version = data.Version

	return c.gRPCClient.UpdateUserData(model)
}
