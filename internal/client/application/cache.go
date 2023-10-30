package application

import (
	"errors"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

var (
	ErrNoDataInCache = errors.New("no data in cache")
)

func (c *Client) AppendDataToCache(model *models.UserData) {
	c.Cache = append(c.Cache, clientCache{
		token:  c.AuthToken,
		dataID: model.ID,
		data:   model,
		actual: true,
	})
}

func (c *Client) UpdateDataInCache(model *models.UserData) {
	newData := true
	for _, m := range c.Cache {
		if m.dataID == model.ID {
			m.data = model
			m.actual = true
			newData = false
			break
		}
	}
	if newData {
		c.AppendDataToCache(model)
	}
}

func (c *Client) SyncCache(records []models.UserDataList) {
	cacheMap := make(map[int64]int64) // мапа версий данных, которые есть в кеше
	for _, m := range c.Cache {
		cacheMap[m.dataID] = m.data.Version
	}
	for _, r := range records {
		if v, ok := cacheMap[r.ID]; ok && v == r.Version {
			continue
		}
		c.DataIDSyncChan <- r.ID
	}

}

func (c *Client) GetUserDataFromCache(model models.UserDataModel) (*models.UserData, error) {
	for _, m := range c.Cache {
		if model.ID == m.dataID {
			return m.data, nil
		}
	}
	return nil, ErrNoDataInCache
}

func (c *Client) GetUserDataListFromCache() []models.UserDataList {
	records := make([]models.UserDataList, 0, len(c.Cache))
	for _, cacheData := range c.Cache {
		if c.AuthToken == cacheData.token {
			m := models.UserDataList{
				Name:     cacheData.data.Name,
				DataType: cacheData.data.DataType,
				ID:       cacheData.dataID,
				Version:  cacheData.data.Version,
			}
			records = append(records, m)
		}
	}

	return records
}
