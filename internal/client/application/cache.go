package application

import "github.com/arseniy96/GophKeeper/internal/client/models"

func (c *Client) SyncCache(records []models.UserDataList) {
	for _, rec := range records {
		cachedData, err := c.cache.GetUserData(models.UserDataModel{ID: rec.ID})
		if err != nil || cachedData.Version != rec.Version {
			c.appendDataToCache(rec.ID)
		}
	}
}

func (c *Client) appendDataToCache(dataID int64) {
	c.dataSyncChan <- dataID
}

func (c *Client) UpdateDataInCache(data *models.UserData) {
	c.cache.Append(data)
}

func (c *Client) GetDataFromCache(m models.UserDataModel) (*models.UserData, error) {
	c.Logger.Log.Debug("load data from cache...")
	return c.cache.GetUserData(m)
}

func (c *Client) GetDataListFromCache() []models.UserDataList {
	c.Logger.Log.Debug("load data list from cache...")
	return c.cache.GetUserDataList()
}
