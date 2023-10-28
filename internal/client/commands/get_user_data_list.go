package commands

import (
	"fmt"

	"github.com/arseniy96/GophKeeper/internal/client/utils"
)

func GetUserDataList(c Client) error {
	records, err := c.GetUserDataList()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return ErrNoData
	}

	utils.SlowPrint("You have these saved data:")
	for _, el := range records {
		fmt.Printf("id: %d, name: %s, type: %s, version: %d\n", el.ID, el.Name, el.DataType, el.Version)
	}

	return nil
}
