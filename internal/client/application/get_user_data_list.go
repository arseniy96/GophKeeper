package application

import (
	"fmt"
)

func (c *Client) GetUserDataList() error {
	records, err := c.gRPCClient.GetUserDataList()
	if err != nil {
		c.Logger.Log.Warnf("get user data list error: %v", err)
		// records = c.GetUserDataListFromCache() TODO

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
