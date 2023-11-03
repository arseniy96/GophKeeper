package application

import "github.com/arseniy96/GophKeeper/internal/client/models"

func (c *Client) DataSyncWorker() {
	for {
		select {
		case id := <-c.DataIDSyncChan:
			c.Logger.Log.Debug(id) // For testing cache
			m, err := c.GetUserData(models.UserDataModel{ID: id})
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
