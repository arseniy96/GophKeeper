package application

import "github.com/arseniy96/GophKeeper/internal/client/models"

func (c *Client) StartWorker() {
	for {
		select {
		case dataID := <-c.dataSyncChan:
			c.Logger.Log.Debugf("add data with id %v to cache", dataID) // For testing cache
			m, err := c.gRPCClient.GetUserData(models.UserDataModel{ID: dataID})
			if err != nil {
				c.Logger.Log.Errorf("get user data error: %v", err)
				continue
			}
			c.UpdateDataInCache(m)
		default:
			continue
		}
	}
}
