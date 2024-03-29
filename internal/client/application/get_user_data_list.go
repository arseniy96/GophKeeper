package application

import (
	"fmt"
)

// GetUserDataList – получение всех сохранённых данных (мета-данных) пользователя.
func (c *Client) GetUserDataList() error {
	records, err := c.gRPCClient.GetUserDataList()
	if err != nil { // что-то с сервером
		c.Logger.Log.Warnf("get user data list error: %v", err)
		records = c.GetDataListFromCache() // пытаемся достать из кеша
	} else {
		c.SyncCache(records)
	}
	if len(records) == 0 {
		return ErrNoData
	}

	c.printer.Print("You have these saved data:")
	for _, el := range records {
		fmt.Printf("id: %d, name: %s, type: %s, version: %d\n", el.ID, el.Name, el.DataType, el.Version)
	}

	return nil
}
