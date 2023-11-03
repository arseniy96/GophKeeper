package clientcache

import (
	"errors"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

type Cache struct {
	mem map[int64]cacheData
}

type cacheData struct {
	dataID int64
	data   *models.UserData
}

var (
	ErrNotFound = errors.New(`not found in cache`)
)

func NewCache() *Cache {
	return &Cache{
		mem: make(map[int64]cacheData),
	}
}

func (c *Cache) Append(model *models.UserData) {
	c.mem[model.ID] = cacheData{
		dataID: model.ID,
		data:   model,
	}
}

func (c *Cache) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	d, ok := c.mem[model.ID]
	if !ok {
		return nil, ErrNotFound
	}

	return d.data, nil
}

func (c *Cache) GetUserDataList() []models.UserDataList {
	records := make([]models.UserDataList, 0, len(c.mem))
	for _, d := range c.mem {
		rec := models.UserDataList{
			Name:     d.data.Name,
			DataType: d.data.DataType,
			ID:       d.data.ID,
			Version:  d.data.Version,
		}
		records = append(records, rec)
	}

	return records
}
